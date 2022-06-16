package helper

import (
	"encoding/json"
	"github.com/firdausalif/challenge-todolist/pkg/constant"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	successJson struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	errorJson struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    struct{} `json:"data"`
	}

	successDeleteJson struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    struct{} `json:"data"`
	}

	ErrorWithCode struct {
		CodeID     int         `json:"-"`
		Msg        string      `json:"message"`
		Status     string      `json:"status"`
		StatusCode int         `json:"-"`
		Data       interface{} `json:"data,omitempty"`
	}
)

func NewErrorMsg(code int, err error) error {
	msg := constant.CodeMapping[code]
	return ErrorWithCode{
		CodeID:     code,
		Msg:        msg,
		Status:     msg,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func NewErrorRecordNotFound(code int, err error) error {
	msg := constant.CodeMapping[code]
	return ErrorWithCode{
		CodeID:     code,
		Msg:        msg,
		Status:     msg,
		StatusCode: http.StatusNotFound,
		Data:       err.Error(),
	}
}

func (c ErrorWithCode) Error() string {
	b, _ := json.Marshal(&c)

	return string(b)
}

func JsonSUCCESS(c *fiber.Ctx, data interface{}) error {
	res := successJson{
		Message: "Success",
		Status:  "Success",
		Data:    data,
	}

	return c.Status(http.StatusOK).JSON(res)
}

func JsonSuccessDelete(c *fiber.Ctx) error {
	res := successDeleteJson{
		Message: "Success",
		Status:  "Success",
	}

	return c.Status(http.StatusOK).JSON(res)
}
func JsonCreated(c *fiber.Ctx, data interface{}) error {
	res := successJson{
		Message: "Success",
		Status:  "Success",
		Data:    data,
	}
	return c.Status(http.StatusCreated).JSON(res)
}

func JsonValidationError(c *fiber.Ctx, message string) error {
	res := errorJson{
		Message: message,
		Status:  "Bad Request",
	}
	return c.Status(http.StatusBadRequest).JSON(res)
}

func JsonNotFound(c *fiber.Ctx, message string) error {
	res := errorJson{
		Message: message,
		Status:  "Not Found",
	}

	return c.Status(http.StatusNotFound).JSON(res)
}

func JsonERROR(c *fiber.Ctx, err error) error {

	switch err.(type) {
	case ErrorWithCode:
		errMsg := err.(ErrorWithCode)
		res := ErrorWithCode{
			CodeID: http.StatusUnprocessableEntity,
			Msg:    errMsg.Msg,
		}

		return c.Status(errMsg.StatusCode).JSON(res)
	default:
		res := ErrorWithCode{
			CodeID: http.StatusInternalServerError,
			Msg:    "Error message type not defined.",
		}
		return c.Status(res.CodeID).JSON(res)
	}
}
