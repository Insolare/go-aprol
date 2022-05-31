package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/insolare/goaprol/iosys"
)

type MyTimeVar struct {
	v    iosys.IosVar
	Self string
}

func (m *MyTimeVar) OnChange() {
	v, err := m.v.GetTimestamp()
	if err != nil {
		fmt.Println(m)
	}
	fmt.Println(m.Self, "changed to", v)
}

func (m *MyTimeVar) OnChangeRequest() {

}

func (m *MyTimeVar) OnIdleChange() {

}

type MyIntVar struct {
	v    iosys.IosVar
	Self string
}

func (m *MyIntVar) OnChange() {
	v, _ := m.v.GetInt()
	fmt.Println(m.Self, "changed to", v)
}

func (m *MyIntVar) OnChangeRequest() {

}

func (m *MyIntVar) OnIdleChange() {

}

func main() {
	wg := sync.WaitGroup{}
	readTimeVar := MyTimeVar{}
	readIntVar := MyIntVar{}
	stopCh := make(chan struct{})

	iosys.Initialize()
	ios := iosys.NewIosysConnection("10.7.10.182:0")
	ios.Connect(5)

	readTimeVar.Self = "p_timedisp_AprFbGetDT.OUT"
	readIntVar.Self = "p_demotrd_AprFbPulse_v.Out"

	readIntVar.v = iosys.NewIosVar(readIntVar.Self, &readIntVar)
	readTimeVar.v = iosys.NewIosVar(readTimeVar.Self, &readTimeVar)

	cancel := iosys.StartMainloop(&wg)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
		}

		stopCh <- struct{}{}
		wg.Done()
	}()

	<-stopCh
	cancel()

	wg.Wait()
	ios.Delete()
	iosys.Finalize()
}
