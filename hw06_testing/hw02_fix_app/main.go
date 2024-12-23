package main

import (
	"fmt"

	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw02_fix_app/printer"
	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw02_fix_app/reader"
	"github.com/napryag/hw_otus_basic/hw06_testing/hw06_testing/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
