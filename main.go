package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type InputData struct {
	VAL int `json:"value"`
	BAL int `json:"balance"`
	FQ  int `json:"frequ"`
}

type ResponseData struct {
	Resp int    `json:"response"`
	Str  string `json:"string"`
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.POST("/calc", calculate)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func calculate(c echo.Context) error {
	var in InputData
	err := c.Bind(&in)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Wrong data")
	}

	res := CalclucateSomething(in)

	return echo.NewHTTPError(http.StatusOK, res)
}

func CalclucateSomething(inp InputData) ResponseData {
	var response ResponseData
	response.Resp = inp.BAL + inp.VAL
	response.Str = "Все ОК"
	return response
}
