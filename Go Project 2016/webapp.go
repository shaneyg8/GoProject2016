package main

import (
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"net/http"
	//"fmt"
	
	
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Post("/", save)
	m.Get("/search/:id", func(ctx *macaron.Context) {
		// Adapted from: https://go-macaron.com/docs/middlewares/templating
		//ctx.Data["id"] = search(ctx.Params(":id"))
		//ctx.HTML(200, "hello")
	})
	m.Run(4000)
}

func save(req *http.Request) {
	//fmt.Println("Uploadhandler start")
	session, err := mgo.Dial("127.0.0.1:50864")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	
}