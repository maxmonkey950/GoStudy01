package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"os/exec"
)

// 返回docker api 链接指针
func Conn() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return cli
}

// 判断重复环境
func FindByName(name string) (bool bool, id string) {
	setname := "/" + name
	containers, err := Conn().ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return false, ""
	}
	for _, v := range containers {
		if setname == v.Names[0] {
			return true, v.ID[:4]
		}
	}
	return false, ""
}

// 删除容器
func removeContainer(containerID string, cli *client.Client) (string, error) {
	err := cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{
		Force: true,
	})
	return containerID, err
}

// 删除容器
func (this *Menu) Del() {
	choice := ""
	fmt.Printf("请输入要删除的容器名称:(谨慎操作,不可逆!)")
	fmt.Scanln(&choice)
	res, id := FindByName(choice)
	if !res {
		fmt.Println("找不到名称= %s的容器,请仔细确认...\n", choice)
		return
	}
	fmt.Printf("已经找到名称 = %s的容器...ID = %s\n", choice, id)
	fmt.Println("del.....")
	cid, err := removeContainer(id, this.cli)
	if err == nil {
		fmt.Println("删除容器", cid, "成功")
	} else {
		fmt.Println("删除失败 id = ", cid)
		fmt.Println(err)
	}
}

// 退出程序
func (this *Menu) Exit() {
	fmt.Printf("请确认是否退出Y/n: ")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Printf("请确认是否退出Y/n: ")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}

}

// 显示所有环境函数
func (this *Menu) showDetails() {
	fmt.Println("---所有环境---")

	containers, err := this.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("容器ID:[%s]\t 容器Name:%s\t 容器镜像:[%s]\t 容器状态:[%s]\n", container.ID[:4], container.Names, container.Image, container.Status)
	}
	fmt.Println("---环境如上---")
}

// 创建环境
func (this *Menu) TchEn() {
	var (
		imgname string
		enname  string
	)
	fmt.Println("请输入镜像名(exp: dev:v1): ")
	fmt.Scanln(&imgname)
	fmt.Println("请输入容器名称(exp: qa-test): ")
	fmt.Scanln(&enname)
	ctx := context.Background()
	resp, err := this.cli.ContainerCreate(ctx, &container.Config{
		Image: imgname,
	}, &container.HostConfig{}, nil, enname)
	if err != nil {
		fmt.Printf("镜像[%v]不存在...\n", imgname)
		return
	}

	if err := this.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Println("New env has created...")
	fmt.Printf("容器ID:[%s]\n", resp.ID[:4])
}

// 本地镜像列表
func listImage(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, image := range images {
		fmt.Printf("%v\t %v\t\n", image.ID, image.RepoTags)
	}
}

func (this *Menu) ImagesList() {
	listImage(this.cli)
}

func (this *Menu) ImageBuild() bool {
	return BuildImage()
}

func BuildImage() bool {
	command := `./scripts/auto.sh`
	cmd := exec.Command("/bin/bash", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s\n", command, err.Error())
		return false
	}
	fmt.Printf("Execute Shell:%s finished with output:%s\n", command, string(output))
	return false
}

