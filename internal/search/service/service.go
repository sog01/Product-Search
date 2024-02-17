package service

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"github.com/sog01/productdiscovery/internal/search/service/mutation"
	"github.com/sog01/productdiscovery/internal/search/service/query"
)

type Service interface {
	Search(ctx context.Context, req model.SearchReq) (model.SearchResponse, error)
	SearchAutocomplete(ctx context.Context, req model.AutocompleteReq) (model.AutocompleteResp, error)
	SearchTotal(ctx context.Context, req model.SearchTotalReq) (model.SearchTotalResp, error)
	SearchCatalogs(ctx context.Context, req model.SearchCatalogsReq) (model.SearchCatalogsResp, error)
	SearchTopProductCatalogs(ctx context.Context, req model.SearchTopProductCatalogReq) (model.SearchTopProductCatalogResp, error)
	BulkInsert(ctx context.Context, req model.BulkInsertReq) (model.BulkInsertResp, error)
	BulkUpdate(ctx context.Context, req model.BulkUpdateReq) (model.BulkUpdateResp, error)
	UploadProductCsv(ctx context.Context, req model.UploadProductCsvReq) (model.UploadProductCsvResp, error)
}

type SearchService struct {
	searchRepo         repository.SearchRepository
	searchTotalRepo    repository.SearchTotalRepository
	searchCategoryRepo repository.SearchCategoryRepository
	bulkUpdateRepo     repository.BulkUpdateRepository
	bulkInsertRepo     repository.BulkInsertRepository
}

func (s SearchService) Search(ctx context.Context, req model.SearchReq) (model.SearchResponse, error) {
	return query.Search(ctx, req, s.searchRepo)
}
func (s SearchService) SearchAutocomplete(ctx context.Context, req model.AutocompleteReq) (model.AutocompleteResp, error) {
	return query.SearchAutocomplete(ctx, req, s.searchRepo)
}
func (s SearchService) SearchTotal(ctx context.Context, req model.SearchTotalReq) (model.SearchTotalResp, error) {
	return query.SearchTotal(ctx, req, s.searchTotalRepo)
}
func (s SearchService) SearchCatalogs(ctx context.Context, req model.SearchCatalogsReq) (model.SearchCatalogsResp, error) {
	return query.SearchCatalogs(ctx, req, s.searchCategoryRepo)
}
func (s SearchService) SearchTopProductCatalogs(ctx context.Context, req model.SearchTopProductCatalogReq) (model.SearchTopProductCatalogResp, error) {
	return query.SearchTopProductCatalog(ctx, req, s.searchRepo)
}
func (s SearchService) BulkInsert(ctx context.Context, req model.BulkInsertReq) (model.BulkInsertResp, error) {
	return mutation.BulkInsert(ctx, req, s.bulkInsertRepo)
}
func (s SearchService) BulkUpdate(ctx context.Context, req model.BulkUpdateReq) (model.BulkUpdateResp, error) {
	return mutation.BulkUpdate(ctx, req, s.bulkUpdateRepo)
}
func (s SearchService) UploadProductCsv(ctx context.Context, req model.UploadProductCsvReq) (model.UploadProductCsvResp, error) {
	return mutation.UploadProductCSV(ctx, req, s.bulkInsertRepo)
}

func NewService(es *elastic.Client) Service {
	return SearchService{
		searchRepo:         repository.NewSearchRepository(es),
		searchTotalRepo:    repository.NewSearchTotalRepository(es),
		searchCategoryRepo: repository.NewSearchCatalogsRepository(es),
		bulkInsertRepo:     repository.NewBulkInsertRepository(es),
		bulkUpdateRepo:     repository.NewBulkUpdateRepository(es),
	}
}
