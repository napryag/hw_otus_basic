package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/napryag/hw_otus_basic/hw09_serialize/types"
)

func main() {
	p1 := &types.Book{
		ID:     1,
		Title:  "Отцы и дети",
		Author: "Грибоедов",
		Year:   1924,
		Size:   132,
		Rate:   4.6,
	}
	var p2, p3, p4, p5, p6 types.Book

	b, err := types.MarshalJSON(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = types.UnmarshalJSON(b, &p2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("JSON: %v\n", p2)

	c, err := types.MarshalXML(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = types.UnmarshalXML(c, &p3)
	if err != nil && errors.Is(err, io.EOF) {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("XML: %v\n", p3)

	d, err := types.MarshalYAML(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = types.UnmarshalYAML(d, &p4)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("YAML: %v\n", p4)

	buff, err := types.EncodeGOB(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = types.DecodeGOB(buff, &p5)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("GOB: %v\n", p5)

	e, err := types.MarshalBSON(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = types.UnmarshalBSON(e, &p6)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("BSON: %v\n", p6)
}
