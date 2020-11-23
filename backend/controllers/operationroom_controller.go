package controllers

import (
   "context"
   "strconv"
   "github.com/Toneone11/app/ent"
   "github.com/Toneone11/app/ent/operationroom"
   "github.com/gin-gonic/gin"
)

// OperationroomController defines the struct for the operationroom controller
type OperationroomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateOperationroom handles POST requests for adding operationroom entities
// @Summary Create operationroom
// @Description Create operationroom
// @ID create-operationroom
// @Accept   json
// @Produce  json
// @Param operationroom body ent.Operationroom true "Operationroom entity"
// @Success 200 {object} ent.Operationroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operationroom [post]
func (ctl *OperationroomController) CreateOperationroom(c *gin.Context) {
	obj := ent.Operationroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "operationroom binding failed",
		})
		return
	}

	o, err := ctl.client.Operationroom.
		Create().
		SetOperationroomName(obj.OperationroomName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, o)
}

// GetOperationroom handles GET requests to retrieve a operationroom entity
// @Summary Get a operationroom entity by ID
// @Description get operationroom by ID
// @ID get-operationroom
// @Produce  json
// @Param id path int true "Operationroom ID"
// @Success 200 {object} ent.Operationroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operationroom/{id} [get]
func (ctl *OperationroomController) GetOperationroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	o, err := ctl.client.Operationroom.
		Query().
		Where(operationroom.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, o)
}

// ListOperationroom handles request to get a list of operationroom entities
// @Summary List operationroom entities
// @Description list operationroom entities
// @ID list-operationroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Operationroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /operationroom [get]
func (ctl *OperationroomController) ListOperationroom(c *gin.Context) {
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

	operationroom, err := ctl.client.Operationroom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, operationroom)
}

// NewOperationroomController creates and registers handles for the operationroom controller
func NewOperationroomController(router gin.IRouter, client *ent.Client) *OperationroomController {
	o := &OperationroomController{
		client: client,
		router: router,
	}

	o.register()

	return o

}

func (ctl *OperationroomController) register() {
	operationroom := ctl.router.Group("/operationroom")

	operationroom.GET("", ctl.ListOperationroom)
	operationroom.POST("", ctl.CreateOperationroom)
	operationroom.GET(":id", ctl.GetOperationroom)
}