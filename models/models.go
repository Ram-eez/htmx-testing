package models

import (
	"errors"
	"time"
)

type Country struct {
	Name       string
	States     int64
	Population int64
}

type CountryListViewModel struct {
	Countries    []Country
	ErrorMessage error
}

type NowHandler struct {
	Now func() time.Time
}

func NewCountryListViewModel(countries []Country, err error) CountryListViewModel {
	var vm CountryListViewModel
	new_err := errors.New("error to load the countries")
	if err != nil {
		vm.ErrorMessage = new_err
		return vm
	}
	vm.Countries = countries
	return vm
}
func NewNowHandler(abhi func() time.Time) NowHandler {
	return NowHandler{Now: abhi}
}

func GetCountries(msg string) ([]Country, error) {
	err := errors.New("could not fetch contries")
	if msg != "country" {
		return nil, err
	}
	return Countries, nil

}

type ContextKey string

var ThemeContextKey ContextKey = "theme"

var Countries = []Country{
	{
		Name:       "india",
		States:     1,
		Population: 293892,
	},
	{
		Name:       "india1",
		States:     2,
		Population: 293892,
	},
	{
		Name:       "india2",
		States:     3,
		Population: 293892,
	},
	{
		Name:       "india3",
		States:     4,
		Population: 293892,
	},
}
