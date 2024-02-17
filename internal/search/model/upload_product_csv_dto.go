package model

import "mime/multipart"

type UploadProductCsvReq struct {
	CsvFile *multipart.FileHeader
}

type UploadProductCsvResp struct {
}
