package http

import (
	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// CartHandler  represent the httphandler for Cart
type UploadsHandler struct {
	UUsecase domain.UploadUsecase
}

// NewCartHandler will initialize the Cart resources endpoint
func NewUploadsHandler(e *echo.Echo, ps domain.UploadUsecase) {
	handler := &UploadsHandler{
		UUsecase: ps,
	}
	v1 := e.Group("v1")
	v1.POST("/upload", handler.UploadFile)
}

func (h UploadsHandler) UploadFile(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var uploads []dto.UploadResponse
	uploads, err = h.UUsecase.UploadFile(ctx.Request().Context(), c.Response(), c.Request())
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(uploads, 0, 0, 0)

	return ctx.Serve(err)
}
