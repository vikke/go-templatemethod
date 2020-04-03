package main

import "log"

type Script interface {
	Send()
}
type Script1 struct {
}
type Script2 struct {
}

type InOutLogger interface {
	In()
	Out()
}

type InOutLoggerImpl struct {
}

func (l *InOutLoggerImpl) In() {
	log.Println("in")
}
func (l *InOutLoggerImpl) Out() {
	log.Println("out")
}

func (m *Script1) Send() {
	log.Println("do script")
}

func (m *Script2) Send() {
	log.Println("do script2")
}

func main() {
	script := Script1{}
	boot(&script)

	script2 := Script2{}
	boot(&script2)
}

func boot(mailer Script) {
	SendMail(mailer)
}

func SendMail(script Script) {
	impl := InOutLoggerImpl{}
	impl.In()
	script.Send()
	impl.Out()
}
