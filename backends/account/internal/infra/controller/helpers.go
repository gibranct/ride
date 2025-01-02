package controller

import (
	"errors"
	"net/http"

	myErrors "github.com.br/gibranct/account/internal/domain/errors"
	"github.com/labstack/echo/v4"
)

type envelope map[string]any

func writeError(err error, e echo.Context) error {
	if errors.Is(err, myErrors.ErrAccountNotFound) {
		return e.JSON(http.StatusNotFound, envelope{"error": err.Error()})
	}

	for _, cError := range myErrors.AllDomainErrors() {
		if errors.Is(err, cError) {
			return e.JSON(http.StatusBadRequest, envelope{"error": err.Error()})
		}
	}

	return e.JSON(http.StatusInternalServerError, envelope{"error": err.Error()})
}
