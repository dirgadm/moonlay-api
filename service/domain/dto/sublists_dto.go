package dto

import (
	"time"
)

type SubListsResponse struct {
	Id          int              `json:"id"`
	ListId      int              `json:"list_id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Priority    int              `json:"priority"`
	List        ListsResponse    `json:"list,omitempty"`
	Files       []UploadResponse `json:"files,omitempty"`
	CreatedAt   time.Time        `json:"created_at,omitempty"`
	UpdatedAt   time.Time        `json:"updated_at,omitempty"`
}

type SubListsRequest struct {
	Id          int      `json:"id"`
	ListId      int      `json:"list_id"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Files       []string `json:"files"`
}
