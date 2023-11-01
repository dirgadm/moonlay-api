package http

import (
	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// CartHandler  represent the httphandler for Cart
type SubListsHandler struct {
	LUsecase domain.SubListsUsecase
}

// NewCartHandler will initialize the Cart resources endpoint
func NewSubListsHandler(e *echo.Echo, ps domain.SubListsUsecase) {
	handler := &SubListsHandler{
		LUsecase: ps,
	}
	v1 := e.Group("v1")
	v1.GET("/sublist", handler.GetList)
	v1.GET("/sublist/:id", handler.GetDetail)
}

func (h SubListsHandler) GetList(c echo.Context) (err error) {
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
	listId := ctx.GetParamInt("list_id")

	var subLists []dto.SubListsResponse
	var total int64
	subLists, total, err = h.LUsecase.GetList(ctx.Request().Context(), page.Start, page.Limit, search, listId)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.DataList(subLists, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h SubListsHandler) GetDetail(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		log.Error(err)
		return
	}

	var sublists dto.SubListsResponse
	sublists, err = h.LUsecase.GetDetail(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	ctx.Data(sublists)

	return ctx.Serve(err)
}
