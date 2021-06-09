package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	return WrapError(errors.New("asdaa"), "foo failed")
	//return errors.Wrap()
}

func bar() error {
	return errors.WithMessage(foo(), "bar 函数调用")
}

func main() {
	err := bar()

	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		fmt.Printf("sa === %+v\n", err)
	}
}
