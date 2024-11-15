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

func Test_SignUpDriver(t *testing.T) {
	type SignUpOutput struct {
		AccountId string `json:"accountId"`
	}
	signUpJson := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
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

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, responseBody.AccountId)
	}
}

func Test_SignUpDriverWithInvalidCarPlate(t *testing.T) {
	type SignUpOutput struct {
		Message int `json:"message"`
	}
	signUpJson := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA123",
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

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, responseBody.Message, -5)
	}
}

func Test_SignUpPassenger(t *testing.T) {
	type SignUpOutput struct {
		AccountId string `json:"accountId"`
	}
	signUpJson := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret",
	}
	jsonBytes, _ := json.Marshal(signUpJson)
	var responseBody SignUpOutput

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(jsonBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, responseBody.AccountId)
	}
}

func Test_SignUpPassengerWithInvalidCPF(t *testing.T) {
	type SignUpOutput struct {
		Message int `json:"message"`
	}
	signUpJson := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "9745632155",
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

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, responseBody.Message, -1)
	}
}

func Test_SignUpPassengerWithInvalidEmail(t *testing.T) {
	type SignUpOutput struct {
		Message int `json:"message"`
	}
	signUpJson := Account{
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

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, responseBody.Message, -2)
	}
}

func Test_SignUpPassengerWithInvalidName(t *testing.T) {
	type SignUpOutput struct {
		Message int `json:"message"`
	}
	signUpJson := Account{
		Name:        "John",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
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

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, responseBody.Message, -3)
	}
}

func Test_SignUpDuplicatedPassenger(t *testing.T) {
	type SignUpOutput struct {
		Message int `json:"message"`
	}
	signUpJson := Account{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
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
	SignUp(ctx)

	req = httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(jsonBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)

	if assert.NoError(t, SignUp(ctx)) {
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, responseBody.Message, -4)
	}
}

func Test_SignUpWithInvalidJSON(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/sign-up", strings.NewReader("{"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.Error(t, SignUp(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
