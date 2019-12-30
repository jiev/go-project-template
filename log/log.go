package log

import "fmt"

type LogConfig struct {

}

func Init(config *LogConfig) {

}

func Debug(format string,values ...interface{}) {
	fmt.Printf(format,values)
}

func Info(format string,values ...interface{}) {
	fmt.Printf(format,values)
}

func Error(format string,values ...interface{}) {
	fmt.Printf(format,values)
}