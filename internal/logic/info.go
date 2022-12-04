package logic

import (
	"context"

	"github.com/pgillich/meals-demo/internal/api"
	"github.com/pgillich/meals-demo/internal/buildinfo"
)

func (fs *FoodStore) GetLivez(ctx context.Context, request api.GetLivezRequestObject) (api.GetLivezResponseObject, error) {
	return &api.GetLivez200Response{}, nil
}

func (fs *FoodStore) GetVersion(ctx context.Context, request api.GetVersionRequestObject) (api.GetVersionResponseObject, error) {
	return &api.GetVersion200JSONResponse{
		AppName:   buildinfo.AppName,
		BuildTime: buildinfo.BuildTime,
		Version:   buildinfo.Version,
		GoMod:     buildinfo.GoMod,
	}, nil
}
