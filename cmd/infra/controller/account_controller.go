package controller

import (
	"net/http"

	"github.com.br/gibranct/ride/cmd/application"
	"github.com.br/gibranct/ride/cmd/application/usecase"
	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountService application.AccountService
}

func (accountCtrl *AccountController) SignUpHandler(c echo.Context) error {
	var input usecase.SignUpInput

	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	output, err := accountCtrl.accountService.SignUp.Execute(input)

	if err != nil {
		response := map[string]any{"message": err.Error()}
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
