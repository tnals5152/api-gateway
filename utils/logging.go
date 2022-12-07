package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var logger *log.Logger
var LogFile *os.File
var today string

const DEFAULT = iota + 3

func Error(err error) {
	logging(err)
}

func Panic(err error) {
	logging(err)
	panic(err)
}

func logging(err error) {
	if err == nil {
		return
	}

	now := time.Now()
	nowString := now.Format("2006-01-02")

	if today != nowString {
		SetLog()
	}

	logFuncName, _, _, _ := runtime.Caller(DEFAULT - 2)
	funcName1, file1, line1, ok1 := runtime.Caller(DEFAULT)
	funcName2, file2, line2, ok2 := runtime.Caller(DEFAULT - 1)

	logNames := strings.Split(runtime.FuncForPC(logFuncName).Name(), ".")
	logger.SetPrefix(fmt.Sprintf("[%s]", strings.ToUpper(logNames[len(logNames)-1])))

	funcName2Split := strings.Split(runtime.FuncForPC(funcName2).Name(), ".")
	funcName1Split := strings.Split(runtime.FuncForPC(funcName1).Name(), ".")

	if ok1 || ok2 {
		logger.Printf("[%s][%s](%d line)/[%s][%s](%d line) - %s",
			filepath.Base(file2), funcName2Split[len(funcName2Split)-1], line2,
			filepath.Base(file1), funcName1Split[len(funcName1Split)-1], line1,
			err.Error())
	}
}

func SetLog() *os.File {
	var err error

	now := time.Now()
	today = now.Format("2006-01-02")

	if LogFile != nil { // 매일 매일 새로운 로그 파일에 쓰기 위해 이전 날짜에 열었던 파일이 있으면 닫아주고 시작
		err = LogFile.Close()
		Error(err)
	}

	LogFile, err = os.OpenFile(fmt.Sprintf("./log/%s.log", today), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)

	if err != nil {
		Panic(err)
	}

	logger = log.New(LogFile, "", log.Ldate|log.Ltime)

	return LogFile
}
