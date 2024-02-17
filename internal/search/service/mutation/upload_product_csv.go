package mutation

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/sog01/pipe"
	"github.com/sog01/productdiscovery/internal/search/model"
	"github.com/sog01/productdiscovery/internal/search/repository"
	"gopkg.in/guregu/null.v4"
)

func UploadProductCSV(ctx context.Context, req model.UploadProductCsvReq, repo repository.BulkInsertRepository) (model.UploadProductCsvResp, error) {
	exec := pipe.PCtx(
		readCSV,
		bulkInsert(repo),
	)
	_, err := exec(ctx, req)
	if err != nil {
		log.Printf("failed to upload product csv: \n", err)
	}
	return model.UploadProductCsvResp{}, nil
}

func readCSV(ctx context.Context, req model.UploadProductCsvReq, responses pipe.Responses) (response any, err error) {
	file, err := req.CsvFile.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	resp := []model.ProductSearchInsert{}

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv rows: %v", err)
	}
	headerIndex := make(map[string]int)
	for i, h := range header {
		headerIndex[strings.ReplaceAll(strings.ToLower(h), " ", "_")] = i
	}

	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read csv rows: %v", err)
		}

		price, _ := strconv.ParseInt(records[headerIndex["price"]], 10, 64)
		resp = append(resp, model.ProductSearchInsert{
			Title:       records[headerIndex["title"]],
			Description: null.StringFrom(records[headerIndex["description"]]),
			CTAURL:      records[headerIndex["cta_url"]],
			ImageURL:    records[headerIndex["image_url"]],
			Price:       float64(price),
			Catalog:     null.StringFrom(records[headerIndex["catalog"]]),
		})
	}

	return resp, nil
}

func bulkInsert(repo repository.BulkInsertRepository) pipe.FuncCtx[model.UploadProductCsvReq] {
	return func(ctx context.Context, args model.UploadProductCsvReq, responses pipe.Responses) (response any, err error) {
		exec := pipe.PCtx(repo.BulkInsert)
		_, err = exec(ctx, model.BulkInsertReq{
			ProductSearchInput: pipe.Get[[]model.ProductSearchInsert](responses),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to bulk insert product csv: %v", err)
		}
		return nil, nil
	}
}
