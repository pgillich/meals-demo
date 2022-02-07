package logic

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/buildinfo"
	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi/operations"
	"github.com/pgillich/meals-demo/internal/restapi/operations/info"
)

func SetInfoAPI(config configs.Options, api *operations.OpenAPIFoodstoreAPI) {
	api.InfoGetLivezHandler = info.GetLivezHandlerFunc(GetLivez)
	api.InfoGetVersionHandler = info.GetVersionHandlerFunc(GetVersion)
}

func GetLivez(params info.GetLivezParams) middleware.Responder {
	return info.NewGetLivezOK()
}

func GetVersion(params info.GetVersionParams) middleware.Responder {
	return info.NewGetVersionOK().WithPayload(&models.Version{
		AppName:   buildinfo.AppName,
		BuildTime: buildinfo.BuildTime,
		Version:   buildinfo.Version,
		GoMod:     buildinfo.GoMod,
	})
}
