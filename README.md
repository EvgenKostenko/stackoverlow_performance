# MEV hackathon

Read [assigment](README_assignment.md.md)

### Dependencies
 
* Golang https://golang.org/
* GOPATH should be setted [GOPATH DOC](https://golang.org/doc/code.html#GOPATH)
* MongoDB 3~ https://www.mongodb.com/

### Installation

* ``` go get github.com/EvgenKostenko/stackoverlow_performance ```
* ``` cd $GOPATH/src/github.com/EvgenKostenko/stackoverlow_performance/ ```
* Download zip archive from [dump](https://drive.google.com/drive/folders/0B4edT1hr54MWTFc1WXEwY3J3ZVU?usp=sharing)
* Extract zip with dump
* Set in config/config.yml path to dump By default it ``` scifi.stackexchange.com/ ``` means dump extracted in project directory
* Set database name and database host By default it ``` localhost ``` and ``` SO ```
* Run ``` go run ./importer.go ``` for importing data from xml to mongodb
* Run ``` go run ./main.go ``` for run API

### Test

* Check tags ``` curl -X "GET" "http://localhost:8080/toptags/373" ```
* Check users ``` curl -X "GET" "http://localhost:8080/topusers/star" ```


