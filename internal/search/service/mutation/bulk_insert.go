package mutation

import (
	"context"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

func BulkInsert(ctx context.Context, req model.BulkInsertReq, repo repository.BulkInsertRepository) (model.BulkInsertResp, error) {
	exec := pipe.PCtx(
		repo.BulkInsert,
	)

	_, err := exec(ctx, req)
	if err != nil {
		return model.BulkInsertResp{}, err
	}

	return model.BulkInsertResp{}, nil
}
