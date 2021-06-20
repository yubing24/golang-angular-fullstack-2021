package view

import "github.com/yubing24/golang-angular-fullstack-2021/model"

type CountryVM struct {
	Name string `json:"name"`
	Alpha2Code string `json:"alpha2Code"`
	Alpha3Code string `json:"alpha3Code"`
}

func NewCountryVM(country model.Country) CountryVM {
	return CountryVM{
		Name: country.Name,
		Alpha2Code: country.Alpha2Code,
		Alpha3Code: country.Alpha3Code,
	}
}