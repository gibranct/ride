package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com.br/gibranct/account/internal/application"
	"github.com.br/gibranct/account/internal/infra/controller/dto"
	myHttp "github.com.br/gibranct/account/internal/infra/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func runTestServer() (*httptest.Server, *postgres.PostgresContainer) {
	dbContainer, connString := getDBConnStrAndContainer()
	os.Setenv("POSTGRES_DSN", connString)

	httpServer := myHttp.NewHttpServer(application.NewApplication())
	httpServer.SetUpRoutes()

	return httptest.NewServer(httpServer.GetHandler()), dbContainer
}

func Test_SignUpAndGetAccount(t *testing.T) {
	testServer, dbContainer := runTestServer()

	fakeId := uuid.NewString()
	testCases := []struct {
		name                   string
		expectedSignUpCode     int
		expectedGetAccountCode int
		accountId              *string
		signUpRequest          *dto.SignUpInputRequestDto
	}{
		{"should create passenger account and fetch account successfully", 201, 200, nil, getSignUpRequestPassengerDto()},
		{"should create driver account and fetch account successfully", 201, 200, nil, getSignUpRequestDriverDto()},
		{"should create passenger account and not find account", 201, 404, &fakeId, getSignUpRequestPassengerDto()},
		{"should not create account with empty name and not find account", 400, 404, &fakeId, &dto.SignUpInputRequestDto{
			Name:        "",
			Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
			CPF:         "97456321558",
			CarPlate:    "AAA1234",
			IsPassenger: false,
			IsDriver:    true,
			Password:    "secret123",
		}},
	}

	defer t.Cleanup(func() {
		err := dbContainer.Terminate(context.Background())
		assert.NoError(t, err)
		testServer.Close()
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			jsonBytes, err := json.Marshal(tc.signUpRequest)
			assert.NoError(t, err)

			request1, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/sign-up", testServer.URL), bytes.NewBuffer(jsonBytes))
			assert.NoError(t, err)
			request1.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

			response, err := http.DefaultClient.Do(request1)
			assert.NoError(t, err)
			assert.Equal(t, response.StatusCode, tc.expectedSignUpCode)

			var resp dto.SignUpInputResponseDto
			err = json.NewDecoder(response.Body).Decode(&resp)
			assert.NoError(t, err)

			var actualAccountId *string
			if tc.accountId != nil {
				actualAccountId = tc.accountId
			} else {
				actualAccountId = &resp.AccountId
			}

			request2, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/accounts/%s", testServer.URL, *actualAccountId), nil)
			assert.NoError(t, err)

			response, err = http.DefaultClient.Do(request2)
			assert.NoError(t, err)

			var account dto.GetAccountResponseDto
			err = json.NewDecoder(response.Body).Decode(&account)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedGetAccountCode, response.StatusCode)

			if tc.expectedGetAccountCode == http.StatusOK {
				assert.Equal(t, account.ID, resp.AccountId)
				assert.Equal(t, account.Name, tc.signUpRequest.Name)
				assert.Equal(t, account.Email, tc.signUpRequest.Email)
				assert.Equal(t, account.CPF, tc.signUpRequest.CPF)
				assert.Equal(t, account.IsPassenger, tc.signUpRequest.IsPassenger)
				assert.Equal(t, account.IsDriver, tc.signUpRequest.IsDriver)
				assert.Equal(t, account.CarPlate, tc.signUpRequest.CarPlate)
			}
		})
	}
}

func getSignUpRequestPassengerDto() *dto.SignUpInputRequestDto {
	return &dto.SignUpInputRequestDto{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "secret123",
	}
}

func getSignUpRequestDriverDto() *dto.SignUpInputRequestDto {
	return &dto.SignUpInputRequestDto{
		Name:        "John Doe",
		Email:       fmt.Sprintf("john_%d@mail.com", rand.Int()),
		CPF:         "97456321558",
		CarPlate:    "AAA1234",
		IsPassenger: false,
		IsDriver:    true,
		Password:    "secret123",
	}
}
