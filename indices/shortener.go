package indices

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func CreateProductSearchShortener(ec *elastic.Client) error {
	body := `{
        "mappings": {
            "properties": {
                "id": {
                    "type": "keyword"
                },
                "slug": {
					"type": "keyword"
				},
				"real_url": {
					"type": "keyword"
				},
                "created_at": {
                    "type": "date"
                },
                "updated_at": {
                    "type": "date"
                }
            }
        }
    }`
	_, err := ec.CreateIndex("product_search_shortener").
		BodyString(body).
		Do(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
