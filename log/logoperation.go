package log

import (
	"io"
	"log"
	"os"
)

var (
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
)

func init() {
	file, err := os.OpenFile("src/go.test/log/log_go.log4g", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	multi := io.MultiWriter(file, os.Stdout)

	INFO = log.New(multi, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARNING = log.New(multi, "WRN: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Logger() bool {
	INFO.Println("program running...")
	WARNING.Println("program running...")
	return true
}
