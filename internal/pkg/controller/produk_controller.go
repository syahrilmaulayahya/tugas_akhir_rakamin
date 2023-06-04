package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/usecase"
	"strconv"
	"strings"
)

type ProdukController interface {
	UploadProduk(ctx *fiber.Ctx) (err error)
	GetProdukByID(ctx *fiber.Ctx) (err error)
	UpdateProdukByID(ctx *fiber.Ctx) (err error)
	DeleteProdukByID(ctx *fiber.Ctx) (err error)
	GetAllProduk(ctx *fiber.Ctx) (err error)
}

type ProdukControllerImpl struct {
	produkUseCase usecase.ProdukUseCase
}

func NewProdukController(produkUseCase usecase.ProdukUseCase) ProdukController {
	return &ProdukControllerImpl{produkUseCase: produkUseCase}
}

func (pc *ProdukControllerImpl) UploadProduk(ctx *fiber.Ctx) (err error) {
	// get tokoID (tokoID is the same as userID) from middleware
	tokoIDMiddleware := ctx.Locals("userID")
	tokoID, _ := strconv.Atoi(fmt.Sprintf("%v", tokoIDMiddleware))

	// get form value
	namaProduk := ctx.FormValue("nama_produk")
	categoryId, errConv := strconv.Atoi(ctx.FormValue("category_id"))
	if errConv != nil {
		response := BaseResponse{
			Status:  false,
			Message: "category_id must be number > 0 ",
			Error:   []string{errConv.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	hargaReseller, errConv := strconv.Atoi(ctx.FormValue("harga_reseller"))
	if errConv != nil {
		response := BaseResponse{
			Status:  false,
			Message: "harga_reseller must be number > 0 ",
			Error:   []string{errConv.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	hargaKonsumen, errConv := strconv.Atoi(ctx.FormValue("harga_konsumen"))
	if errConv != nil {
		response := BaseResponse{
			Status:  false,
			Message: "harga_konsumen must be number > 0 ",
			Error:   []string{errConv.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	stok, errConv := strconv.Atoi(ctx.FormValue("stok"))
	if errConv != nil {
		response := BaseResponse{
			Status:  false,
			Message: "stok must be number > 0 ",
			Error:   []string{errConv.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	deskripsi := ctx.FormValue("deskripsi")

	// map form-data value to local struct
	var data dto.UploadProdukRequest
	data = dto.UploadProdukRequest{
		NamaProduk:    namaProduk,
		CategoryID:    uint(categoryId),
		TokoID:        uint(tokoID),
		HargaReseller: uint(hargaReseller),
		HargaKonsumen: uint(hargaKonsumen),
		Stok:          uint(stok),
		Deskripsi:     deskripsi,
		Photos:        nil,
	}

	// initiate multiplatform to get data from form-data file
	form, err := ctx.MultipartForm()
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	files := form.File["photos"]

	// saving file input to local
	for i, file := range files {
		if file != nil {
			var filename = file.Filename

			filename = fmt.Sprintf("%d-%s-%d-%v", tokoID, namaProduk, i, filename)
			filename = strings.ToLower(filename)
			filenameSplit := strings.Split(filename, " ")
			filename = strings.Join(filenameSplit, "_")
			if errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/toko/%s", filename)); errSaveFile != nil {
				response := BaseResponse{
					Status:  false,
					Message: "Failed to POST data",
					Error:   []string{errSaveFile.Error()},
					Data:    nil,
				}
				return ctx.Status(fiber.StatusInternalServerError).JSON(response)
			}
			photo := dto.Photos{URL: fmt.Sprintf("./public/images/toko/%s", filename)}
			data.Photos = append(data.Photos, photo)
		}
	}
	// call UploadProduk from produk useCase
	c := ctx.Context()
	responseUseCase, errUseCase := pc.produkUseCase.UploadProduk(c, data)
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    responseUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pc *ProdukControllerImpl) GetProdukByID(ctx *fiber.Ctx) (err error) {
	// get id from url parameter
	GetParams := ctx.Params("id")

	// check if id valid
	ID, err := strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetProdukByID from produk useCase
	c := ctx.Context()
	responseUseCase, errUseCase := pc.produkUseCase.GetProdukByID(c, uint(ID))
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data:    responseUseCase,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pc *ProdukControllerImpl) UpdateProdukByID(ctx *fiber.Ctx) (err error) {
	// get tokoID (tokoID is the same as userID) from middleware
	tokoIDMiddleware := ctx.Locals("userID")
	tokoID, _ := strconv.Atoi(fmt.Sprintf("%v", tokoIDMiddleware))

	// get id from url parameter
	GetParams := ctx.Params("id")
	// check if id is valid
	ID, err := strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	var data dto.UpdateProdukRequest

	// get form value
	namaProduk := ctx.FormValue("nama_produk")
	var categoryID int
	var errConv error
	if ctx.FormValue("category_id") != "" {
		categoryID, errConv = strconv.Atoi(ctx.FormValue("category_id"))
		if errConv != nil {
			response := BaseResponse{
				Status:  false,
				Message: "category_id must be number > 0 ",
				Error:   []string{errConv.Error()},
				Data:    nil,
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}
	}
	var hargaReseller int
	if ctx.FormValue("harga_reseller") != "" {
		hargaReseller, errConv = strconv.Atoi(ctx.FormValue("harga_reseller"))
		if errConv != nil {
			response := BaseResponse{
				Status:  false,
				Message: "harga_reseller must be number > 0 ",
				Error:   []string{errConv.Error()},
				Data:    nil,
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}
	}
	var hargaKonsumen int
	if ctx.FormValue("harga_konsumen") != "" {
		hargaKonsumen, errConv = strconv.Atoi(ctx.FormValue("harga_konsumen"))
		if errConv != nil {
			response := BaseResponse{
				Status:  false,
				Message: "harga_konsumen must be number > 0 ",
				Error:   []string{errConv.Error()},
				Data:    nil,
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}
	}
	var stok int
	if ctx.FormValue("stok") != "" {
		stok, errConv = strconv.Atoi(ctx.FormValue("stok"))
		if errConv != nil {
			response := BaseResponse{
				Status:  false,
				Message: "stok must be number > 0 ",
				Error:   []string{errConv.Error()},
				Data:    nil,
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(response)
		}
	}

	deskripsi := ctx.FormValue("deskripsi")

	// map form-data value to local struct
	data = dto.UpdateProdukRequest{
		ID:            uint(ID),
		NamaProduk:    namaProduk,
		CategoryID:    uint(categoryID),
		TokoID:        uint(tokoID),
		HargaReseller: uint(hargaReseller),
		HargaKonsumen: uint(hargaKonsumen),
		Stok:          uint(stok),
		Deskripsi:     deskripsi,
		Photos:        nil,
	}

	// initiate multiplatform to get data from form-data file
	form, err := ctx.MultipartForm()
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// get files from form-data with key photos
	files := form.File["photos"]

	for i, file := range files {

		if file != nil {
			var filename = file.Filename
			filename = fmt.Sprintf("%d-%s-%d-%v", tokoID, namaProduk, i, filename)
			filename = strings.ToLower(filename)
			filenameSplit := strings.Split(filename, " ")
			filename = strings.Join(filenameSplit, "_")
			// save file to local folder with specified name
			if errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/toko/%s", filename)); errSaveFile != nil {
				response := BaseResponse{
					Status:  false,
					Message: "Failed to POST data",
					Error:   []string{errSaveFile.Error()},
					Data:    nil,
				}
				return ctx.Status(fiber.StatusInternalServerError).JSON(response)
			}
			// add photo to local struct
			photo := dto.Photos{URL: fmt.Sprintf("./public/images/toko/%s", filename)}
			data.Photos = append(data.Photos, photo)
		}
	}
	// call UploadProduk from produk useCase to update produk record with specified toko_id and id
	c := ctx.Context()
	errUseCase := pc.produkUseCase.UpdateProdukByID(c, data)
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to POST data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Error:   nil,
		Data:    "",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pc *ProdukControllerImpl) DeleteProdukByID(ctx *fiber.Ctx) (err error) {
	// get tokoID (tokoID is the same as userID) from middleware
	tokoIDMiddleware := ctx.Locals("userID")
	tokoID, _ := strconv.Atoi(fmt.Sprintf("%v", tokoIDMiddleware))

	// get id from url parameter
	GetParams := ctx.Params("id")
	// check if id is valid
	ID, err := strconv.Atoi(GetParams)
	if err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "ID must integer > 0",
			Error:   []string{err.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	// call deleteTokoByID from produk useCase
	c := ctx.Context()
	errUseCase := pc.produkUseCase.DeleteProdukByID(c, uint(tokoID), uint(ID))
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to DELETE data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to DELETE data",
		Error:   nil,
		Data:    "",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pc *ProdukControllerImpl) GetAllProduk(ctx *fiber.Ctx) (err error) {
	// get limit and page from query parameter url
	filter := new(dto.FilterProduk)
	if errQuery := ctx.QueryParser(filter); errQuery != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errQuery.Error()},
			Data:    nil,
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	// call GetAllProduk from produk useCase to get produk records
	c := ctx.Context()
	responseUseCase, errUseCase := pc.produkUseCase.GetAllProduk(c, *filter)
	// error checking
	if errUseCase.Err != nil {
		response := BaseResponse{
			Status:  false,
			Message: "Failed to GET data",
			Error:   []string{errUseCase.Err.Error()},
			Data:    nil,
		}
		return ctx.Status(errUseCase.Code).JSON(response)
	}
	type DataProduct struct {
		Data []dto.GetProduk `json:"data"`
	}
	// success response
	response := BaseResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Error:   nil,
		Data: DataProduct{
			Data: responseUseCase,
		},
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
