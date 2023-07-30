package newPkg

import "log"

func Logger2(s string) {
	log.Printf("INFO: %s", s)
}
func LoggerWarn(s string) {
	log.Printf("Warn: %s", s)
}
