package loglib

import (
"log"
"os"
"path"
)

var Mylog *log.Logger
var Log_file_dir string
var Log_pw bool

func init()  {
	process_name := path.Base(os.Args[0])
	if Log_file_dir == ""{
		if Log_pw {
			if Log_file_dir == "" {
				Log_file_dir = process_name+".log"
			}
			logFile, err := os.OpenFile(Log_file_dir, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			Mylog = log.New(logFile, "["+process_name+"] ", log.Ldate|log.Ltime|log.LstdFlags)
		}else {
			Mylog = log.New(os.Stdout, "["+process_name+"] ", log.Ldate|log.Ltime|log.LstdFlags)
		}
	}
}