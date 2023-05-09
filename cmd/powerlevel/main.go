package main

import (
	"fmt"
	"github.com/BenjaminGlusa/powerlevel/pkg/adapter"
)

func main() {
	fmt.Println("Power level")

	var db adapter.DatabaseAdapter = adapter.NewMySqlAdapter("power", "power", "tv")
	defer db.Close()
	db.CreateTableIfNotExits()

	kwhToday := db.KwhToday()
	kwhThisMonth := db.KwhThisMonth()
	KwhThisYear := db.KwhThisYear()
	KwhTotal := db.KwhTotal()
	fmt.Printf("kwh today: %.3f \n", kwhToday)
	fmt.Printf("kwh this month: %.3f \n", kwhThisMonth)
	fmt.Printf("kwh this year: %.3f \n", KwhThisYear)
	fmt.Printf("kwh total: %.3f \n", KwhTotal)

	fmt.Println("all done")

}
