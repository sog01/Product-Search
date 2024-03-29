package indices

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func CreateProductSearch(ec *elastic.Client) error {
	body := `
    {
        "settings": {
            "analysis": {
                "analyzer": {
                    "autocomplete": {
                        "tokenizer": "autocomplete",
                        "filter": [
                            "lowercase"
                        ]
                    }
                },
                "tokenizer": {
                    "autocomplete": {
                        "type": "edge_ngram",
                        "min_gram": 1,
                        "max_gram": 10,
                        "token_chars": [
                            "letter",
                            "digit"
                        ]
                    }
                }
            }
        },
        "mappings": {
            "properties": {
                "id": {
                    "type": "keyword"
                },
                "title": {
                    "type": "text",
                    "analyzer": "autocomplete",
                    "fields": {
                        "keyword": {
                            "type": "keyword"
                        }
                    }
                },
                "description": {
                    "type": "text"                    
                },
                "cta_url": {
                    "type": "keyword"
                },
                "image_url": {
                    "type": "keyword"
                },
                "price": {
                    "type": "double"
                },
                "catalog": {
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
	_, err := ec.CreateIndex("product_search").
		BodyString(body).
		Do(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
