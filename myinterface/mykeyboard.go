package myinterface

import "fmt"

type MyKeyboard struct {
	id string
}

func (kb MyKeyboard) USBA() {
	fmt.Println("Keyboard connected by USB Type-A")
}

func (kb MyKeyboard) Bluetooth() {
	fmt.Println("Keyboard connected by Bluetooth (wireless)")
}
