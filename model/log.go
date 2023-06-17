package model

import (
	"github.com/sirupsen/logrus"
	"ls-kh-download/log"
)

var myLog = log.Log

func init() {
	myLog.SetLevel(logrus.TraceLevel)
}
