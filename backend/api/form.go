package api

import (
	"github.com/gorilla/schema"
	"github.com/yubing24/golang-angular-fullstack-2021/model"
	"net/http"
)

var queryParamDecoder = schema.NewDecoder()

func parseQueryForm(r *http.Request, form interface{}) error {
	if err := r.ParseForm(); err != nil {
		return r.Context().Err()
	} else {
		return queryParamDecoder.Decode(form, r.Form)
	}
}

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