package main

import (
	"fmt"
	"os"
	"strconv"
)

func readableFileSize(size int64) string {
	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case size >= TB:
		return strconv.FormatFloat(float64(size)/TB, 'f', 2, 64) + " TB"
	case size >= GB:
		return strconv.FormatFloat(float64(size)/GB, 'f', 2, 64) + " GB"
	case size >= MB:
		return strconv.FormatFloat(float64(size)/MB, 'f', 2, 64) + " MB"
	case size >= KB:
		return strconv.FormatFloat(float64(size)/KB, 'f', 2, 64) + " KB"
	default:
		return strconv.FormatInt(size, 10) + " B"
	}
}

func GetFileSize(filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	fileSize := fileInfo.Size()
	return readableFileSize(fileSize), nil
}
