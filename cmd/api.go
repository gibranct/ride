package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const PORT = "127.0.0.1:3333"

func StartServer() {
	e := echo.New()

	e.POST("/sign-up", SignUpHandler)

	e.Logger.Fatal(e.Start(PORT))
}

func SignUpHandler(c echo.Context) error {
	var input Account

	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	output, err := SignUp(input)

	if err != nil {
		response := map[string]any{"message": err.Error()}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	return c.JSON(http.StatusCreated, output)
}
