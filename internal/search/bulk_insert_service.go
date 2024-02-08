package search

import (
	"context"

	"github.com/sog01/pipe"
)

type BulkInsertRepository struct {
	BulkInsert pipe.Func[BulkInsertReq]
}

func BulkInsert(ctx context.Context, req BulkInsertReq, repo BulkInsertRepository) (BulkInsertResp, error) {
	exec := pipe.P(
		repo.BulkInsert,
	)

	req.ctx = ctx
	_, err := exec(req)
	if err != nil {
		return BulkInsertResp{}, err
	}

	return BulkInsertResp{}, nil
}
