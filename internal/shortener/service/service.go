package service

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/internal/shortener/model"
	"github.com/sog01/productdiscovery/internal/shortener/repository"
)

type Service interface {
	CreateShortener(ctx context.Context, req model.CreateShortenerReq) (model.CreateShortenerResp, error)
}

type ShortenerService struct {
	createShortenerRepo repository.CreateShortenerRepository
}

func (s ShortenerService) CreateShortener(ctx context.Context, req model.CreateShortenerReq) (model.CreateShortenerResp, error) {
	return CreateShortener(ctx, req, s.createShortenerRepo)
}

func NewService(ec *elastic.Client) Service {
	return ShortenerService{
		createShortenerRepo: repository.NewCreateShortenerRepository(ec),
	}
}
