package level_log

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// TODO:
	//log.Println("level_log init")
}

func Info(message string) {
	log.Println("INFO ", message)
}

func Infof(format string, a ...interface{}) {
	log.Println("INFO ", fmt.Sprintf(format, a...))
}

func Fatal(message string) {
	log.Fatalln("FATAL", message)
}

func Fatalf(format string, a ...interface{}) {
	log.Fatalln("FATAL", fmt.Sprintf(format, a...))
}
