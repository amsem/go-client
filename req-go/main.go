package main

import (
	"log"

	"github.com/levigross/grequests"
)


func main()  {
    resp, err := grequests.Get("https://httpbin.org/get", nil)
    if err != nil {
        log.Fatalln("Unable to make request : ", err)
    }
    var jsonRes interface{}
    resp.JSON(&jsonRes)
   // log.Println(resp.String())
    log.Println(jsonRes)
}

