package api

import (
	"github.com/yubing24/golang-angular-fullstack-2021/model"
	"github.com/yubing24/golang-angular-fullstack-2021/repository"
	"github.com/yubing24/golang-angular-fullstack-2021/view"
	"net/http"
)

type searchCountryForm struct {
	Name string `schema:"name"`
	Code string `schema:"code"`
}

func (form searchCountryForm) ToCriteria() model.SearchCountryCriteria {
	return model.SearchCountryCriteria{
		Name: form.Name,
		Code: form.Code,
		MatchFull: false,
	}
}

func SearchCountry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo := repository.RepositoryImpl{}
	form := new(searchCountryForm)
	parseQueryForm(r, form)
	result, err := repo.SearchCountries(ctx, form.ToCriteria())
	if err != nil {
		respond(ctx, w, http.StatusInternalServerError, err, nil)
	} else {
		output := make([]view.CountryVM, 0)
		for _, each := range result {
			output = append(output, view.NewCountryVM(each))
		}
		w.Header().Set("Content-Type", "application/json")
		respond(ctx, w, http.StatusOK, nil, output)
	}
}
