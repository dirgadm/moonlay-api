package ehttp

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type Binder struct{}

func (b Binder) Bind(i interface{}, ctx echo.Context) (err error) {
	req := ctx.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	if strings.HasPrefix(ctype, echo.MIMEApplicationJSON) {
		if err = json.NewDecoder(req.Body).Decode(i); err != nil && err != io.EOF {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid payload request")
			} else if _, ok := err.(*json.SyntaxError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid json request")
			} else if _, ok := err.(*time.ParseError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid datetime format")
			} else {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
		return
	}
	return echo.ErrUnsupportedMediaType
}
