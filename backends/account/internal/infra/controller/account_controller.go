package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/infra/controller/dto"
	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountService application.AccountService
}

func (accountCtrl *AccountController) SignUpHandler(c echo.Context) error {
	var input dto.SignUpInputRequestDto

	if err := c.Bind(&input); err != nil {
		fmt.Printf("SignUp input: %+v\n", err)
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	output, err := accountCtrl.accountService.SignUp.Execute(context.Background(), input.ToSignUpInput())

	if err != nil {
		return writeError(err, c)
	}

	return c.JSON(http.StatusCreated, dto.NewSignUpInputResponseDto(output.AccountId))
}

func (accountCtrl *AccountController) GetAccountByIDHandler(c echo.Context) error {
	accountId := c.Param("id")
	account, err := accountCtrl.accountService.GetAccount.Execute(context.Background(), accountId)

	if err != nil {
		return writeError(err, c)
	}

	return c.JSON(http.StatusOK, dto.NewGetAccountResponseDto(account))
}

func NewAccountController(accountService *application.AccountService) *AccountController {
	return &AccountController{
		accountService: *accountService,
	}
}
