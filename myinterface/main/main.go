package main

import myinterface "MyInterface"

// MyBluetooth bluetooth interface
type MyBluetooth interface {
	Bluetooth()
}

// MyUSBA usb-A interface
type MyUSBA interface {
	USBA()
}

func main() {
	var bt MyBluetooth
	var usba MyUSBA
	bt = new(myinterface.MyEarphone)
	usba = new(myinterface.MyKeyboard)
	bt.Bluetooth()
	usba.USBA()
	bt = new(myinterface.MyKeyboard)
	bt.Bluetooth()
}
