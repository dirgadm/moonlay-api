package ehttp

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	ResponseFormat *FormatResponse
	ResponseData   interface{}
}

func NewContext(c echo.Context) *Context {
	return &Context{c, NewResponse(), nil}
}

func NewResponse() *FormatResponse {
	return &FormatResponse{
		Code:   http.StatusOK,
		Status: "success",
	}
}

func HTTPErrorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		ctx, ok := c.(*Context)
		if !ok {
			ctx = NewContext(c)
		}
		ctx.Serve(err)
	}
}

func (c *Context) GetParamUri(paramName string) (id int, err error) {
	paramValue := c.Param(paramName)
	var value int
	value, err = strconv.Atoi(paramValue)
	id = value
	return
}

func (c *Context) GetParamInt(param string) (value int) {
	paramValue := c.QueryParam(param)
	value, _ = strconv.Atoi(paramValue)
	return
}

func (c *Context) GetParamString(param string) (value string) {
	value = c.QueryParam(param)
	return
}

func (c *Context) Data(data interface{}, total ...int64) {
	c.ResponseFormat.SetData(data)
}

func (c *Context) DataList(data interface{}, total int64, page int, perPage int) {
	c.ResponseFormat.SetDataList(data, total, page, perPage)
}

func (c *Context) Failure(fail ...string) {
	c.ResponseFormat.Errors = map[string]string{fail[0]: fail[1]}
}

func (c *Context) Serve(e error) (err error) {
	c.ResponseFormat.Code = http.StatusOK
	if e != nil {
		c.ResponseFormat.SetError(e)
	} else {
		if c.ResponseData != nil {
			c.ResponseFormat.SetData(c.ResponseData)
		}
	}

	if c.Request().Method == echo.HEAD || c.Request().Method == echo.OPTIONS {
		err = c.NoContent(http.StatusNoContent)
	} else {
		err = c.JSON(c.ResponseFormat.Code, c.ResponseFormat)
	}

	c.ResponseFormat.Reset()
	return
}

func (c *Context) Message(statusCode int, status string, message string) (err error) {
	c.ResponseFormat.Code = statusCode
	c.ResponseFormat.Status = status
	c.ResponseFormat.Message = message

	if c.Request().Method == echo.HEAD || c.Request().Method == echo.OPTIONS {
		err = c.NoContent(http.StatusNoContent)
	} else {
		err = c.JSON(c.ResponseFormat.Code, c.ResponseFormat)
	}
	c.ResponseFormat.Reset()
	return
}

func (r *FormatResponse) SetError(err error) *FormatResponse {
	if he, ok := err.(*echo.HTTPError); ok {
		r.Code = he.Code
		r.Status = "failed"
		r.Message = http.StatusText(r.Code)
		r.Errors = map[string]string{
			"error": fmt.Sprintf("%v", he.Message),
		}
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		r.Code = http.StatusBadRequest
		r.Status = "failed"
		r.Message = http.StatusText(r.Code)
		errorMaps := map[string]string{}

		for _, e := range ve {
			field := strings.ToLower(e.StructField())
			var message string
			switch e.Tag() {
			case "required":
				message = fmt.Sprintf("The %s is required", field)
			case "gte":
				message = fmt.Sprintf("The %s must be greater than %s", field, e.Param())
			case "lte":
				message = fmt.Sprintf("The %s must be lower than %s", field, e.Param())
			default:
				message = fmt.Sprintf("The %s is invalid", field)
			}
			errorMaps[field] = message
		}
		r.Errors = errorMaps
	} else if o, ok := err.(*Errors); ok {
		// validation error
		r.Code = http.StatusBadRequest
		r.Status = "failed"
		r.Message = http.StatusText(r.Code)
		r.Errors = o.Messages()
	} else {
		message := err.Error()
		r.Code = http.StatusBadRequest
		r.Status = "failed"
		r.Message = http.StatusText(r.Code)
		r.Errors = map[string]string{
			"error": message,
		}
	}

	return r
}

func (r *FormatResponse) Reset() {
	r.Data = nil
	r.Errors = nil
	r.Message = nil
}
