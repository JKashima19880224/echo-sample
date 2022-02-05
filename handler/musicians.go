package handler

import (
	"echo-sample/model"
	"echo-sample/model/request"
	"echo-sample/model/response"
	"echo-sample/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Musicians interface {
	List(echo.Context) error
}

type musiciansImpl struct {
	musicianService service.Musician
}

func NewMusicians(musicianService service.Musician) Musicians {
	return &musiciansImpl{
		musicianService: musicianService,
	}
}

func (impl *musiciansImpl) List(c echo.Context) error {
	defaultPerPage, defaultPage := 10, 1
	req := &request.MusicianList{
		PerPage: defaultPerPage,
		Page:    defaultPage,
	}

	musicianList, count, err := impl.musicianService.List(c, req.PerPage, req.Page)
	if err != nil {
		err = errors.Wrapf(err, "fail.List")
		return err
	}

	res := convertMusicianListResponse(count, musicianList)
	return c.JSON(http.StatusOK, res)
}

func convertMusicianListResponse(count int, modelMusicianList []model.Musician) response.MusicianList {
	m := []response.Musician{}
	for _, modelMusician := range modelMusicianList {
		m = append(m, response.Musician{
			ID:   modelMusician.ID,
			Name: modelMusician.Name,
		})
	}

	return response.MusicianList{
		TotalCount: count,
		Items:      m,
	}
}
