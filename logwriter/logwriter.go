package logwriter

import (
	"path/filepath"
	"log"
	"os"
)

//Logger custom logger
type Logger struct {
	Path  string
	Level string
	File  *os.File
}

//Log custom logger
var Log = &Logger{}

//Init logger initialization
func (l *Logger) Init() (err error) {

	dir := filepath.Dir(l.Path)
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm)
		}
	}

	f, err := os.OpenFile(l.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	//defer f.Close()

	log.SetOutput(f)
	l.File = f
	return err
}

//Debug Debug
func (l *Logger) Debug(v ...interface{}) {
	log.Println("DEBUG:", v, "\r\n")
}

//Error Error
func (l *Logger) Error(v ...interface{}) {
	log.Println("ERROR:", v, "\r\n")
}
