# GO Project Instructions
## TODO List
### Shane Gleeson - G00311793
#### Download the project or clone it to your desktop


**In command prompt:**
You must set your GOPATH to your workspace for example:
```
set GOPATH=C:\Users\SHANE\Desktop\GO2016
```
This "GO2016" as being your folder for all your work that you're about to do

**After you set your GOPATH you must install the following packages**
```
go get gopkg.in/macaron.v1
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```

If mongo is installed correctly then run the command below. If you need to install mongo follow the procedure in my wiki
```
mongod
```

Then opening up a new cmd window run:
```
mongo
```
In a new command prompt window without closing the others go to the directory of the project using the command:
```
go run webapp.go
```
or
```
go run webapp.exe
```



#### References
https://www.mongodb.com/ - MongoDB
