/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "This is the a cli tool to get data jokes",
	Long: `this is the long desction for the project`,
	Run: func(cmd *cobra.Command, args []string) {
    getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type joke struct {
  ID     string `json:"id"`
  Joke   string `json:"joke"`
  status int    `json:"status"`
}

func getRandomJoke(){
  url := "https://icanhazdadjoke.com/"
  responseBytes := getJokeData(url)
  joke := joke{}
  err := json.Unmarshal(responseBytes, &joke)
  if err != nil {
    fmt.Println("Error parsing the response - %v", err)
      }
  fmt.Printf("%s\n", joke.Joke)
}

func getJokeData(baseAPI string)[]byte{
  request, err := http.NewRequest(
    "GET",
    baseAPI,
    nil,
  )
  if err != nil {
    fmt.Println("Error getting the request - %v", err)
  }
  request.Header.Set("Accept", "application/json")
  request.Header.Set("User-Agent", "Dad Joke CLI (github.com/basola21/dadjoke-cli)")

  response, err := http.DefaultClient.Do(request)
  if err != nil {
    fmt.Println("Error getting the response - %v", err)
      }

  responseBytes, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println("Error reading the response body - %v", err)
  }
  return responseBytes
}
