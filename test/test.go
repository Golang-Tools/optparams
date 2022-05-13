package main

import (
	"fmt"

	"github.com/Golang-Tools/optparams"
)

type NameOpt struct {
	Name string
}

func WithName(name string) optparams.Option[NameOpt] {
	return optparams.NewFuncOption(
		func(o *NameOpt) {
			o.Name = name
		})
}

func expfunc(opts ...optparams.Option[NameOpt]) {
	opt := optparams.GetOption(&NameOpt{Name: "a"}, opts...)
	fmt.Println("opt:", opt)
}

func main() {
	expfunc(WithName("b"))
}
