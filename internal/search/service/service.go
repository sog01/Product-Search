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
	SearchTotal(ctx context.Context, req model.SearchTotalReq) (model.SearchTotalResp, error)
	SearchCategory(ctx context.Context, req model.SearchCategoryReq) (model.SearchCategoryResp, error)
	BulkInsert(ctx context.Context, req model.BulkInsertReq) (model.BulkInsertResp, error)
	BulkUpdate(ctx context.Context, req model.BulkUpdateReq) (model.BulkUpdateResp, error)
}

type SearchService struct {
	searchRepo         repository.SearchRepository
	searchTotalRepo    repository.SearchTotalRepository
	searchCategoryRepo repository.SearchCategoryRepository
	bulkUpdateRepo     repository.BulkUpdateRepository
	bulkInsertRepo     repository.BulkInsertRepository
}

func (s SearchService) Search(ctx context.Context, req model.SearchReq) (model.SearchResponse, error) {
	return Search(ctx, req, s.searchRepo)
}
func (s SearchService) SearchSuggestion(ctx context.Context, req model.SuggestionReq) (model.SuggestionResp, error) {
	return SearchSuggestion(ctx, req, s.searchRepo)
}
func (s SearchService) SearchTotal(ctx context.Context, req model.SearchTotalReq) (model.SearchTotalResp, error) {
	return SearchTotal(ctx, req, s.searchTotalRepo)
}
func (s SearchService) SearchCategory(ctx context.Context, req model.SearchCategoryReq) (model.SearchCategoryResp, error) {
	return SearchCategory(ctx, req, s.searchCategoryRepo)
}
func (s SearchService) BulkInsert(ctx context.Context, req model.BulkInsertReq) (model.BulkInsertResp, error) {
	return BulkInsert(ctx, req, s.bulkInsertRepo)
}
func (s SearchService) BulkUpdate(ctx context.Context, req model.BulkUpdateReq) (model.BulkUpdateResp, error) {
	return BulkUpdate(ctx, req, s.bulkUpdateRepo)
}

func NewService(es *elastic.Client) Service {
	return SearchService{
		searchRepo:         repository.NewSearchRepository(es),
		searchTotalRepo:    repository.NewSearchTotalRepository(es),
		searchCategoryRepo: repository.NewSearchCategoryRepository(es),
		bulkInsertRepo:     repository.NewBulkInsertRepository(es),
		bulkUpdateRepo:     repository.NewBulkUpdateRepository(es),
	}
}
