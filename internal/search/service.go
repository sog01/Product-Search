package search

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type Service interface {
	Search(ctx context.Context, req SearchReq) (SearchResponse, error)
	SearchSuggestion(ctx context.Context, req SuggestionReq) (SuggestionResp, error)
	BulkInsert(ctx context.Context, req BulkInsertReq) (BulkInsertResp, error)
	BulkUpdate(ctx context.Context, req BulkUpdateReq) (BulkUpdateResp, error)
}

type SearchService struct {
	searchRepo     SearchRepository
	bulkUpdateRepo BulkUpdateRepository
	bulkInsertRepo BulkInsertRepository
}

func (s SearchService) Search(ctx context.Context, req SearchReq) (SearchResponse, error) {
	return Search(ctx, req, s.searchRepo)
}
func (s SearchService) SearchSuggestion(ctx context.Context, req SuggestionReq) (SuggestionResp, error) {
	return SearchSuggestion(ctx, req, s.searchRepo)
}
func (s SearchService) BulkInsert(ctx context.Context, req BulkInsertReq) (BulkInsertResp, error) {
	return BulkInsert(ctx, req, s.bulkInsertRepo)
}
func (s SearchService) BulkUpdate(ctx context.Context, req BulkUpdateReq) (BulkUpdateResp, error) {
	return BulkUpdate(ctx, req, s.bulkUpdateRepo)
}

func NewService(es *elastic.Client) Service {
	return SearchService{
		searchRepo:     NewSearchRepository(es),
		bulkInsertRepo: NewBulkInsertRepository(es),
		bulkUpdateRepo: NewBulkUpdateRepository(es),
	}
}
