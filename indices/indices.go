package indices

import (
	"log"

	"github.com/olivere/elastic/v7"
)

func Create(ec *elastic.Client) {
	log.Println("creating ES indices")
	CreateProductSearch(ec)
	CreateProductSearchShortener(ec)
}
