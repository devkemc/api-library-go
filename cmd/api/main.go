package main

import (
	"fmt"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/router"
)

func main() {
	router.Initialize()
	fmt.Print("Server is running on port 8081")
}
