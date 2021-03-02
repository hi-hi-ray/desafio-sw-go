package integration

import "github.com/hi-hi-ray/desafio-sw-go/api/models"

var residents = []string{
	"http://swapi.dev/api/people/1/",
	"http://swapi.dev/api/people/2/",
	"http://swapi.dev/api/people/4/",
	"http://swapi.dev/api/people/6/",
	"http://swapi.dev/api/people/7/",
	"http://swapi.dev/api/people/8/",
	"http://swapi.dev/api/people/9/",
	"http://swapi.dev/api/people/11/",
	"http://swapi.dev/api/people/43/",
	"http://swapi.dev/api/people/62/",
}

var films = []string{
	"http://swapi.dev/api/films/1/",
	"http://swapi.dev/api/films/3/",
	"http://swapi.dev/api/films/4/",
	"http://swapi.dev/api/films/5/",
	"http://swapi.dev/api/films/6/",
}

var Result = models.Results{
	Name:           "Tatooine",
	RotationPeriod: "23",
	OrbitalPeriod:  "304",
	Diameter:       "10465",
	Climate:        "arid",
	Gravity:        "1 standard",
	Terrain:        "desert",
	SurfaceWater:   "1",
	Population:     "200000",
	ResidentURLs:   residents,
	FilmURLs:       films,
	Created:        "2014-12-09T13:50:49.641000Z",
	Edited:         "2014-12-20T20:58:18.411000Z",
	URL:            "http://swapi.dev/api/planets/1/",
}
