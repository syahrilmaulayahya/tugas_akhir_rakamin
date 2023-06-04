package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/infrastructure/container"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/controller"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/repository"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/utils/apicall"
)

// AuthRoute group user endpoint
func AuthRoute(r fiber.Router, containerConf *container.Container) {

	repo := repository.NewUserRepository(containerConf.Mysqldb)
	userUseCase := usecase.NewUserUseCase(repo)
	middleware := usecase.NewMiddleware(usecase.Config{SharedKey: containerConf.Apps.SecretJwt})
	userController := controller.NewUserController(userUseCase, middleware)

	// auth endpoint
	AuthAPI := r.Group("/auth")
	AuthAPI.Post("/register", userController.Register)
	AuthAPI.Post("/login", userController.Login)

}

func TokoRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewTokoRepository(containerConf.Mysqldb)
	tokoUseCase := usecase.NewTokoUseCase(repo)
	tokoController := controller.NewTokoController(tokoUseCase)
	middleware := usecase.NewMiddleware(usecase.Config{SharedKey: containerConf.Apps.SecretJwt})
	auth := controller.NewAuthImpl(middleware)

	// toko endpoint
	tokoAPI := r.Group("/toko")
	tokoAPI.Get("/my", auth.CheckJwtUser, tokoController.GetMyToko)
	tokoAPI.Get("/:id_toko", auth.CheckJwtUser, tokoController.GetTokoByID)
	tokoAPI.Get("", auth.CheckJwtUser, tokoController.GetAllToko)
	tokoAPI.Put("/:id_toko", auth.CheckJwtUser, tokoController.UpdateToko)

}

func CategoryRoute(r fiber.Router, containerConf *container.Container) {
	// setup middleware service
	middleware := usecase.NewMiddleware(usecase.Config{
		SharedKey: containerConf.Apps.SecretJwt,
	})
	auth := controller.NewAuthImpl(middleware)

	// setup category service
	repo := repository.NewCategoryRepository(containerConf.Mysqldb)
	categoryUseCase := usecase.NewCategoryUseCase(repo)
	categoryController := controller.NewCategoryController(categoryUseCase)

	// category endpoint
	categoryAPI := r.Group("/category")
	categoryAPI.Post("", auth.CheckJwtAdmin, categoryController.CreateCategory)
	categoryAPI.Get("", categoryController.GetAllCategory)
	categoryAPI.Get("/:id", auth.CheckJwt, categoryController.GetCategoryByID)
	categoryAPI.Put("/:id", auth.CheckJwtAdmin, categoryController.UpdateCategoryByID)
	categoryAPI.Delete("/:id", auth.CheckJwtAdmin, categoryController.DeleteCategoryByID)
}

func UserRoute(r fiber.Router, containerConf *container.Container) {

	// setup middleware service
	middleware := usecase.NewMiddleware(usecase.Config{SharedKey: containerConf.Apps.SecretJwt})
	auth := controller.NewAuthImpl(middleware)

	// setup user service
	userRepo := repository.NewUserRepository(containerConf.Mysqldb)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUseCase, middleware)

	// setup alamat service
	alamatRepo := repository.NewAlamatRepository(containerConf.Mysqldb)
	alamatUseCae := usecase.NewAlamatUseCase(alamatRepo)
	alamatController := controller.NewAlamatController(alamatUseCae)

	// user endpoint
	userAPI := r.Group("/user")
	userAPI.Get("", auth.CheckJwtUser, userController.GetMyProfile)
	userAPI.Put("", auth.CheckJwtUser, userController.UpdateProfile)
	userAPI.Post("/alamat", auth.CheckJwtUser, alamatController.CreateAlamat)
	userAPI.Get("/alamat", auth.CheckJwtUser, alamatController.GetMyAlamat)
	userAPI.Get("/alamat/:id", auth.CheckJwtUser, alamatController.GetAlamatByID)
	userAPI.Put("/alamat/:id", auth.CheckJwtUser, alamatController.UpdateAlamatByID)
	userAPI.Delete("/alamat/:id", auth.CheckJwtUser, alamatController.DeleteAlamatByID)

}

// ProvinceCityRoute group provinceCity endpoint
func ProvinceCityRoute(r fiber.Router, containerConf *container.Container) {

	provinceCityApi := apicall.NewProvinceCityImpl(containerConf.Apps.URLPrvovinceCity)
	provinceCityController := controller.NewProvincCityController(provinceCityApi)

	// auth endpoint
	ProvCityApi := r.Group("/provcity")
	ProvCityApi.Get("/listprovincies", provinceCityController.GetAllProvince)
	ProvCityApi.Get("/listcities/:prov_id", provinceCityController.GetListCities)
	ProvCityApi.Get("/detailprovince/:prov_id", provinceCityController.GetDetailProvince)

}

func ProdukRoute(r fiber.Router, containerConf *container.Container) {
	// setup middleware service
	middleware := usecase.NewMiddleware(usecase.Config{SharedKey: containerConf.Apps.SecretJwt})
	auth := controller.NewAuthImpl(middleware)

	produkRepo := repository.NewProdukRepository(containerConf.Mysqldb)
	produkUseCase := usecase.NewProdukUseCase(produkRepo)
	produkController := controller.NewProdukController(produkUseCase)

	produkAPI := r.Group("/product")
	produkAPI.Post("", auth.CheckJwtUser, produkController.UploadProduk)
	produkAPI.Get("/:id", produkController.GetProdukByID)
	produkAPI.Put("/:id", auth.CheckJwtUser, produkController.UpdateProdukByID)
	produkAPI.Delete("/:id", auth.CheckJwtUser, produkController.DeleteProdukByID)
	produkAPI.Get("", produkController.GetAllProduk)

}

func TRXRoute(r fiber.Router, containerConf *container.Container) {
	// setup middleware service
	middleware := usecase.NewMiddleware(usecase.Config{SharedKey: containerConf.Apps.SecretJwt})
	auth := controller.NewAuthImpl(middleware)

	trxRepo := repository.NewTRXRepository(containerConf.Mysqldb)
	trxUseCase := usecase.NewTRXUseCase(trxRepo)
	trxController := controller.NewTRXController(trxUseCase)

	trxAPI := r.Group("/trx")
	trxAPI.Post("", auth.CheckJwtUser, trxController.CreateTRX)
	trxAPI.Get("/", auth.CheckJwtUser, trxController.GetALlTRX)
	trxAPI.Get("/:id", auth.CheckJwtUser, trxController.GetTRXByID)

}
