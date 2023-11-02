package http

import (
	"project-version3/moonlay-api/pkg/ehttp"
	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ListsHandler struct {
	LUsecase domain.ListsUsecase
}

func NewListsHandler(e *echo.Echo, ps domain.ListsUsecase) {
	handler := &ListsHandler{
		LUsecase: ps,
	}
	v1 := e.Group("v1")
	v1.GET("/list", handler.GetList)
	v1.GET("/list/:id", handler.GetDetail)
	v1.POST("/list", handler.Create)
	v1.PUT("/list/:id", handler.Update)
	v1.DELETE("/list/:id", handler.Delete)
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

func (h ListsHandler) Update(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.ListsRequest
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

func (h ListsHandler) Create(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)
	validator := validator.New()

	var req dto.ListsRequest
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

func (h ListsHandler) Delete(c echo.Context) (err error) {
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
