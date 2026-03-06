package logger

import (
	"log"
	"os"
)

func New() *log.Logger {
	return log.New(os.Stdout, "[site-sentinel] ", log.LstdFlags|log.Lshortfile)
}
