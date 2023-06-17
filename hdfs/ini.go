package hdfs

import (
	"github.com/colinmarc/hdfs"
	"github.com/spf13/viper"
	"ls-kh-download/log"
)

var Client *hdfs.Client

func init() {
	client, err := hdfs.New(viper.GetString("Hdfs.NameNode"))
	if err != nil {
		log.Log.Error("hdfs init failed,err: ", err)
		return
	} else {
		log.Log.Info("hdfs init successfully")
	}
	Client = client
}
