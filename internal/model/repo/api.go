package repo

import "github.com/Cheng1622/news_go_server/internal/model"

type ApiTreeResp struct {
	ID       int          `json:"ID"`
	Desc     string       `json:"desc"`
	Category string       `json:"category"`
	Children []*model.Api `json:"children"`
}
