package dto

import (
	"time"
)

type SubListsResponse struct {
	Id          int           `json:"id"`
	ListId      int           `json:"list_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Priority    int           `json:"priority"`
	List        ListsResponse `json:"list"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
}

type SubListsRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Files       []string `json:"files"`
}
