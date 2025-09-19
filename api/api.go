package api

import (
	"github.com/BellaMez/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dataBase := db.Init()
	studentsDB := db.NewStudentHandler(dataBase)

	return &API{
		Echo: e,
		DB:   studentsDB,
	}

}

func (api *API) Configureroutes() {

	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/student/:id", api.getStudent)
	api.Echo.PUT("/student/:id", api.updateStudent)
	api.Echo.DELETE("/student/:id", api.deleteStudent)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")

}
