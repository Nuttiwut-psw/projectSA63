package main

import (
	"context"
	//"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/Toneone11/app/controllers"
	"github.com/Toneone11/app/ent"

)

//Users defines the struct for the Users object
type Users struct {
	User []User
}

//User defines the struct for the User object
type User struct {
	Doctorname  string
	Doctoremail string
}

//Operationrooms defines the struct for the Operationrooms object
type Operationrooms struct {
	Operationroom []Operationroom
}

//Operationroom defines the struct for the Operationroom object
type Operationroom struct {
	operationroomname string
}

//Patients defines the struct for the Patients object
type Patients struct {
	Patient []Patient
}

//Patient defines the struct for the Patient object
type Patient struct {
	patientname  string
	patientage   int
}

// @title SUT SA Project API Booking
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewBookingController(v1, client)
	controllers.NewOperationroomController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewUserController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"doctor Test1", "docorTest1@gmail.com"},
			User{"doctor Test2", "doctorTest2@gmail.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetDoctorName(u.Doctorname).
			SetDoctorEmail(u.Doctoremail).
			Save(context.Background())
	}

	// Set Operationrooms Data
	Operationrooms := Operationrooms{
		Operationroom: []Operationroom{
			Operationroom{"room01"},
			Operationroom{"room02"},
			Operationroom{"room03"},
			Operationroom{"room04"},
			Operationroom{"room05"},
		},
	}

	for _, o := range Operationrooms.Operationroom {
		client.Operationroom.
			Create().
			SetOperationroomName(o.operationroomname).
			Save(context.Background())
	}

	// Set Patient Data
	Patients := Patients{
		Patient: []Patient{
			Patient{"patient Test1", 22},
			Patient{"patient Test2", 33},
			Patient{"patient Test3", 44},
		},
	}

	for _, p := range Patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.patientname).
			SetPatientAge(p.patientage).
			Save(context.Background())
	}
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}