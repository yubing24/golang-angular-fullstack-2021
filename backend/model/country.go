package model

import (
	"context"
	"strings"
)

type Country struct {
	Name       string
	Alpha2Code string
	Alpha3Code string
}

type SearchCountryCriteria struct {
	Name      string
	MatchFull bool
	Code      string
}

func (criteria SearchCountryCriteria) IsEmpty() bool {
	return criteria.Name == "" && !criteria.MatchFull && criteria.Code == ""
}

func (country Country) Match(criteria SearchCountryCriteria) bool {
	if criteria.IsEmpty() {
		return true
	}
	if len(criteria.Name) > 0 {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(criteria.Name)) {
			return true
		}
		if criteria.MatchFull {
			if strings.ToLower(country.Name) == strings.ToLower(criteria.Name) {
				return true
			}
		}
	}
	if len(criteria.Code) > 0 {
		if strings.ToLower(country.Alpha3Code) == strings.ToLower(criteria.Code) ||
			strings.ToLower(country.Alpha2Code) == strings.ToLower(criteria.Code) {
			return true
		}
	}
	return false
}

type ICountryRepository interface {
	SearchCountries(ctx context.Context, criteria SearchCountryCriteria) ([]Country, error)
}
