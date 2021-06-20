package view

import "github.com/yubing24/golang-angular-fullstack-2021/model"

type LanguageVM struct {
	Name string `json:"name"`
}

func NewLanguageVM(language model.Language) LanguageVM {
	return LanguageVM{
		Name: language.Name,
	}
}

type CountryVM struct {
	Name         string       `json:"name"`
	Alpha2Code   string       `json:"alpha2Code"`
	Alpha3Code   string       `json:"alpha3Code"`
	FlagImageUrl string       `json:"flagImageUrl"`
	Region       string       `json:"region"`
	SubRegion    string       `json:"subRegion"`
	Population   int          `json:"population"`
	Languages    []LanguageVM `json:"languages"`
}

func NewCountryVM(country model.Country) CountryVM {
	languages := make([]LanguageVM, 0)
	for _, each := range country.Languages {
		languages = append(languages, NewLanguageVM(each))
	}
	return CountryVM{
		Name:         country.Name,
		Alpha2Code:   country.Alpha2Code,
		Alpha3Code:   country.Alpha3Code,
		FlagImageUrl: country.FlagImageUrl,
		Region:       country.Region,
		SubRegion:    country.SubRegion,
		Population:   country.Population,
		Languages:    languages,
	}
}
