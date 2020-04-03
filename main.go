package main

import (
	m "github.com/vikke/go-templatemethod/module"
)

func main() {
	script1 := m.Script1{}
	m.SendMail(script1)

	script2 := m.Script2{}
	m.SendMail(script2)
}
