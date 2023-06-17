package main

import (
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"ls-kh-download/hdfs"
	"os"
)

func main() {
	hdfs.Client.ReadFile("/2/hosts.txt")
}

func CompressFile(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil || info.IsDir() {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = prefix + "/" + header.Name
	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}
	if _, err = io.Copy(writer, file); err != nil {
		return err
	}
	return nil
}

func main02() {
	f, _ := os.Open("main.go")
	// 压缩文件
	dst, _ := os.Create("test.zip")
	zipWriter := zip.NewWriter(dst)
	if err := CompressFile(f, "", zipWriter); err != nil {
		log.Fatalln(err.Error())
	}
	// Make sure to check the error on Close.
	if err := zipWriter.Close(); err != nil {
		log.Fatalln(err.Error())
	}
	return
}
func main0() {
	r := gin.Default()
	r.GET("/download", func(c *gin.Context) {
		f1, err := os.Create("f1.zip")
		if err != nil {
			fmt.Println(err)
			return
		}
		zow := zip.NewWriter(f1)
		ziw, err := zow.Create("t/a.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		f2, err := os.Open("./main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(ziw, f2)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = zow.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok")
		f1.Close()
	})
	r.Static("/file", ".")
	_ = r.Run()
}
