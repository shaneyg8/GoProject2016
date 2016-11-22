package main

import (
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"net/http"
	"fmt"
	
	
)
type List struct {
 	ID       bson.ObjectId `bson:"_id,omitempty"`
 	Date    string
    ToDo    string
 }
func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Post("/", save)
	m.Run(4000)
}

func save(req *http.Request) {
	//fmt.Println("Uploadhandler start")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()
    req.ParseForm()
    date := req.FormValue("date")
    todo := req.FormValue("todo")
	//Check if there is an error
	if err != nil {
		fmt.Println(err)
	}
    
    // Collection
 	c := session.DB("Todo").C("List")
 
 	// Insert
 	err = c.Insert(&List{Date: date, ToDo: todo})
    if err != nil {
        panic(err)
    }
	
}