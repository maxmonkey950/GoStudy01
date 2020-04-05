package main

import (
	"fmt"
	"github.com/docker/docker/client"
)

type Menu struct {
	key  string // 捕获用户输入
	loop bool   // 标记主菜单是否循环
	cli  *client.Client
}

// 主菜单函数
func (this *Menu) MainMenu() {
	for {
		fmt.Println("1. 创建环境")
		fmt.Println("2. 删除环境")
		fmt.Println("3. 显示所有环境")
		fmt.Println("4. 本地镜像列表")
		fmt.Println("5. 构建镜像")
		fmt.Println("q. 退出")
		fmt.Println("---请选择(1-5 or Q)---")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("1. 创建环境")
			this.TchEn()
		case "2":
			fmt.Println("2. 删除环境")
			this.Del()
		case "3":
			fmt.Println("3. 显示所有环境")
			this.showDetails()
		case "4":
			fmt.Println("4. 本地镜像列表")
			this.ImagesList()
		case "5":
			fmt.Println("5. 构建镜像")
			this.ImageBuild()
		case "q":
			fmt.Println("q. 退出")
			this.Exit()
		default:
			fmt.Println("输入有误...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你已退出程序...")
}

func main() {
	menu := Menu{
		key:  "",
		loop: true,
		cli:  Conn(),
	}
	menu.MainMenu()
}

