package main

import (
	"fmt"

	"github.com/pozelim/product-review-app/user/internal/config"
)

func main() {
	application := config.NewApplication()
	err := application.Start()
	fmt.Println(err)
}
