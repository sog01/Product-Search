package search

import (
	"context"
	"fmt"

	"github.com/sog01/pipe"
)

type BulkUpdateRepository struct {
	BulkUpdate pipe.Func[BulkUpdateReq]
}

func BulkUpdate(ctx context.Context, req BulkUpdateReq, repo BulkUpdateRepository) (BulkUpdateResp, error) {
	exec := pipe.P(
		validateBulkUpdateRequest,
		repo.BulkUpdate,
	)

	req.ctx = ctx
	_, err := exec(req)
	if err != nil {
		return BulkUpdateResp{}, err
	}

	return BulkUpdateResp{}, nil
}

func validateBulkUpdateRequest(args BulkUpdateReq, responses pipe.Responses) (response any, err error) {
	for i, productUpdate := range args.ProductSearchUpdate {
		var notEmpty bool
		if productUpdate.Title.String != "" {
			notEmpty = true
		}
		if productUpdate.CTAURL.String != "" {
			notEmpty = true
		}
		if productUpdate.ImageURL.String != "" {
			notEmpty = true
		}
		if !notEmpty {
			return nil, fmt.Errorf("empty product[%d]", i)
		}
	}
	return nil, nil
}
