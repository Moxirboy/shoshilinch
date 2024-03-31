package utils

import (
	"database/sql"
	"errors"
	"net/http"

	dto2 	"shoshilinch/internal/controller/v1/dto"


	"github.com/jackc/pgx"
	echo 	"github.com/gin-gonic/gin"
)

const (
	ErrMessageUnknown  = "Unknown server error"
	ErrMessageNotFound = "Not found"
)

const (
	StatusNotFound            = "NOT_FOUND"
	StatusInternalServerError = "INTERNAL_SERVER_ERROR"
	StatusInvalidField        = "INVALID_FIELD"
	StatusAlreadyExist        = "ALREADY_EXIST"
)

var (
	ErrInvalidField      = errors.New("invelid field")
	ErrAlreadyExist      = errors.New("already exist")
	ErrClientAlreadyUsed = errors.New("clientID: this client already used")
	
)

func SendResponse(
	ctx *echo.Context,
	response interface{},
	err error,
) error {
	if err != nil {
		code, status, message := parseError(err)
		ctx.JSON(
			code,
			dto2.BaseResponse{
				Status:  status,
				Message: message,
			},
		)
	}

	switch response.(type) {
	case []*dto2.InvalidParams:
		ctx.JSON(
			http.StatusBadRequest,
			dto2.BaseResponse{
				Data: response,
			})
	}

	ctx.JSON(
		http.StatusOK,
		dto2.BaseResponse{
			Data: response,
		},
	)
	return nil
}

func parseError(err error) (code int, status, message string) {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
		return http.StatusNotFound, StatusNotFound, ErrMessageNotFound
	} else if errors.Is(err, ErrInvalidField) {
		return http.StatusBadRequest, StatusInvalidField, ""
	} else if errors.Is(err, ErrAlreadyExist) {
		return http.StatusBadRequest, StatusAlreadyExist, ""
	} else if errors.Is(err, ErrClientAlreadyUsed) {
		return http.StatusBadRequest, StatusAlreadyExist, ErrClientAlreadyUsed.Error()
	}
	return http.StatusInternalServerError, StatusInternalServerError, err.Error()
}