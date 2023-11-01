package usecase

import (
	"context"
	"time"

	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type subListsUsecase struct {
	subListsRepo domain.SubListsRepository
	listsRepo    domain.ListsRepository
}

// NewListsUsecase will create new an articleUsecase object representation of domain.ListsUsecase interface
func NewSubListsUsecase(u domain.SubListsRepository, l domain.ListsRepository, timeout time.Duration) domain.SubListsUsecase {
	return &subListsUsecase{
		subListsRepo: u,
		listsRepo:    l,
	}
}

func (s *subListsUsecase) GetList(ctx context.Context, offset int, limit int, search string, listId int) (res []dto.SubListsResponse, total int64, err error) {
	var sublists []domain.SubLists
	sublists, total, err = s.subListsRepo.GetList(ctx, offset, limit, search, listId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, sl := range sublists {
		var list domain.Lists
		list, err = s.listsRepo.GetDetail(ctx, sl.ListId)
		if err != nil {
			log.Error(err)
			return
		}

		res = append(res, dto.SubListsResponse{
			Id:          sl.Id,
			Title:       sl.Title,
			ListId:      sl.ListId,
			Description: sl.Description,
			Priority:    sl.Priority,
			CreatedAt:   sl.CreatedAt,
			UpdatedAt:   sl.UpdatedAt,
			List: dto.ListsResponse{
				Id:          list.Id,
				Title:       list.Title,
				Description: list.Description,
				Priority:    list.Priority,
				CreatedAt:   list.CreatedAt,
				UpdatedAt:   list.UpdatedAt,
			},
		})
	}

	return
}

func (s *subListsUsecase) GetDetail(ctx context.Context, id int) (res dto.SubListsResponse, err error) {
	var subList domain.SubLists
	subList, err = s.subListsRepo.GetDetail(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	var list domain.Lists
	list, err = s.listsRepo.GetDetail(ctx, subList.ListId)
	if err != nil {
		log.Error(err)
		return
	}

	res = dto.SubListsResponse{
		Id:          subList.Id,
		Title:       subList.Title,
		ListId:      subList.ListId,
		Description: subList.Description,
		Priority:    subList.Priority,
		CreatedAt:   subList.CreatedAt,
		UpdatedAt:   subList.UpdatedAt,
		List: dto.ListsResponse{
			Id:          list.Id,
			Title:       list.Title,
			Description: list.Description,
			Priority:    list.Priority,
			CreatedAt:   list.CreatedAt,
			UpdatedAt:   list.UpdatedAt,
		},
	}

	return
}
