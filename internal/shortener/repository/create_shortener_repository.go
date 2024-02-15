package repository

import (
	"context"
	"time"

	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/shortener/model"
	sid "github.com/teris-io/shortid"
)

type CreateShortenerRepository struct {
	CreateShortener pipe.FuncCtx[model.CreateShortenerReq]
}

func NewCreateShortenerRepository(ec *elastic.Client) CreateShortenerRepository {
	return CreateShortenerRepository{
		CreateShortener: func(ctx context.Context, args model.CreateShortenerReq, responses pipe.Responses) (response any, err error) {
			id := uuid.NewV4().String()
			slug := sid.MustGenerate()
			_, err = ec.Index().Index("product_search_shortener").BodyJson(map[string]any{
				"id":         id,
				"slug":       slug,
				"real_url":   args.RealURL,
				"created_at": time.Now().UTC(),
				"updated_at": time.Now().UTC(),
			}).Do(ctx)
			return map[string]any{
				"slug": slug,
			}, err
		},
	}
}
