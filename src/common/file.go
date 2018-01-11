package common

import (
	"os"
	"path/filepath"

	"github.com/qiniu/log.v1"
)

// 读取或新建一个文件，用于读写
// 同时会创建相应的目录
func OpenOrCreateFile(filename string) (*os.File, error) {
	fileExists, err := IsFileExists(filename)
	if err != nil {
		return nil, err
	}

	var flag int
	if fileExists {
		flag = os.O_RDWR
	} else {
		if err = os.MkdirAll(filepath.Dir(filename), 0775); err != nil {
			return nil, err
		}
		flag = os.O_RDWR | os.O_CREATE | os.O_EXCL | os.O_TRUNC
	}
	file, err := os.OpenFile(filename, flag, 0666)
	if err != nil {
		log.Errorf("Failed to open file, fileExsits: %v, err: %v", fileExists, err)
		return nil, err
	}
	return file, nil
}

func IsFileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
