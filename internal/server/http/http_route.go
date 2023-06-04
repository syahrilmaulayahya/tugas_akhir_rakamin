package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/infrastructure/container"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/server/http/handler"
)

func RouteInit(r *fiber.App, containerConf *container.Container) {
	// base url
	api := r.Group("/api/v1")

	// user endpoint
	handler.AuthRoute(api, containerConf)
	handler.UserRoute(api, containerConf)
	handler.TokoRoute(api, containerConf)
	handler.CategoryRoute(api, containerConf)
	handler.ProvinceCityRoute(api, containerConf)
	handler.ProdukRoute(api, containerConf)
	handler.TRXRoute(api, containerConf)
}
