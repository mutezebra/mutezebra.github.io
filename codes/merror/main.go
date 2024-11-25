package main

import (
	"errors"
	"fmt"
	"mutezebra/pkg/merror"
)

func main() {
	e := merror.New("a error", 1024)
	fmt.Printf("%s\n", e.Error())

	var ex merror.Merror
	if errors.As(e, &ex) {
		fmt.Printf("%d\n", ex.Extra())
		fmt.Printf("%+v\n", ex)
	}
}
