package router

import (
	"archive/zip"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"ls-kh-download/dao"
	"ls-kh-download/hdfs"
	"ls-kh-download/log"
	"os"
	"strings"
)

var (
	myLog  = log.Log
	db     = dao.DB
	client = hdfs.Client
)

func InitRouters() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/download", JWT(), func(c *gin.Context) {
		fpath := c.Query("filepath")
		file, err2 := os.Create("file.zip")
		if handleError(c, err2) {
			return
		}
		defer file.Close()
		zWriter := zip.NewWriter(file)
		defer zWriter.Close()
		if !strings.HasPrefix(fpath, "/") {
			fpath = "/" + fpath
		}
		reader, err := client.Open("/" + c.GetString("username") + fpath)
		if handleError(c, err) {
			return
		}
		defer reader.Close()
		strArray := strings.Split(fpath, "/")
		if len(strArray) < 2 {
			jsonError(c, errors.New("invalid filepath: "+fpath))
		}
		zInnerWriter, err := zWriter.Create("file/" + strArray[len(strArray)-1])
		if handleError(c, err) {
			return
		}
		_, err2 = io.CopyBuffer(zInnerWriter, reader, make([]byte, 1024))
		if handleError(c, err2) {
			return
		}
		zWriter.Close()
		zipFile, err := os.Open("file.zip")
		if handleError(c, err) {
			return
		}
		defer zipFile.Close()
		_, err = io.CopyBuffer(c.Writer, zipFile, make([]byte, 1024))
		if handleError(c, err) {
			return
		}
	})
	_ = r.Run(":" + viper.GetString("Router.Port"))
}
