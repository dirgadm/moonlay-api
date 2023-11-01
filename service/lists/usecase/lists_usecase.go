package usecase

import (
	"context"
	"time"

	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type listsUsecase struct {
	listsRepo domain.ListsRepository
}

// NewListsUsecase will create new an articleUsecase object representation of domain.ListsUsecase interface
func NewListsUsecase(u domain.ListsRepository, timeout time.Duration) domain.ListsUsecase {
	return &listsUsecase{
		listsRepo: u,
	}
}

func (s *listsUsecase) GetList(ctx context.Context, offset int, limit int, search string) (res []dto.ListsResponse, total int64, err error) {
	var lists []domain.Lists
	lists, total, err = s.listsRepo.GetList(ctx, offset, limit, search)
	if err != nil {
		log.Error(err)
		return
	}

	for _, l := range lists {
		res = append(res, dto.ListsResponse{
			Id:          l.Id,
			Title:       l.Title,
			Description: l.Description,
			Priority:    l.Priority,
			CreatedAt:   l.CreatedAt,
			UpdatedAt:   l.UpdatedAt,
		})
	}

	return
}

func (s *listsUsecase) GetDetail(ctx context.Context, id int) (res dto.ListsResponse, err error) {
	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.ListsResponse{
		Id:          list.Id,
		Title:       list.Title,
		Description: list.Description,
		Priority:    list.Priority,
		CreatedAt:   list.CreatedAt,
		UpdatedAt:   list.UpdatedAt,
	}

	return
}
