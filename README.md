# Go-Project-2016
## Shane Gleeson - G00311793
## [WIKI LINK FOR IN DEPTH DOCUMENTATION](https://github.com/shaneyg8/GoProject2016/wiki)

## Abstract
A simple ToDo List web application using Golang and HTML. In this application you will be given two text boxes where you enter a Date and the other text box will be used to enter your activity you want to do. This information will be then posted to a mongodb to store your information and then retrieved to show you what you your activities you are about to do. 

## [PROJECT INSTRUCTIONS ON HOW TO RUN](https://github.com/shaneyg8/GoProject2016/blob/master/Go%20Project%202016/README.md)

##Project Planning
As part of this Go Project it originally was designed as a group project but with regards differences between myself and my working colleague I decided to do the project by myself.

-> As part of my brainstorming I had come up with a few ideas to do for a project and wanted to produce something that could improve my knowledge on the Go programming language as well as using another Database other than MySQL

At first Frameworks that could be used with GoLang:
* Macaron
* Gorilla
* Beego
* Martini
* GoCraft

-> I decided to go for Macaron just because I had used Macaron in previous labs and I was familiar with it.

When coming to the Front End side of things I was always going to use HTML5 as I wanted to incorporate languages that I had done before with GoLang. Along with HTML5 I wanted to use a Front End Framework.

After researching I had a list to pick from:
* Bootstrap
* Zurb
* Semantic UI
* Pure By Yahoo

-> Because it has grown rapidly in popularity I decided to use bootstrap as I had never used it before and wanted to see what everyone was talking about. 

Finally as part of the technologies I were to use I wanted to pick a database that I have never worked on before.
* MongoDB
* CouchDB
* NoSQL
* MySQL

-> I left MySQL in my options as if I couldn't understand how to work with any other database in time I know i'd have something to fall back on. I ended up choosing MongoDB.

##Ideas
For Ideas for an Web Application I had many in mind, due to this originally being a group project I wanted to create something a small bit more elegant with ideas such as:
* Mobile Banking Application with incorporated Login Feature 
* File Sharing Dropbox with incorporated Login Feature

But because I am now doing this project on my own I had decided I wouldn't have been able to get that amount of work done in so much little time so because none of them Ideas fell through I started producing Ideas for myself and the work i'd be able to fit in in such a small space of time, these ideas included:
* Counties Application where it retrieves info on all counties of Ireland from database
* ToDo List POSTing and GETting information from the database
* A simple Authentication Page 

-> In the end I decided to go for a ToDo List as in my opinion it had features that I would like to learn in the Go Programming Language. I felt it would be a challenging task but I felt it would enhance my knowledge and also I felt I would be able to get the work done in that short time frame.  

# MongoDB
### My Web Application

As I have never used mongoDB before I felt like this was a new challenge. Having already worked with MySQL I wanted to see what was so different to it. With my web application I wanted to use a database so I could POST and GET data. As it was a ToDo List all I wanted to incorporate in this project was two text boxes which would be all that you would need, a date you wanted to do the activity and another text box to input your activity, this would then be POSTing data to my Mongo database through the save button. The POST and GET functions are set up through the GoLang which can be seen on my source code. 


# Project Architecture

[![Architecture.png](https://s21.postimg.org/yn3avt20n/Architecture.png)](https://postimg.org/image/rjvfg6wkz/)

This application is written in GoLang using the Macaron Framework.The backend technology used is MongoDB and is ran locally. This application using GoLang POSTs and retrieves data with the GET method through the service layer(GoLang). The Front End/ Client Side is done through HTML5 and Bootstrap. 
