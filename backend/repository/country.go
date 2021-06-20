package repository

import (
	"context"
	"errors"
	"github.com/yubing24/golang-angular-fullstack-2021/integration/restcountry"
	"github.com/yubing24/golang-angular-fullstack-2021/model"
)

type RepositoryImpl struct {}

func (repo RepositoryImpl) SearchCountries(ctx context.Context, criteria model.SearchCountryCriteria) ([]model.Country, error) {
	var err error
	client, err := restcountry.NewRestCountryClient()
	if err != nil {
		ctx.Done()
	}
	if result, err := client.GetCountries(ctx); err != nil {
		ctx.Done()
	} else {
		output := make([]model.Country, 0)
		for _, each := range result {
			if each.Match(criteria) {
				output = append(output, each)
			}
		}
		return output, nil
	}
	err = errors.New("an unexpected error occurred")
	return make([]model.Country, 0), err
}