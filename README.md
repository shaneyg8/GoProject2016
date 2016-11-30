# Go-Project-2016
## Abstract
A simple ToDo List web application using Golang and HTML. In this application you will be given two text boxes where you enter a Date and the other text box will be used to enter your activity you want to do. This information will be then posted to a mongodb to store your information and then retrieved to show you what you your activities you are about to do. 

##GO & MACARON
![Macaron](https://raw.githubusercontent.com/go-macaron/macaron/v1/macaronlogo.png) ![Golang](https://2.bp.blogspot.com/-Yt5QMzQYEoQ/V-Y48KZ0NyI/AAAAAAAAAAY/7_cgxeVqIb8GcVMibMLuZSIgh7O8p9zjQCLcB/s1600/go.png)
* The application in this repository is written in the programming language [Go](https://golang.org/). 
* Macaron is required when running this application and can be installed through the simple steps [here](https://go-macaron.com/). 
* Macaron must be installed in your [GOPATH](https://golang.org/doc/code.html#GOPATH).

As long as those steps are followed you should be able to run the application, the project can be compiled and run with:

`$ go run webapp.go`

## Bootstrap
![Bootstrap Logo](http://www.allosamerica.com/wp-content/uploads/2015/10/bootstrap-logo.png)

For a more descriptive guide on how to install bootstrap follow the link [here](http://getbootstrap.com/getting-started/).

* Installing with bower:
 
`$ bower install bootstrap`

* Installing with npm:

`$ npm install bootstrap@3`

Taken from the bootstrap getting started page which can be clicked on the link above :

"require('bootstrap') will load all of Bootstrap's jQuery plugins onto the jQuery object. The bootstrap module itself does not export anything. You can manually load Bootstrap's jQuery plugins individually by loading the /js/*.js files under the package's top-level directory.

Bootstrap's package.json contains some additional metadata under the following keys:

* less - path to Bootstrap's main Less source file
* style - path to Bootstrap's non-minified CSS that's been precompiled using the default settings (no customization)"

* You can also install bootstrap with composer:

`$ composer require twbs/bootstrap`

* Once downloaded unzip the folder to view the structure which should look something like this:
```
    bootstrap/
    ├── css/
    │   ├── bootstrap.css
    │   ├── bootstrap.css.map
    │   ├── bootstrap.min.css
    │   ├── bootstrap.min.css.map
    │   ├── bootstrap-theme.css
    │   ├── bootstrap-theme.css.map
    │   ├── bootstrap-theme.min.css
    │   └── bootstrap-theme.min.css.map
    ├── js/
    │   ├── bootstrap.js
    │   └── bootstrap.min.js
    └── fonts/
        ├── glyphicons-halflings-regular.eot
        ├── glyphicons-halflings-regular.svg
        ├── glyphicons-halflings-regular.ttf
        ├── glyphicons-halflings-regular.woff
        └── glyphicons-halflings-regular.woff2
```
## MongoDB -Backend
![MongoDBLogo](https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/640px-MongoDB-Logo.svg.png)

### [Advantages](https://www.tutorialspoint.com/mongodb/mongodb_advantages.htm)

* Uses internal memory for storing the (windowed) working set, enabling faster access of data
* Structure of a single object is clear
* MongoDB is a document database in which one collection holds different documents. Number of fields, content and size of the document can differ from one document to another
* Replication and high availability
* Index on any attribute
* Fast in-place updates

### My Web Application

As I have never used mongoDB before I felt like this was a new challenge. Having already worked with MySQL I wanted to see what was so different to it. With my web application I wanted to use a database so I could POST and GET data. As it was a ToDo List all I wanted to incorporate in this project was two text boxes which would be all that you would need, a date you wanted to do the activity and another text box to input your activity, this would then be POSTing data to my Mongo database through the save button. The POST and GET functions are set up through the GoLang which can be seen on my source code. 

Skip to content
This repository
Search
Pull requests
Issues
Gist
 @shaneyg8
 Unwatch 2
  Star 0
  Fork 0 shaneyg8/GoProject2016
 Code  Issues 0  Pull requests 0  Projects 1  Wiki  Pulse  Graphs  Settings
 Delete Page Page History New Page
Editing 5 . Project Architecture
 
5 . Project Architecture
Write  Preview
h1 h2 h3      B i        Edit mode:  
Block Elements
Span Elements
Miscellaneous
Paragraphs & Breaks
Headers
Blockquotes
Lists
Code Blocks
Horizontal Rules
To create a paragraph, simply create a block of text that is not separated by one or more blank lines. Blocks of text separated by one or more blank lines will be parsed as paragraphs.

If you want to create a line break, end a line with two or more spaces, then hit Return/Enter.


# Project Architecture

[![Architecture.png](https://s21.postimg.org/yn3avt20n/Architecture.png)](https://postimg.org/image/rjvfg6wkz/)

This application is written in GoLang using the Macaron Framework.The backend technology used is MongoDB and is ran locally. This application using GoLang POSTs and retrieves data with the GET method through the service layer(GoLang). The Front End/ Client Side is done through HTML5 and Bootstrap. 
