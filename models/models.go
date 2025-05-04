package models

import "time"

type Country struct {
	Name       string
	States     int64
	Population int64
}

type NowHandler struct {
	Now func() time.Time
}

func NewNowHandler(abhi func() time.Time) NowHandler {
	return NowHandler{Now: abhi}
}

var Countries = []Country{
	{
		Name:       "india",
		States:     28,
		Population: 293892,
	},
	{
		Name:       "india1",
		States:     28,
		Population: 293892,
	},
	{
		Name:       "india2",
		States:     28,
		Population: 293892,
	},
	{
		Name:       "india3",
		States:     28,
		Population: 293892,
	},
}
