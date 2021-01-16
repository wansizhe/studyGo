package main

import (
	"fmt"
	"mycustomer/customerManage/model"
	"mycustomer/customerManage/service"
)

type customerView struct {
	key  string
	loop bool
	cS   *service.CustomerService
}

func (this *customerView) mainMenu() {
	for this.loop {
		fmt.Println("\n客户信息管理")
		fmt.Println("1、添加客户")
		fmt.Println("2、修改客户")
		fmt.Println("3、删除客户")
		fmt.Println("4、客户列表")
		fmt.Println("5、退出")
		fmt.Print("请选择（1-5）：")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.showList()
		case "5":
			this.exit()
		default:
			fmt.Println("输入错误")
		}
	}
	fmt.Println("已退出")
}

func (this *customerView) showList() {
	customers := this.cS.List()
	fmt.Println("\n客户列表")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, c := range customers {
		fmt.Println(c.Info())
	}
	fmt.Println("已显示全部客户")
}

func (this *customerView) add() {
	fmt.Println("\n添加新客户")
	name := ""
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	gender := ""
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	age := 0
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	phone := ""
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	email := ""
	fmt.Print("邮箱：")
	fmt.Scanln(&email)
	c := model.NewCustomer(name, gender, age, phone, email)
	if this.cS.Add(c) {
		fmt.Println("添加完成")
	} else {
		fmt.Println("添加失败")
	}
}

func (this *customerView) delete() {
	fmt.Println("\n删除某客户")
	id := -1
	confirm := ""
	fmt.Print("ID：")
	fmt.Scanln(&id)
	for confirm != "y" && confirm != "n" {
		fmt.Print("确认删除（y/n）：")
		fmt.Scanln(&confirm)
	}
	if confirm == "y" {
		if this.cS.Delete(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("查无此人")
		}
	} else {
		fmt.Println("放弃删除")
	}

}

func (this *customerView) update() {
	fmt.Println("\n修改某客户")
	id := -1
	fmt.Print("ID：")
	fmt.Scanln(&id)
	idx := this.cS.FindByID(id)
	if idx != -1 {
		name := ""
		fmt.Print("姓名：")
		fmt.Scanln(&name)
		gender := ""
		fmt.Print("性别：")
		fmt.Scanln(&gender)
		age := 0
		fmt.Print("年龄：")
		fmt.Scanln(&age)
		phone := ""
		fmt.Print("电话：")
		fmt.Scanln(&phone)
		email := ""
		fmt.Print("邮箱：")
		fmt.Scanln(&email)
		this.cS.Update(idx, name, gender, age, phone, email)
		fmt.Println("修改完成")
	} else {
		fmt.Println("查无此人")
	}
}

func (this *customerView) exit() {
	confirm := ""
	for confirm != "y" && confirm != "n" {
		fmt.Print("确认退出（y/n）：")
		fmt.Scanln(&confirm)
	}
	if confirm == "y" {
		this.loop = false
	}
}

func main() {
	cV := customerView{
		key:  "",
		loop: true,
	}
	cV.cS = service.Constructor()
	cV.mainMenu()
}
