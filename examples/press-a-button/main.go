package main

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/theclapp/hvue"
)

type Data struct {
	*js.Object
	Message string `js:"message"`
}

func main() {
	hvue.NewVM(
		hvue.El("#app-5"),
		hvue.DataS(NewData("Hello, Vue!")),
		hvue.MethodsOf(&Data{}))
}

func NewData(message string) *Data {
	d := &Data{Object: js.Global.Get("Object").New()}
	d.Message = message
	return d
}

func (d *Data) ReverseMessage(event *js.Object) {
	// event ignored
	d.Message = reverse(d.Message)
}

func reverse(s string) string {
	runes := strings.Split(s, "")
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return strings.Join(runes, "")
}
