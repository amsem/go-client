package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
    Use: "details",
    Short: "This Project takes student info",
    Long: "A long String about description",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := cmd.PersistentFlags().Lookup("name").Value
        age := cmd.PersistentFlags().Lookup("age").Value
        log.Printf("Hello %s (%d), welcome to the cli world", name, age)
    },
} 




func Execute()  {
    rootCmd.PersistentFlags().StringP("name", "n", "stranger", "Name of the student")
    rootCmd.PersistentFlags().IntP("age", "a", 25, "Age of the student")
    err := rootCmd.Execute()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
