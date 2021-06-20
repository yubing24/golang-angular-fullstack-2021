package restcountry

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/yubing24/golang-angular-fullstack-2021/model"
	"net/http"
)

const endpoint = "https://restcountries.eu/rest/v2/all"

type RestCountryClient struct{}

func NewRestCountryClient() (RestCountryClient, error) {
	// if client needs to be authenticated, do it here
	return RestCountryClient{}, nil
}

type CurrencyDTO struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type LanguageDTO struct {
	ISO639_1   string `json:"iso639_1"`
	ISO639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type RestCountryDTO struct {
	Name              string   `json:"name"`
	TopLevelDomain    []string `json:"topLevelDomain"`
	Alpha2Code        string   `json:"alpha2Code"`
	Alpha3Code        string   `json:"alpha3Code"`
	CallingCode       []string
	Capital           string
	AltSpellings      []string
	Region            string
	SubRegion         string
	Population        int
	LatitudeLongitude []float32
	Demonym           string
	Area              float32
	Gini              float32
	Timezone          []string
	Borders           []string
	NativeName        string
	NumericCode       string
	Currencies        []CurrencyDTO
	Languages         []LanguageDTO
	Translations      map[string]string
	Flag              string
	RegionalBlocs     interface{}
	Cioc              string
}

func (dto RestCountryDTO) ToModel() model.Country {
	return model.Country{
		Name: dto.Name,
		Alpha2Code: dto.Alpha2Code,
		Alpha3Code: dto.Alpha3Code,
	}
}

func (client RestCountryClient) GetCountries(ctx context.Context) ([]model.Country, error) {
	var rawResponse = make([]byte, 0)
	countries := make([]model.Country, 0)
	if req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, bytes.NewBuffer(rawResponse)); err != nil {
		ctx.Done()
	} else {
		httpClient := &http.Client{}
		res, err := httpClient.Do(req)
		defer res.Body.Close()

		output := make([]RestCountryDTO, 0)
		err = json.NewDecoder(res.Body).Decode(&output)
		if err != nil {
			return make([]model.Country, 0), err
		}
		for _, each := range output {
			countries = append(countries, each.ToModel())
		}
	}
	return countries, nil
}
