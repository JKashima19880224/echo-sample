package repository

import (
	"echo-sample/model/database"

	"github.com/labstack/echo/v4"
)

type Musician interface {
	List(echo.Context, int, int) ([]database.Musician, error)
	Count(echo.Context) (int, error)
}

type musicianImpl struct {
}

func NewMusician() Musician {
	return &musicianImpl{}
}

func (*musicianImpl) List(c echo.Context, _ int, _ int) ([]database.Musician, error) {
	musicianList := []database.Musician{
		{
			ID:   1,
			Name: "Miles Davis",
		},
		{
			ID:   2,
			Name: "John Coltrane",
		},
		{
			ID:   3,
			Name: "Thelonious Monk",
		},
		{
			ID:   4,
			Name: "Dizzy Gillespie",
		},
		{
			ID:   5,
			Name: "Charlie Parker",
		},
	}
	return musicianList, nil
}

func (*musicianImpl) Count(c echo.Context) (int, error) {
	return 5, nil
}
