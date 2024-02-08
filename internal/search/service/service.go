package service

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
)

type Service interface {
	Search(ctx context.Context, req model.SearchReq) (model.SearchResponse, error)
	SearchSuggestion(ctx context.Context, req model.SuggestionReq) (model.SuggestionResp, error)
	BulkInsert(ctx context.Context, req model.BulkInsertReq) (model.BulkInsertResp, error)
	BulkUpdate(ctx context.Context, req model.BulkUpdateReq) (model.BulkUpdateResp, error)
}

type SearchService struct {
	searchRepo     repository.SearchRepository
	bulkUpdateRepo repository.BulkUpdateRepository
	bulkInsertRepo repository.BulkInsertRepository
}

func (s SearchService) Search(ctx context.Context, req model.SearchReq) (model.SearchResponse, error) {
	return Search(ctx, req, s.searchRepo)
}
func (s SearchService) SearchSuggestion(ctx context.Context, req model.SuggestionReq) (model.SuggestionResp, error) {
	return SearchSuggestion(ctx, req, s.searchRepo)
}
func (s SearchService) BulkInsert(ctx context.Context, req model.BulkInsertReq) (model.BulkInsertResp, error) {
	return BulkInsert(ctx, req, s.bulkInsertRepo)
}
func (s SearchService) BulkUpdate(ctx context.Context, req model.BulkUpdateReq) (model.BulkUpdateResp, error) {
	return BulkUpdate(ctx, req, s.bulkUpdateRepo)
}

func NewService(es *elastic.Client) Service {
	return SearchService{
		searchRepo:     repository.NewSearchRepository(es),
		bulkInsertRepo: repository.NewBulkInsertRepository(es),
		bulkUpdateRepo: repository.NewBulkUpdateRepository(es),
	}
}
