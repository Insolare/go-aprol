package main

import (
	"fmt"

	"github.com/Insolare/goaprol/tbase"
)

type Enumerator struct {
	Vsets []tbase.Vset
}

func (e *Enumerator) EnumeratorCallback(a tbase.Vset, b interface{}) {
	e.Vsets = append(e.Vsets, a)
}

func main() {
	base := tbase.Tbase{}
	en := Enumerator{}

	err := base.Connect("10.7.10.182")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer base.Disconnect()

	err = base.OpenDatabase("/home/engin/ENGIN/caedb")
	if err != nil {
		fmt.Println(err)
		return
	}

	base.Enumerate("GLOBAL", &en)

	defer func() {
		for i := range en.Vsets {
			en.Vsets[i].Free()
		}
	}()

	for _, v := range en.Vsets {
		fmt.Println(v.Self)

		ch := v.ListFields()
		for fs := range ch {
			fmt.Println(fs)
		}

		base.Enumerate(v.Self, &v)
		for i := range v.Children {
			fmt.Println(v.Children[i].Self)
		}
	}
}
