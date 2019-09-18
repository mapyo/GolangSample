package main

import (
	"fmt"
	"go.uber.org/multierr"
)
import (
	"golang.org/x/xerrors"
)

func errXerrorsSample() {
	err := xerrors.New("first error")
	err = xerrors.Errorf("second error: %w", err)

	fmt.Println("------------")
	fmt.Println("errXerrorsSample")
	fmt.Println("------------")
	fmt.Println(err)
	fmt.Println("------------")
	fmt.Printf("%+v\n", err)
	fmt.Println("------------")
}

func errMultierrSample() {
	err := xerrors.New("first error")
	err = multierr.Append(xerrors.New("second error"), err)

	fmt.Println("------------")
	fmt.Println("errMultierrSample")
	fmt.Println("------------")
	fmt.Println(err)
	fmt.Println("------------")
	fmt.Printf("%+v\n", err)
	fmt.Println("------------")
}

func errMultierrNilSample() {
	var err error
	err = multierr.Append(xerrors.New("second error"), err)

	fmt.Println("------------")
	fmt.Println("errMultierrNilSample")
	fmt.Println("------------")
	fmt.Println(err)
	fmt.Println("------------")
	fmt.Printf("%+v\n", err)
	fmt.Println("------------")
}

func errXerrorsNilSample() {
	var err error
	err = xerrors.Errorf("second error: %w", nil)

	fmt.Println("------------")
	fmt.Println("errXerrorsNilSample")
	fmt.Println("------------")
	fmt.Println(err)
	fmt.Println("------------")
	fmt.Printf("%+v\n", err)
	fmt.Println("------------")
}

func main() {
	errXerrorsSample()
	errXerrorsNilSample()

	errMultierrSample()
	errMultierrNilSample()
}
