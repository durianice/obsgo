package main

import (
	"fmt"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func PutFile(obsClient *obs.ObsClient, bucket, key, sourceFile string) error {
	input := &obs.PutFileInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucket,
				Key:    key,
				// 可选：设置访问控制列表（ACL）
				// ACL: obs.AclPrivate, // 其他选项包括 obs.AclPublicRead, obs.AclPublicReadWrite 等
				// 可选：设置自定义元数据
				// Metadata: map[string]string{
				// 	"author":  "",
				// 	"project": "",
				// },
				// 可选：设置存储类型
				//StorageClass: obs.StorageClassStandard, // 其他选项包括 obs.StorageClassIA, obs.StorageClassArchive 等
				// 可选：设置网站重定向位置
				// WebsiteRedirectLocation: "http://www.example.com/",
				// 可选：设置过期时间（Unix 时间戳）
				// Expires: 1633036800,
				// 可选：设置内容的 MD5 校验
				// ContentMD5: "base64-encoded-md5",
				// 可选：设置内容长度
				// ContentLength: 123456789,
			},
		},
		// 设置 SourceFile 为要上传的本地文件路径
		SourceFile: sourceFile,
	}

	output, err := obsClient.PutFile(input)
	if err == nil {
		fmt.Printf("上传文件(%s)到存储桶(%s)成功！\n", input.Key, input.Bucket)
		fmt.Printf("StorageClass: %s, ETag: %s\n", output.StorageClass, output.ETag)
		return nil
	}

	fmt.Printf("上传文件(%s)到存储桶(%s)失败！\n", input.Key, input.Bucket)
	if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println("发生 OBS 错误，您的请求被 OBS 拒绝。")
		fmt.Println(obsError.Error())
	} else {
		fmt.Println("发生异常，客户端在与 OBS 通信时遇到内部问题。")
		fmt.Println(err)
	}
	return err
}
