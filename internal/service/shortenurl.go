package service

import (
	"context"

	"github.com/EBVN/bookmark-management/internal/repository"
)

const DefaultCodeLength = 7

//go:generate mockery --name=ShortenUrl --filename shortenurl.go --outpkg genpass_mocks
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

	// store url with code
	err = s.r.StoreURL(ctx, code, url)
	if err != nil {
		return "", err
	}

	// return code
	return code, nil
}

func (s *shortenUrl) GetUrlFromCode(ctx context.Context, code string) (string, error) {
	return s.r.GetURL(ctx, code)
}
