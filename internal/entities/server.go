package entities

import (
	"context"
	api "ozon-test/pkg/api/proto"
)

type Database interface {
	AddURL(context.Context, *api.AddURLRequest) (*api.AddURLResponse, error)
	GetURL(context.Context, *api.GetURLRequest) (*api.GetURLResponse, error)
	GetStorageType() string
}

type Link struct {
	OriginalLink string
	ShortLink    string
}
