package service

import (
	"study/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(customerService.customerNum, 20, "tom", "male", "110", "tom@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i, _ := range this.customers {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (this *CustomerService) Update(customer model.Customer) bool {
	index := this.FindById(customer.Id)
	//index == -1 ，没有这个客户
	if index == -1 {
		return false
	}
	if customer.Name != "" {
		this.customers[index].Name = customer.Name
	}
	if customer.Gender != "" {
		this.customers[index].Gender = customer.Gender
	}
	if customer.Age != 0 {
		this.customers[index].Age = customer.Age
	}
	if customer.Phone != "" {
		this.customers[index].Phone = customer.Phone
	}
	if customer.Email != "" {
		this.customers[index].Email = customer.Email
	}
	return true
}
