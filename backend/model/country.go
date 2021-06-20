package model

import (
	"context"
	"strings"
)

type Language struct {
	Name           string
	NativeName     string
	Iso_639_1_Name string
	Iso_639_2_Name string
}

type Country struct {
	Name         string
	Alpha2Code   string
	Alpha3Code   string
	Population   int
	FlagImageUrl string
	Region       string
	SubRegion    string
	Languages    []Language
}

type ByPopulation []Country

func (p ByPopulation) Len() int {
	return len(p)
}

func (p ByPopulation) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPopulation) Less(i, j int) bool {
	return p[i].Population < p[j].Population
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
		if strings.Contains(strings.ToLower(country.Alpha3Code), strings.ToLower(criteria.Code)) ||
			strings.Contains(strings.ToLower(country.Alpha2Code), strings.ToLower(criteria.Code)) {
			return true
		}
	}
	return false
}

type SearchCountryResultStatistics struct {
	TotalCountries    int
	RegionCounters    map[string]int
	SubRegionCounters map[string]int
}

func GetSearchCountryResultStatistics(result []Country) SearchCountryResultStatistics {
	output := SearchCountryResultStatistics{
		TotalCountries:    len(result),
		RegionCounters:    make(map[string]int),
		SubRegionCounters: make(map[string]int),
	}

	for _, each := range result {
		output.RegionCounters[each.Region] = output.RegionCounters[each.Region] + 1
		output.SubRegionCounters[each.SubRegion] = output.SubRegionCounters[each.SubRegion] + 1
	}

	if output.RegionCounters[""] != 0 {
		output.RegionCounters["Other"] = output.RegionCounters[""]
		delete(output.RegionCounters, "")
	}
	if output.SubRegionCounters[""] != 0 {
		output.SubRegionCounters["Other"] = output.SubRegionCounters[""]
		delete(output.SubRegionCounters, "")
	}

	return output
}

type ICountryRepository interface {
	SearchCountries(ctx context.Context, criteria SearchCountryCriteria) ([]Country, error)
}
