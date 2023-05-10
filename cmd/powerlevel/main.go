package main

import (
	"fmt"
	"net/http"
	"github.com/BenjaminGlusa/powerlevel/pkg/adapter"
	"github.com/BenjaminGlusa/powerlevel/pkg/handler"
)


func main() {
	fmt.Println("Power level")
	port := ":4081"

	var db adapter.DatabaseAdapter = adapter.NewMySqlAdapter("power", "power", "tv")
	defer db.Close()
	db.CreateTableIfNotExits()

	fmt.Printf("Server started on port %s...\n", port)

	http.ListenAndServe(port, handler.NewRouter(db))

}
