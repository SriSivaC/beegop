package main

import (
	"fmt"
	"log"

	beegoServerWeb "github.com/beego/beego/v2/server/web"
)

type mainController struct {
	beegoServerWeb.Controller
}

func (this *mainController) Get() {
	bytesx, err := fmt.Fprint(this.Ctx.ResponseWriter, "<h1>Testing My Go Skills</h1>")
	log.Println(":INFO :Bytes written - ", bytesx)
	warnError(&err)
}

type supportController struct {
	beegoServerWeb.Controller
}

func (this *supportController) Get() {
	bytesx, err := fmt.Fprint(this.Ctx.ResponseWriter, "<h1>Support</h1> <p>To get in touch, please <a href=\"mailto:example@picgallery.com\">send an email</a></p>")
	log.Println(":INFO :Bytes written - ", bytesx)
	warnError(&err)
}
func main() {
	beegoServerWeb.Router("/", &mainController{})
	beegoServerWeb.Router("/support", &supportController{})
	beegoServerWeb.Run()
}

func warnError(err *error) {
	if *err != nil {
		log.Println(":WARN :Error encounter - ", *err)
		// panic(fmt.Sprintf("Panicking at Handler"))
	} else if *err == nil {
		log.Println(":INFO :No Error encounter - ", *err)
	}
}
