package module

import "log"

type Mailer struct {

}

type InOutLogger interface {
	In()
	Out()
}
type InOutLoggerImpl struct {
}

(l *InOutLoggerImple)In() {
	log.Println("in")
}
(l *InOutLoggerImple)Out() {
	log.Println("out")
}

func SendMail(mailer Mailer) {
	impl := InOutLoggerImpl{}
	impl.In()
	Mailer.send()
	imple.Out()
}
