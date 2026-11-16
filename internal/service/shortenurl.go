package service

import (
	"context"
	"errors"

	"github.com/EBVN/bookmark-management/internal/repository"
)

const DefaultCodeLength = 7

//go:generate mockery --name=ShortenUrl --filename shorten.go --outpkg genpass_svc_mocks
type ShortenUrl interface {
	CreateShortenLink(ctx context.Context, url string) (string, error)
	GetUrlFromCode(ctx context.Context, code string) (string, error)
}

type shortenUrl struct {
	r       repository.URLStorage
	codeGen GenPass
}

func NewShortenUrl(r repository.URLStorage, codeGen GenPass) ShortenUrl {
	return &shortenUrl{r: r, codeGen: codeGen}
}

func (s *shortenUrl) CreateShortenLink(ctx context.Context, url string) (string, error) {
	// generate code
	code, err := s.codeGen.GeneratePassword(DefaultCodeLength)
	if err != nil {
		return "", err
	}

	// store url theo code
	err = s.r.StoreURL(ctx, code, url)
	if err != nil {
		return "", err
	}

	// return code
	return code, nil
}

var ErrUrlDoesNotExist = errors.New("url does not exist!")

func (s *shortenUrl) GetUrlFromCode(ctx context.Context, code string) (string, error) {
	res, err := s.r.GetURL(ctx, code)
	if errors.Is(err, repository.ErrUrlDoesNotExist) {
		return "", ErrUrlDoesNotExist
	}

	return res, err
}
