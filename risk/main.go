// A service that creates a snapshot repository the way go-elasticsearch
// 8.12 does it. 8.13 removed createrepository.NewRequest, so a minor bump
// that every other gate lets through would break this build.
package main

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/createrepository"
)

func main() {
	req := createrepository.NewRequest()
	fmt.Println(req != nil)
}
