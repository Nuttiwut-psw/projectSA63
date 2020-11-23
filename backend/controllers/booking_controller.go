package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Toneone11/app/ent"
	"github.com/Toneone11/app/ent/booking"
	"github.com/Toneone11/app/ent/operationroom"
	"github.com/Toneone11/app/ent/patient"
	"github.com/Toneone11/app/ent/user"
	"github.com/gin-gonic/gin"
)

// BookingController defines the struct for the booking controller
type BookingController struct {
	client *ent.Client
	router gin.IRouter
}

// Booking defines the struct for the booking object
type Booking struct {
	User          int
	Patient       int
	Operationroom int
	Added         string
}

// CreateBooking handles POST requests for adding booking entities
// @Summary Create booking
// @Description Create booking
// @ID create-booking
// @Accept   json
// @Produce  json
// @Param booking body Booking true "Booking entity"
// @Success 200 {object} Booking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /booking [post]
func (ctl *BookingController) CreateBooking(c *gin.Context) {
	obj := Booking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
	}

	o, err := ctl.client.Operationroom.
		Query().
		Where(operationroom.IDEQ(int(obj.Operationroom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "operationroom not found",
		})
	}

	time, err := time.Parse(time.RFC3339, obj.Added)
	b, err := ctl.client.Booking.
		Create().
		SetDoctorID(u).
		SetPatientID(p).
		SetOperationroomID(o).
		SetDate(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// ListBooking handles request to get a list of booking entities
// @Summary List booking entities
// @Description list booking entities
// @ID list-booking
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Booking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /booking [get]
func (ctl *BookingController) ListBooking(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	booking, err := ctl.client.Booking.
		Query().
		WithDoctorID().
		WithPatientID().
		WithOperationroomID().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, booking)
}

// DeleteBooking handles DELETE requests to delete a booking entity
// @Summary Delete a booking entity by ID
// @Description get booking by ID
// @ID delete-booking
// @Produce  json
// @Param id path int true "Booking ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /booking/{id} [delete]
func (ctl *BookingController) DeleteBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Booking.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// GetBooking handles GET requests to retrieve a booking entity
// @Summary Get a booking entity by ID
// @Description get booking by ID
// @ID get-booking
// @Produce  json
// @Param id path int true "Booking ID"
// @Success 200 {object} ent.Booking
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /booking/{id} [get]
func (ctl *BookingController) GetBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.Booking.
		Query().
		Where(booking.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// NewBookingController creates and registers handles for the booking controller
func NewBookingController(router gin.IRouter, client *ent.Client) *BookingController {
	bc := &BookingController{
		client: client,
		router: router,
	}

	bc.register()

	return bc

}

func (ctl *BookingController) register() {
	booking := ctl.router.Group("/booking")

	booking.POST("", ctl.CreateBooking)
	booking.GET("", ctl.ListBooking)
	booking.GET(":id", ctl.GetBooking)
	booking.DELETE(":id", ctl.DeleteBooking)
}
