package main

import (
	"fmt"
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
)

func main() {
	var inputYear int = 2026
	var dateOfEaster patype.FullDate = palib.GetDateOfEaster(inputYear)

	fmt.Printf("Date of Easter for %v is %v/%v/%v\n", inputYear, dateOfEaster.Month, dateOfEaster.Day, dateOfEaster.Year)
	fmt.Println("The library is accessible and working.")
}
