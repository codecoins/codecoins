package log

import (
	"fmt"
	"time"
	"strings"
)

var logLevel string

func logFormatter(level string) string {
	t := time.Now()
	wrapper := func() string {
		switch(level) {
		case "debug":
			return "\033[89m   DEBUG\t[%v] \033[0m"
		case "warning":
			return "\033[90m   WARNING\t[%v] \033[0m"
		case "error":
			return "\033[91m   ERROR\t[%v] \033[0m"
		case "green":
			return "\033[91m   GREEN\t[%v] \033[0m"
		default:
			return "\033[94m   INFO\t[%v] \033[0m"
		}
	}()
	timeString := t.Format("01/02/2006 15:04:05 MST")
	logString := fmt.Sprintf(wrapper, timeString)

	return logString + "%v\n"
}

func SetLevel(s string){
	logLevel = s
}

func Debug(s string) error {
	var level = "debug"
	return printLog(level,s)
}

func Info(s string) error {
	var level = "info"
	return printLog(level,s)
}

func Warning(s string) error {
	var level = "warning"
	return printLog(level,s)
}

func Error(s string) error {
	var level = "error"
	return printLog(level,s)
}

func Green(s string) error {
	var level = "green"
	return printLog(level,s)
}

func printLog(level,s string) error {
	var err error
	if strings.Contains(logLevel,level) {
		_, err = fmt.Printf(logFormatter(level), s)
	}
	return err
}

func PrintError(err error) error {
	if err != nil {
		Error(err.Error())
	}
	return err
}

func DieFatal(err error){
	if err != nil {
		panic(Error(err.Error()))
	}
}