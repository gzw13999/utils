package utils

import (
	"os"
	"path/filepath"
	"strings"
	"time"
	"fmt"
	"io"
)

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	// Check if error is "no such file or directory"
	if _, ok := err.(*os.PathError); ok {
		return false
	}
	return false
}

func FileRemove(path string) bool {
	if !FileIsExist(path) {
		return false
	}
	err := os.Remove(path)
	if err != nil {
		return false
	}
	return true
}

func FileGetExt(file string) string {
	f := filepath.Ext(file)
	if f == "" {
		return f
	}
	return f[1:]
}

func GetCurrentDir()string{
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return ""
	} else {
		return strings.Replace(dir, "\\", "/", -1)
	}
}


func FileGetModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

func ReadFileByte(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(" file.Close() err:", err)
		}
	}()

	if fileByte, err := io.ReadAll(file); err != nil {
		return nil, err
	} else {
		return fileByte, nil
	}
}
