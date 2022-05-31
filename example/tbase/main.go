package main

import (
	"fmt"

	"github.com/Insolare/goaprol/tbase"
)

func main() {

	base := tbase.Tbase{}

	err := base.Connect("10.7.10.182")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = base.OpenDatabase("/home/runtime/RUNTIME/runtimedb")
	if err != nil {
		fmt.Println(err)
		return
	}

	vs := base.Enumerate("")

	defer func() {
		for i := range vs {
			vs[i].Free()
		}
	}()

	for _, v := range vs {
		fmt.Println(v.Self)
		if v.Self == "LSC" {
			if s, err := v.GetString("self"); err == nil {
				fmt.Println("==>", s)
			}
		}
	}

	alarmGroups := base.Enumerate("G\\A")
	defer func() {
		for i := range alarmGroups {
			alarmGroups[i].Free()
		}
	}()

	for _, v := range alarmGroups {
		fmt.Println(v.Self)
	}

	alarmRefs := make([][]tbase.Vset, len(alarmGroups))

	for i := range alarmRefs {
		alarmRefs[i] = base.GetReferences(alarmGroups[i].Self)
	}

	base.Disconnect()
}
