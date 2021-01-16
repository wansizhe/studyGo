package main

import (
	"fmt"
	"myaccount/account"
)

func main() {
	account.Constructor().MainMenu()
	// key := ""
	// loop := true

	// balance := 10000.0
	// money := 0.0
	// note := ""
	// detail := "收支\t账户金额\t收支金额\t说明"
	// flag := false

	// for loop {
	// 	fmt.Println("\n家庭收支记账软件")
	// 	fmt.Println("1、收支明细")
	// 	fmt.Println("2、登记收入")
	// 	fmt.Println("3、登记支出")
	// 	fmt.Println("4、退出软件")
	// 	fmt.Print("请选择（1-4）：")

	// 	fmt.Scanln(&key)
	// 	switch key {
	// 	case "1":
	// 		if flag {
	// 			fmt.Println("收支明细")
	// 			fmt.Println(detail)
	// 		} else {
	// 			fmt.Println("无收支情况")
	// 		}
	// 	case "2":
	// 		fmt.Print("本次收入金额：")
	// 		fmt.Scanln(&money)
	// 		balance += money
	// 		fmt.Print("本次收入说明：")
	// 		fmt.Scanln(&note)
	// 		detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
	// 		flag = true
	// 	case "3":
	// 		fmt.Print("本次支出金额：")
	// 		fmt.Scanln(&money)
	// 		for money > balance {
	// 			fmt.Println("余额不足")
	// 			fmt.Print("本次支出金额：")
	// 			fmt.Scanln(&money)
	// 		}
	// 		balance -= money
	// 		fmt.Print("本次支出说明：")
	// 		fmt.Scanln(&note)
	// 		detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
	// 		flag = true
	// 	case "4":
	// 		fmt.Print("确认退出？[y/n]：")
	// 		choice := ""
	// 		for {
	// 			fmt.Scanln(&choice)
	// 			if choice == "y" || choice == "n" {
	// 				break
	// 			}
	// 			fmt.Print("输入错误，重新输入：")
	// 		}
	// 		if choice == "y" {
	// 			loop = false
	// 		}
	// 	default:
	// 		fmt.Println("输入错误")
	// 	}
	// }
	fmt.Println("\n已退出")
}
