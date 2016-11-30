package main

import (
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"net/http"
	"fmt"
	
	
)
type List struct {
 	//ID       bson.ObjectId `bson:"_id,omitempty"`
 	Date    string
    ToDo    string
 }
 
 type ListTodo struct{
		ListStr	string		`json:"liststr" bson:"liststr"`
 }
 
func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Post("/", save)
	m.Get("/Todo", func(ctx *macaron.Context){
		ctx.Data["List"] = search(nil,nil)
		ctx.HTML(303, "Todo")
	})
	m.Run(4000)
}

func save(w http.ResponseWriter, req *http.Request)string {
	fmt.Println("Uploadhandler start")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	

	defer session.Close()
	

	err = req.ParseForm()
	if err != nil {
		panic(err)
	}
    	date := req.FormValue("date")
    	todo := req.FormValue("todo")
    
    // Collection
 	c := session.DB("Todo").C("List")
 
 	// Insert
 	err = c.Insert(&List{Date: date, ToDo: todo})
	
    if err != nil {
        panic(err)
    }
	response:="test"
	http.Redirect(w, req, "/Todo", 303)
	return response
	
}

func search(w http.ResponseWriter, req *http.Request)[]List {
	fmt.Println("Search called")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	

	defer session.Close()
	
	// Collection
 	db := session.DB("Todo")
	c := db.C("List")
	var listArray []List
	err = c.Find(nil).Select(bson.M{"_id": 0}).All(&listArray)
	
    if err != nil {
        panic(err)
    }
	return listArray
}