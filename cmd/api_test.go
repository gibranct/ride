package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_SignUpDriverAPI(t *testing.T) {
	type SignUpOutput struct {
		AccountId string `json:"accountId"`
	}
	account := SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}
	jsonBytes, _ := json.Marshal(account)
	var responseBody SignUpOutput
	var newAccount GetAccountOutput

	e := echo.New()

	// Sign up
	req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(jsonBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, SignUpHandler(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, responseBody.AccountId)
	}

	// GetAccount
	path := "/v1/accounts/:id"
	req = httptest.NewRequest(http.MethodGet, path, nil)
	req.Header.Set(echo.HeaderAcceptEncoding, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)
	ctx.SetPath(path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(responseBody.AccountId)

	if assert.NoError(t, GetAccountByIDHandler(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &newAccount)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseBody.AccountId, newAccount.ID)
		assert.Equal(t, account.Name, newAccount.Name)
		assert.Equal(t, account.Email, newAccount.Email)
		assert.Equal(t, account.CPF, newAccount.CPF)
		assert.Equal(t, account.IsDriver, newAccount.IsDriver)
		assert.Equal(t, account.IsPassenger, newAccount.IsPassenger)
		assert.Equal(t, account.CarPlate, newAccount.CarPlate)
	}
}

func Test_SignUpPassengerWithInvalidEmailAPI(t *testing.T) {
	type SignUpOutput struct {
		Message string `json:"message"`
	}
	signUpJson := SignUpInput{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d_mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret",
	}
	jsonBytes, _ := json.Marshal(signUpJson)
	var responseBody SignUpOutput

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(jsonBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, SignUpHandler(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, responseBody.Message, "invalid email")
	}
}

func Test_SignUpWithInvalidJSON(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader("{"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.Error(t, SignUpHandler(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
