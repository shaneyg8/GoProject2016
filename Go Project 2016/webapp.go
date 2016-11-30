package main

import (
// Make sure all packages are installed before rrunning the application
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"net/http"
	"fmt"
	
	
)
type List struct {
	// Below id won't be needed as it will produce Hex values if left in
	//Left in incase of change
 	//ID       bson.ObjectId `bson:"_id,omitempty"` 
 	Date    string
    ToDo    string
 }
 
 type ListTodo struct{
		ListStr	string		`json:"liststr" bson:"liststr"`
 }
 
 //Macaron
func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Post("/", save)
	m.Get("/Todo", func(ctx *macaron.Context){
		ctx.Data["List"] = search(nil,nil)
		ctx.HTML(303, "Todo")
	})
	m.Run(8080)
}

func save(w http.ResponseWriter, req *http.Request)string {
	fmt.Println("Uploadhandler start")
	//MongoDB locally... Mongo must be running on your machine
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
	//Giving the names of each "node" in the collection
 	c := session.DB("Todo").C("List")
 
 	// Insert
	//POSTing to database
	//Insert will input values in to that collection
 	err = c.Insert(&List{Date: date, ToDo: todo})
	
    if err != nil {
        panic(err)
    }
	response:="test"
	http.Redirect(w, req, "/Todo", 303)
	return response
	
}
//GET method to find the information from the database
//This information will then be passed to output on the application through the search function
func search(w http.ResponseWriter, req *http.Request)[]List {
	fmt.Println("Search called")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	

	defer session.Close()
	
	// Collection
	//This will print a List Array in the placeholder provided in the HTML
 	db := session.DB("Todo")
	c := db.C("List")
	var listArray []List
	//Finding the Data and producing all that is inside that collection
	err = c.Find(nil).Select(bson.M{"_id": 0}).All(&listArray)
	
    if err != nil {
        panic(err)
    }
	return listArray
}



//ISSUE
//Trying to Get data to show using a button in html in the placeholder
//Error appears when go build webapp.go is run
//Yet to fix
//Button is commented out on the Todo.html also
/*func find(w http.ResponseWriter, req *http.Request)[]List {
	fmt.Println("Search called")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	

	defer session.Close()
	
	response:="todo"
	http.Redirect(w, req, "/Todo", 303)
	return response
	
	// Collection
	//This will print a List Array in the placeholder provided in the HTML
 	db := session.DB("Todo")
	c := db.C("List")
	var listArray []List
	//Finding the Data and producing all that is inside that collection
	err = c.Find(nil).Select(bson.M{"_id": 0}).All(&listArray)
	
    if err != nil {
        panic(err)
    }
	return listArray
}*/