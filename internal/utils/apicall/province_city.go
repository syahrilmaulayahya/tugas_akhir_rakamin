package apicall

import (
	"encoding/json"
	"fmt"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"io/ioutil"
	"net/http"
)

type ProvinceCity interface {
	GetAllProvince() ([]dto.Province, error)
	GetListCity(provinceID uint) ([]dto.Regency, error)
	GetDetailProvince(provinceID uint) (dto.Province, error)
}

type ProvinceCityImpl struct {
	URL string
}

func NewProvinceCityImpl(URL string) ProvinceCity {
	return &ProvinceCityImpl{
		URL: URL,
	}
}

func (pci *ProvinceCityImpl) GetAllProvince() ([]dto.Province, error) {

	// format url to get province data
	URLProvince := fmt.Sprintf("%sprovinces.json", pci.URL)

	// var result to save result from api call
	var result []dto.Province

	// call http get method to get the data
	response, err := http.Get(URLProvince)
	if err != nil {
		return result, err
	}

	// read response from http get method
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	// mapping response from 3rd party api to local variable
	err = json.Unmarshal(responseData, &result)
	if err != nil {

		return result, err
	}
	return result, err
}

func (pci *ProvinceCityImpl) GetListCity(provinceID uint) ([]dto.Regency, error) {

	// format url to get regency data
	URLProvince := fmt.Sprintf("%s/regencies/%d.json", pci.URL, provinceID)

	// var result to save result from api call
	var result []dto.Regency

	// call http get method to get the data
	response, err := http.Get(URLProvince)
	if err != nil {
		return result, err
	}

	// read response from http get method
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	// mapping response from 3rd party api to local variable
	err = json.Unmarshal(responseData, &result)
	if err != nil {

		return result, err
	}
	return result, err
}

func (pci *ProvinceCityImpl) GetDetailProvince(provinceID uint) (dto.Province, error) {

	// format url to get regency data
	URLProvince := fmt.Sprintf("%s/province/%d.json", pci.URL, provinceID)

	// var result to save result from api call
	var result dto.Province

	// call http get method to get the data
	response, err := http.Get(URLProvince)
	if err != nil {
		return result, err
	}

	// read response from http get method
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	// mapping response from 3rd party api to local variable
	err = json.Unmarshal(responseData, &result)
	if err != nil {

		return result, err
	}
	return result, err
}
