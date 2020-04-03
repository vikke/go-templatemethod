package template

import "log"

func (Script1) Exec() {
	log.Println("Script1")
}
func (Script2) Exec() {
	log.Println("Script2")
}

func SendMail(script Script) {
	log.Println("IN")
	script.Exec()
	log.Println("OUT")
}
