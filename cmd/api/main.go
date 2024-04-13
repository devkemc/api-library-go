package main

import (
	"fmt"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/routes"
)

func main() {
	routes.RouterInit()
	fmt.Print("Server is running on port 8081")
}
