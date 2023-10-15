package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main()  {
    app := cli.NewApp()
    app.Name = "MARKS"
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "save",
            Value: "no",
            Usage: "do you want to save to DB ? (yes/no)",
        },
    }
      app.Version = "1.0"
      app.Action = func(c *cli.Context) error {
            var args []string
            if c.NArg() > 0 {
                args = c.Args()
                personName := args[0]
                x := len(args)
                marks := args[1:x]
                log.Println("Person : ",personName)
                log.Println("Marks : ",marks)
            }
            if c.String("save") == "no" {
                log.Println("Skipping saving to the DB")
            }else {
                log.Println("Saving to the DB ", args)
            }
            return nil
        }
    
    app.Run(os.Args)
}
