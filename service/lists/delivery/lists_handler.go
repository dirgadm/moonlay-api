package http

import (
	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// CartHandler  represent the httphandler for Cart
type ListsHandler struct {
	LUsecase domain.ListsUsecase
}

// NewCartHandler will initialize the Cart resources endpoint
func NewListsHandler(e *echo.Echo, ps domain.ListsUsecase) {
	handler := &ListsHandler{
		LUsecase: ps,
	}
	v1 := e.Group("v1")
	v1.GET("/list", handler.GetList)
	v1.GET("/list/:id", handler.GetDetail)
}

func (h ListsHandler) GetList(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	// get pagination
	var page *ehttp.Paginator
	page, err = ehttp.NewPaginator(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	// get params
	search := ctx.GetParamString("search")

	var lists []dto.ListsResponse
	var total int64
	lists, total, err = h.LUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(lists, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h ListsHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var lists dto.ListsResponse
	lists, err = h.LUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(lists)

	return ctx.Serve(err)
}
