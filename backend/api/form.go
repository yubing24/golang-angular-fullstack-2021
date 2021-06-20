package api

import (
	"github.com/gorilla/schema"
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