package model

import "fmt"

type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func Constructor(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this Customer) Info() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", this.ID, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}

func (this *Customer) UpdateCustomer(name string, gender string,
	age int, phone string, email string) {
	if name != "" {
		this.Name = name
	}
	if gender != "" {
		this.Gender = gender
	}
	if age != 0 {
		this.Age = age
	}
	if phone != "" {
		this.Phone = phone
	}
	if email != "" {
		this.Email = email
	}
}
