package http

import (
	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/go-playground/validator/v10"
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
	v1.POST("/sublist", handler.Create)
	v1.PUT("/sublist/:id", handler.Update)
	v1.DELETE("/sublist/:id", handler.Delete)
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

func (h SubListsHandler) Update(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.SubListsRequest
	if req.Id, err = ctx.GetParamUri("id"); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = ctx.Bind(&req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = validator.Struct(req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	err = h.LUsecase.Update(ctx.Request().Context(), req)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}

func (h SubListsHandler) Create(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.SubListsRequest
	if err = ctx.Bind(&req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	if err = validator.Struct(req); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	err = h.LUsecase.Create(ctx.Request().Context(), req)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}

func (h SubListsHandler) Delete(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	if id, err = ctx.GetParamUri("id"); err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}
	err = h.LUsecase.Delete(ctx.Request().Context(), id)
	if err != nil {
		log.Error(err)
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}
