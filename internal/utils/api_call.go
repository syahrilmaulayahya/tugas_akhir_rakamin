package utils

import (
	"encoding/json"
	"fmt"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"io/ioutil"
	"net/http"
)

const (
	URL             = "https://emsifa.github.io/api-wilayah-indonesia/api/"
	currentfilepath = "internal/infrastructure/utils/api_call.go"
)

// GetProvince get all province data from 3rd party api
func GetAllProvince() ([]dto.Province, error) {

	// format url to get province data
	urlProvince := fmt.Sprintf("%sprovinces.json", URL)
	result := []dto.Province{}

	// call http get method to get the data
	response, err := http.Get(urlProvince)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to get data : %s", err.Error()))
		return result, err
	}

	// read response from http get method
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to get read data : %s", err.Error()))
		return result, err
	}

	// mapping response from 3rd party api to local variable
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to unmarshal data : %s", err.Error()))
		return result, err
	}
	return result, err
}

// GetProvince get province data with id as parameter from 3rd party api
func GetProvince(id string) (dto.Province, error) {

	// format url to get province data
	urlProvince := fmt.Sprintf("%sprovince/%s.json", URL, id)
	result := dto.Province{}

	// call http get method to get the data
	response, err := http.Get(urlProvince)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to get data : %s", err.Error()))
		return result, err
	}

	// read response from http get method
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to get read data : %s", err.Error()))
		return result, err
	}

	// mapping response from 3rd party api to local variable
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to unmarshal data : %s", err.Error()))
		return result, err
	}
	return result, err
}

// GetRegency get regency data with id as parameter from 3rd party api
func GetRegency(id string) (dto.Regency, error) {

	// format url to get province data
	urlRegency := fmt.Sprintf("%sregency/%s.json", URL, id)
	result := dto.Regency{}

	// call http get method to get the data
	response, err := http.Get(urlRegency)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to get data : %s", err.Error()))
		return result, err
	}

	// read response from http get method
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to get read data : %s", err.Error()))
		return result, err
	}

	// mapping response from 3rd party api to local variable
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed to unmarshal data : %s", err.Error()))
		return result, err
	}
	return result, err
}
