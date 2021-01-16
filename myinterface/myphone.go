package myinterface

import "fmt"

type MyPhone struct {
}

func (mp MyPhone) Bluetooth() {
	fmt.Println("Mobile phone connected by Bluetooth (wireless)")
}

func (mp MyPhone) MicroUSBA() {
	fmt.Println("Mobile phone connected by Micro-USB Type-A")
}

func (mt MyPhone) USBC() {
	fmt.Println("Mobile phone connected by USB Tpye-C")
}
