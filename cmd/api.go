package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const PORT = "127.0.0.1:3333"

func StartServer() {
	e := echo.New()

	e.POST("/sign-up", SignUpHandler)

	e.GET("/v1/accounts/:id", GetAccountByIDHandler)

	e.Logger.Fatal(e.Start(PORT))
}

func SignUpHandler(c echo.Context) error {
	signUp := NewSignUpUseCase(NewAccountDAO())
	var input Account

	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	output, err := signUp.Execute(input)

	if err != nil {
		response := map[string]any{"message": err.Error()}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	return c.JSON(http.StatusCreated, output)
}

func GetAccountByIDHandler(c echo.Context) error {
	getAccount := NewGetAccountCase(NewAccountDAO())
	accountId := c.Param("id")

	account, err := getAccount.Execute(accountId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}
