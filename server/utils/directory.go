package utils

import (
	"errors"
	"os"
)

func IsPathExists(path string) (bool, error) {
	// 获取目录信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			// 不存在文件夹
			return false, nil
		}
		return false, err
	}
	if fileInfo.IsDir() {
		// 存在文件夹
		return true, nil
	} else {
		// 存在同名文件
		return false, errors.New("存在同名文件")
	}
}
