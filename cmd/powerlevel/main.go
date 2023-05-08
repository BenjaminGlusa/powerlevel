package main

import(
	"fmt"
	"github.com/BenjaminGlusa/powerlevel/pkg/adapter"
)

func main() {
	fmt.Println("Power level")
	
	db := adapter.NewMySqlAdapter("power", "power", "tv")
	defer db.Close()
	db.CreateTableIfNotExits()
	
	fmt.Println("all done")

}
