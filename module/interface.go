package module

import (
	"log"
)

type Do interface {
	in()
	out()
	do()
}

type Receiver struct {
}

func New() Receiver {
	return Receiver{}
}

func (r *Receiver) in() {
	log.Println("in")
}

func (r *Receiver) out() {
	log.Println("out")
}

func (r *Receiver) do() {
	r.in()
	log.Println("do")
	r.out()
}
