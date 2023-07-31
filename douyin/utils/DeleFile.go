package utils

import (
	"fmt"
	"os"
)

func DeleteFile() {
	path, _ := os.Getwd()
	path += "/upload/videos/"
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("删除失败")
	}
	fmt.Println("删除成功")
}
