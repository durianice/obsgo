package main

import (
	"flag"
	"fmt"
	"os"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func main() {
	// 定义命令行参数
	operation := flag.String("op", "", "操作类型：put 或 del")
	bucket := flag.String("bucket", "", "存储桶名称")
	key := flag.String("key", "", "对象键")
	sourceFile := flag.String("file", "", "本地文件路径（仅用于上传）")

	flag.Parse()

	if *operation != "put" && *operation != "del" {
		fmt.Println("请指定有效的操作类型：put 或 del")
		flag.Usage()
		os.Exit(1)
	}

	if *bucket == "" || *key == "" {
		fmt.Println("请指定存储桶名称和对象键")
		flag.Usage()
		os.Exit(1)
	}

	if *operation == "put" && *sourceFile == "" {
		fmt.Println("上传操作需要指定本地文件路径")
		flag.Usage()
		os.Exit(1)
	}

	// 获取环境变量
	ak := os.Getenv("AccessKeyID")
	sk := os.Getenv("SecretAccessKey")
	// securityToken := os.Getenv("SecurityToken") // 如果需要使用临时凭证，可以启用
	endPoint := os.Getenv("endPoint")

	if ak == "" || sk == "" || endPoint == "" {
		fmt.Println("请设置 AccessKeyID、SecretAccessKey 和 endPoint 环境变量")
		flag.Usage()
		os.Exit(1)
	}

	// 创建 OBS 客户端
	// 如果使用临时凭证，取消注释以下代码
	/*
	   obsClient, err := obs.New(ak, sk, endPoint, obs.WithSecurityToken(securityToken))
	   if err != nil {
	       fmt.Printf("创建 OBS 客户端失败，错误信息: %s\n", err.Error())
	       os.Exit(1)
	   }
	*/

	// 不使用临时凭证
	obsClient, err := obs.New(ak, sk, endPoint)
	if err != nil {
		fmt.Printf("创建 OBS 客户端失败，错误信息: %s\n", err.Error())
		os.Exit(1)
	}

	// 根据操作类型调用相应的函数
	if *operation == "put" {
		err = PutFile(obsClient, *bucket, *key, *sourceFile)
	} else if *operation == "del" {
		err = DeleteObject(obsClient, *bucket, *key)
	}

	if err != nil {
		os.Exit(1)
	}
}
