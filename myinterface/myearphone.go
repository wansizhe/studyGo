package myinterface

import "fmt"

type MyEarphone struct {
}

func (ep MyEarphone) Hole35() {
	fmt.Println("Earphone connected by 3.5mm hole")
}

func (ep MyEarphone) Bluetooth() {
	fmt.Println("Earphone connected by Bluetooth (wireless)")
}
