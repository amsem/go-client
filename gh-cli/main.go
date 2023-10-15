package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/levigross/grequests"

	"github.com/urfave/cli"
)



var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

type File struct {
    Content string `json:"content"`
}
type Gist struct {
    Description string `json:"description"`
    Public bool `json:"public"`
    Files map[string]File `json:"files"`
}
type Repo struct {
    ID int `json:"id"`
    Name string `json:"name"`
    FullName string `json:"full_name"`
    Forks int `json:"forks"`
    Private bool `json:"private"`
}

func getStats(url string) *grequests.Response {
    resp, err := grequests.Get(url, requestOptions)
    if err != nil {
        log.Fatalln("Unable to make req : ", err)
    }
    return resp
}

func createPost(url string, args []string ) *grequests.Response {
    description := args[0]
    var fileContents = make(map[string]File)
    for i:= 1; i < len(args); i++ {
        dat, err := os.ReadFile(args[i])
        if err != nil {
            log.Println("Please check the fileNames. Absolute path (or) same directory are allowed")
            return nil
        }
        var file File
        file.Content = string(dat)
        fileContents[args[i]] = file
    }
    var gist = Gist{Description: description, Public: true, Files: fileContents}
    var postBody, _ = json.Marshal(gist)
    var requestOptions_copy = requestOptions
    requestOptions_copy.JSON = string(postBody)
    resp, err := grequests.Post(url, requestOptions_copy)
    if err != nil {
        log.Println("Craete Request dailed for github API")
    }
    return resp
}


func main()  {
    app := cli.NewApp()
    // let's define the commands of our CLI app
    app.Commands = []cli.Command{
        {
            Name: "fetch",
            Aliases: []string{"f"},
            Usage: "Fetch the Repo details. [Usage]: githubAPI fetch user",
            Action: func(c *cli.Context) error {
                if c.NArg() > 0 {
                    var repos []Repo
                    user := c.Args()[0]
                    var repoURL = fmt.Sprintf("https://api.github.com/users/%s/repos", user)
                    resp := getStats(repoURL)
                    resp.JSON(&repos)
                    log.Println(repos)
                }else {
                    log.Println("Please give a valid user name. See -h to get help")
                }
              return nil  
            },
                
        },
        {
            Name : "create",
            Aliases: []string{"c"},
            Usage: "Create a gist from the given text . [Usage]: githubAPI name 'description' sample.txt",
            Action: func(c *cli.Context) error {
                if c.NArg() > 1 {
                    args := c.Args()
                    var postURL = "https://api.github.com/gists"
                    resp :=  createPost(postURL, args)
                    log.Println(resp.String())
                }else {
                    log.Println("please sufficient arguments . See -h for help")
                }
                return nil
            },
        },
    }
    app.Version = "0.1"
    app.Run(os.Args)
}
