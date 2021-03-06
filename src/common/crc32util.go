package common

import (
	"github.com/qiniu/io/crc32util"
	"io"
	"os"
	"qiniupkg.com/x/log.v7"
	"hash/crc32"
)

func TransferToCrc32EncodedFile(dstFile string, srcFile string) error {
	src, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	dst, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	enc := crc32util.SimpleEncoder(src, nil)
	n, err := io.Copy(dst, enc)
	log.Info("file size: %v", n)
	return err
}

func GetFileCrc32Value(file string) (uint32, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return 0, err
	}

	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	value := h.Sum32()
	return value, nil
}