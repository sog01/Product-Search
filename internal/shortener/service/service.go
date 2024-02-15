package service

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/internal/shortener/model"
	"github.com/sog01/productdiscovery/internal/shortener/repository"
)

type Service interface {
	CreateShortener(ctx context.Context, req model.CreateShortenerReq) (model.CreateShortenerResp, error)
	GetShortener(ctx context.Context, req model.GetShortenerReq) (model.GetShortenerResp, error)
}

type ShortenerService struct {
	createShortenerRepo repository.CreateShortenerRepository
	getShortenerRepo    repository.GetShortenerRepository
}

func (s ShortenerService) CreateShortener(ctx context.Context, req model.CreateShortenerReq) (model.CreateShortenerResp, error) {
	return CreateShortener(ctx, req, s.createShortenerRepo)
}

func (s ShortenerService) GetShortener(ctx context.Context, req model.GetShortenerReq) (model.GetShortenerResp, error) {
	return Gethortener(ctx, req, s.getShortenerRepo)
}

func NewService(ec *elastic.Client) Service {
	return ShortenerService{
		createShortenerRepo: repository.NewCreateShortenerRepository(ec),
		getShortenerRepo:    repository.NewGetShortenerRepository(ec),
	}
}
