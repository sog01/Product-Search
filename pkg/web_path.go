package pkg

import (
	"os"
	"sync"
)

var (
	once = sync.Once{}
	webP string
)

func WebPath() string {
	once.Do(func() {
		webP = "./../../web"
		_, err := os.Stat(webP)
		if os.IsNotExist(err) {
			webP = "./web"
		}
	})

	return webP
}
