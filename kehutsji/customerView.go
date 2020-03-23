package main

import (
	"fmt"
	"study/model"
	"study/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (this *customerView) exit() {
	fmt.Println("请确认是否退出Y/n：_")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("请确认是否退出Y/n")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}

}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("\n----------------------客户列表---------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i, _ := range customers {
		fmt.Println(customers[i].Info())
	}
	fmt.Printf("-------------------客户列表完成--------------------\n\n")
}

func (this *customerView) add() {
	fmt.Println("---1.添加用户---")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(age, name, gender, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-----添加完成-----")
	} else {
		fmt.Println("-----添加失败-----")
	}
}

func (this *customerView) delete() {
	fmt.Println("---3.删除用户---")
	fmt.Println("请输入用户ID, -1退出...")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认删除? Y/n")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if !this.customerService.Delete(id) {
			fmt.Println("-----删除失败，输入的用户id不存在-----")
		} else {
			fmt.Println("-----删除成功-----")
		}
	}
}

func (this *customerView) update() {
	fmt.Println("请输入要修改的用户ID")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := this.customerService.FindById(id)
	if index == -1 {
		fmt.Println("找不到用户ID")
		return
	}
	customers := this.customerService.List()
	fmt.Printf("姓名(%v)： ", customers[index].Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("年龄(%v)： ", customers[index].Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("性别(%v)： ", customers[index].Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("Phone(%v)： ", customers[index].Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("电邮(%v)： ", customers[index].Email)
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(id, age, name, gender, phone, email)
	if this.customerService.Update(customer) {
		fmt.Println("-----修改完成-----")
	} else {
		fmt.Println("-----修改失败,用户id不存在-----")
	}
}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("---管理软件---")
		fmt.Println("---1.增添用户---")
		fmt.Println("---2.修改用户---")
		fmt.Println("---3.删除用户---")
		fmt.Println("---4.用户列表---")
		fmt.Println("---5.退出程序---")
		fmt.Println("请输入1-5：_")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你输入有误...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你已退出程序...")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
