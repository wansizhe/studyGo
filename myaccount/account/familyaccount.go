package account

import "fmt"

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	detail  string
	flag    bool
}

func Constructor() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		detail:  "收支\t账户金额\t收支金额\t说明",
		flag:    false,
	}
}

func (this *FamilyAccount) MainMenu() {
	for this.loop {
		fmt.Println("\n家庭收支记账软件")
		fmt.Println("1、收支明细")
		fmt.Println("2、登记收入")
		fmt.Println("3、登记支出")
		fmt.Println("4、退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetail()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("输入错误")
		}
	}
}

func (this *FamilyAccount) showDetail() {
	if this.flag {
		fmt.Println("收支明细")
		fmt.Println("当前余额：", this.balance)
		fmt.Println(this.detail)
	} else {
		fmt.Println("当前余额：", this.balance)
		fmt.Println("无收支情况")
	}
}

func (this *FamilyAccount) income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	for this.money > this.balance {
		fmt.Println("余额不足")
		fmt.Print("本次支出金额：")
		fmt.Scanln(&this.money)
	}
	this.balance -= this.money
	fmt.Print("本次支出说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Print("确认退出？[y/n]：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("输入错误，重新输入：")
	}
	if choice == "y" {
		this.loop = false
	}
}
