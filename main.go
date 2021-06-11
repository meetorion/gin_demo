package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed,err:%v", err)
		return
	}

	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render templete failed,err=%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v", err)
		return
	}

}
