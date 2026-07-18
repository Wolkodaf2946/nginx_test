package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

// Success отправляет успешный ответ
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status: StatusSuccess,
		Data:   data,
	})
}

// ErrorResponse отправляет ошибку с заданным HTTP-кодом
func ErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, Response{
		Status:  StatusError,
		Message: message,
	})
}

// ValidationError разбирает ошибки тегов binding:"required,..."
func ValidationError(c *gin.Context, err error) {
	var errMsgs []string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid (%s)", e.Field(), e.Tag()))
		}
	} else {
		errMsgs = append(errMsgs, "invalid request body")
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Status:  StatusError,
		Message: strings.Join(errMsgs, ", "),
	})
}
