package api

import (
	"errors"
	"github.com/yubing24/golang-angular-fullstack-2021/model"
	"github.com/yubing24/golang-angular-fullstack-2021/repository"
	"github.com/yubing24/golang-angular-fullstack-2021/view"
	"net/http"
	"sort"
)



func SearchCountry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo := repository.RepositoryImpl{}
	form := new(searchCountryForm)
	parseQueryForm(r, form)
	result, err := repo.SearchCountries(ctx, form.ToCriteria())
	if err != nil {
		respond(ctx, w, http.StatusInternalServerError, err, nil)
	} else {

		if result == nil || len(result) == 0 {
			respond(ctx, w, http.StatusNotFound, errors.New("no country matches the criteria"), nil)
			return
		}

		records := make([]view.CountryVM, 0)
		sort.Sort(sort.Reverse(model.ByPopulation(result)))
		for _, each := range result {
			records = append(records, view.NewCountryVM(each))
		}

		statistics := model.GetSearchCountryResultStatistics(result)

		w.Header().Set("Content-Type", "application/json")
		respond(ctx, w, http.StatusOK, nil, map[string]interface{} {
			"records":    records,
			"statistics": statistics,
		})
	}
}
