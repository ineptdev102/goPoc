package sample

import "log"

func Logger(s string) {
	log.Printf("INFO: %s", s)
}

func LogFatal(s string) {
	log.Printf("Fatal_tes: %s", s)
}
