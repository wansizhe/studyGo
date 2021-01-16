package myinterface

import "fmt"

type MyMonitor struct {
}

func (mt MyMonitor) HDMI() {
	fmt.Println("Monitor connected by HDMI")
}

func (mt MyMonitor) VGA() {
	fmt.Println("Monitor connected by VGA")
}

func (mt MyMonitor) DP() {
	fmt.Println("Monitor connected by DisplayPort")
}

func (mt MyMonitor) USBC() {
	fmt.Println("Monitor connected by USB Tpye-C")
}
