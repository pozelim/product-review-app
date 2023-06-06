package main

import (
	"fmt"

	"github.com/pozelim/product-review-app/user/config"
)

func main() {
	application := config.NewApplication()
	err := application.Start()
	fmt.Println(err)
}
