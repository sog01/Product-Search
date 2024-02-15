package service

import (
	"context"
	"log"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/shortener/model"
	"github.com/sog01/productdiscovery/internal/shortener/repository"
)

func Gethortener(ctx context.Context, req model.GetShortenerReq, repo repository.GetShortenerRepository) (model.GetShortenerResp, error) {
	exec := pipe.PCtx(
		repo.GetShortener,
		composeGetShortenerResp,
	)
	resp, err := exec(ctx, req)
	if err != nil {
		log.Printf("failed to get shortener: %v", err)
	}

	r := pipe.Get[model.GetShortenerResp](resp)
	return r, nil
}

func composeGetShortenerResp(ctx context.Context, req model.GetShortenerReq, response pipe.Responses) (any, error) {
	resp := pipe.Get[map[string]any](response)
	realURL, _ := resp["real_url"].(string)
	return model.GetShortenerResp{
		RealURL: realURL,
	}, nil
}
