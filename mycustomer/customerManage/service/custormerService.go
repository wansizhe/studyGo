package service

import "mycustomer/customerManage/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func Constructor() *CustomerService {
	cS := &CustomerService{}
	cS.customerNum = 1
	customer := model.Constructor(1, "张三", "男", 20, "123456789", "zs@example.com")
	cS.customers = append(cS.customers, customer)
	return cS
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(c model.Customer) bool {
	this.customerNum++
	c.ID = this.customerNum
	this.customers = append(this.customers, c)
	return true
}

func (this *CustomerService) FindByID(id int) int {
	idx := -1
	for i, c := range this.customers {
		if c.ID == id {
			idx = i
			break
		}
	}
	return idx
}

func (this *CustomerService) Delete(id int) bool {
	idx := this.FindByID(id)
	if idx == -1 {
		return false
	}
	this.customers = append(this.customers[:idx], this.customers[idx+1:]...)
	return true
}

func (this *CustomerService) Update(idx int, name string, gender string, age int, phone string, email string) {
	this.customers[idx].UpdateCustomer(name, gender, age, phone, email)
}
