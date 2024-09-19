package main

import (
	"fmt"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func DeleteObject(obsClient *obs.ObsClient, bucket, key string) error {
	input := &obs.DeleteObjectInput{
		Bucket: bucket,
		Key:    key,
	}

	output, err := obsClient.DeleteObject(input)
	if err == nil {
		fmt.Printf("从存储桶(%s)删除对象(%s)成功！\n", input.Key, input.Bucket)
		fmt.Printf("RequestId: %s\n", output.RequestId)
		return nil
	}

	fmt.Printf("从存储桶(%s)删除对象(%s)失败！\n", input.Key, input.Bucket)
	if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println("发生 OBS 错误，您的请求被 OBS 拒绝。")
		fmt.Println(obsError.Error())
	} else {
		fmt.Println("发生异常，客户端在与 OBS 通信时遇到内部问题。")
		fmt.Println(err)
	}
	return err
}
