package service

import (
	"echo-sample/model"
	"echo-sample/model/database"
	"echo-sample/repository"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Musician interface {
	List(echo.Context, int, int) ([]model.Musician, int, error)
}

type musicianImpl struct {
	musicianRepository repository.Musician
}

func NewMusician(musicianRepository repository.Musician) Musician {
	return &musicianImpl{
		musicianRepository: musicianRepository,
	}
}

func (impl *musicianImpl) List(c echo.Context, perPage int, page int) ([]model.Musician, int, error) {
	count, err := impl.musicianRepository.Count(c)
	if err != nil {
		err = errors.Wrapf(err, "fail.Count")
		return nil, -1, err
	}

	list, err := impl.musicianRepository.List(c, perPage, page)
	if err != nil {
		err = errors.Wrapf(err, "fail.List")
		return nil, -1, err
	}

	return convertMusicianListModel(list), count, nil
}

func convertMusicianListModel(dbMusicianList []database.Musician) []model.Musician {
	m := []model.Musician{}
	for _, dbMusician := range dbMusicianList {
		m = append(m, model.Musician{
			ID:   dbMusician.ID,
			Name: dbMusician.Name,
		})
	}
	return m
}
