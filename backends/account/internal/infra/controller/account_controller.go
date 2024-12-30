package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/application/usecase"
	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountService application.AccountService
}

func (accountCtrl *AccountController) SignUpHandler(c echo.Context) error {
	var input usecase.SignUpInput

	if err := c.Bind(&input); err != nil {
		fmt.Printf("SignUp input: %+v\n", err)
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	fmt.Printf("SignUp input: %+v\n", input)

	output, err := accountCtrl.accountService.SignUp.Execute(input)

	fmt.Printf("SignUp output: %+v\n", output)

	if err != nil {
		response := map[string]any{"message": err.Error()}
		log.Default().Println(response)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	return c.JSON(http.StatusCreated, output)
}

func (accountCtrl *AccountController) GetAccountByIDHandler(c echo.Context) error {
	accountId := c.Param("id")
	account, err := accountCtrl.accountService.GetAccount.Execute(accountId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)
}

func NewAccountController(accountService *application.AccountService) *AccountController {
	return &AccountController{
		accountService: *accountService,
	}
}
