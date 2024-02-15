package service

import (
	"context"
	"log"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/shortener/model"
	"github.com/sog01/productdiscovery/internal/shortener/repository"
)

func CreateShortener(ctx context.Context, req model.CreateShortenerReq, repo repository.CreateShortenerRepository) (model.CreateShortenerResp, error) {
	exec := pipe.PCtx(
		repo.CreateShortener,
		composeShortenerResp,
	)
	resp, err := exec(ctx, req)
	if err != nil {
		log.Printf("failed to create shortener: %v", err)
	}

	r := pipe.Get[model.CreateShortenerResp](resp)
	return r, nil
}

func composeShortenerResp(ctx context.Context, req model.CreateShortenerReq, response pipe.Responses) (any, error) {
	resp := pipe.Get[map[string]any](response)
	slug, _ := resp["slug"].(string)
	return model.CreateShortenerResp{
		Slug: slug,
	}, nil
}
