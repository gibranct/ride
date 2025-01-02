package usecase_test

import (
	"testing"

	"github.com.br/gibranct/account/internal/application/usecase"
	"github.com.br/gibranct/account/internal/domain/entity"
	"github.com.br/gibranct/account/internal/domain/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) GetAccountByEmail(email string) (*entity.Account, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *MockAccountRepository) GetAccountByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *MockAccountRepository) SaveAccount(account entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type MockMailerGateway struct {
	mock.Mock
}

func (m *MockMailerGateway) Send(to, subject, body string) {
	m.Called(to, subject, body)
}

func Test_SignUpExecute_EmailAlreadyTaken(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	mockMailer := new(MockMailerGateway)
	signUpUseCase := usecase.NewSignUpUseCase(mockRepo, mockMailer)

	input := usecase.SignUpInput{
		Name:        "John Doe",
		Email:       "john@doe.com",
		CPF:         "97456321558",
		CarPlate:    "XYZ1234",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "securepassword",
	}

	existingAccount := &entity.Account{ID: "1234"}
	mockRepo.On("GetAccountByEmail", input.Email).Return(existingAccount, nil)

	output, err := signUpUseCase.Execute(input)

	assert.Nil(t, output)
	assert.Equal(t, errors.ErrEmailAlreadyTaken, err)
	mockRepo.AssertCalled(t, "GetAccountByEmail", input.Email)
	mockRepo.AssertNotCalled(t, "SaveAccount", mock.Anything)
	mockMailer.AssertNotCalled(t, "Send", mock.Anything, mock.Anything, mock.Anything)
}

func Test_SignUpExecute_SuccessfulCreation(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	mockMailer := new(MockMailerGateway)
	signUpUseCase := usecase.NewSignUpUseCase(mockRepo, mockMailer)

	input := usecase.SignUpInput{
		Name:        "Jane Doe",
		Email:       "jane@doe.com",
		CPF:         "12345678909",
		CarPlate:    "ABC1234",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "anothersecurepassword",
	}

	mockRepo.On("GetAccountByEmail", input.Email).Return(nil, nil)
	mockRepo.On("SaveAccount", mock.Anything).Return(nil)
	mockMailer.On("Send", input.Email, "Welcome!", "...")

	output, err := signUpUseCase.Execute(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.NoError(t, uuid.Validate(output.AccountId))
	mockRepo.AssertCalled(t, "GetAccountByEmail", input.Email)
	mockRepo.AssertCalled(t, "SaveAccount", mock.Anything)
	mockMailer.AssertCalled(t, "Send", input.Email, "Welcome!", "...")
}

func Test_SignUpExecute_AccountCreationFails(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	mockMailer := new(MockMailerGateway)
	signUpUseCase := usecase.NewSignUpUseCase(mockRepo, mockMailer)

	input := usecase.SignUpInput{
		Name:        "Invalid User",
		Email:       "invalid@user.com",
		CPF:         "invalidcpf",
		CarPlate:    "XYZ1234",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "weakpassword",
	}

	// Simulate account creation failure
	mockRepo.On("GetAccountByEmail", input.Email).Return(nil, nil)

	output, err := signUpUseCase.Execute(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrInvalidCPF, err) // Assuming ErrInvalidCPF is the error returned by CreateAccount
	mockRepo.AssertNotCalled(t, "SaveAccount", mock.Anything)
	mockMailer.AssertNotCalled(t, "Send", mock.Anything, mock.Anything, mock.Anything)
}

func Test_SignUpExecute_ErrorRetrievingAccountByEmail(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	mockMailer := new(MockMailerGateway)
	signUpUseCase := usecase.NewSignUpUseCase(mockRepo, mockMailer)

	input := usecase.SignUpInput{
		Name:        "Error User",
		Email:       "error@user.com",
		CPF:         "12345678909",
		CarPlate:    "XYZ1234",
		IsPassenger: true,
		IsDriver:    false,
		Password:    "securepassword",
	}

	mockRepo.On("GetAccountByEmail", input.Email).Return(nil, errors.ErrDatabase)

	output, err := signUpUseCase.Execute(input)

	assert.Nil(t, output)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrDatabase, err)
	mockRepo.AssertCalled(t, "GetAccountByEmail", input.Email)
	mockRepo.AssertNotCalled(t, "SaveAccount", mock.Anything)
	mockMailer.AssertNotCalled(t, "Send", mock.Anything, mock.Anything, mock.Anything)
}
