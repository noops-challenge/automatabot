package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rules struct {
	Name     string `json:"name"`
	Birth    []int  `json:"birth"`
	Survival []int  `json:"survival"`
}

type AutomataChallenge struct {
	Cells       [][]int `json:"cells"`
	Rules       Rules   `json:"rules"`
	Generations int     `json:"generations"`
}

type AutomataChallengeResponse struct {
	Challenge     AutomataChallenge `json:"challenge"`
	ChallengePath string            `json:"challengePath"`
}

func main() {
	var response = getChallenge()

	printChallenge(response.Challenge)
	printBoard("Before", response.Challenge.Cells)

	var cells = getSolution(response.Challenge)

	printBoard("After", cells)
}
func getSolution(challenge AutomataChallenge) [][]int {
	// your code goes here
	fmt.Println("Not yet implemented...")
	return challenge.Cells
}

func getChallenge() AutomataChallengeResponse {
	domain := "https://api.noopschallenge.com"

	// get and parse challenge from the api
	res, err := http.Get(domain + "/automatabot/challenges/start")
	if err != nil {
		panic(err.Error())
	}

	return parseApiResponse(res)
}

func parseApiResponse(res *http.Response) AutomataChallengeResponse {
	var response AutomataChallengeResponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err.Error())
	}

	return response
}

// print info about the ruleset we are using
func printChallenge(challenge AutomataChallenge) {
	fmt.Println()
	fmt.Println("Ruleset:", challenge.Rules.Name)
	fmt.Println("Birth:", challenge.Rules.Birth)
	fmt.Println("Survival:", challenge.Rules.Survival)
	fmt.Println("Generations:", challenge.Generations)
}

// print a board for human consumption, with borders
func printBoard(title string, cells [][]int) {
	fmt.Println()
	fmt.Println(title)
	for x := 0; x < len(cells[0])+2; x++ {
		fmt.Print("=")
	}
	fmt.Println()
	for y := range cells {
		fmt.Print("|")
		for x := range cells[y] {
			if cells[y][x] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("|")
		fmt.Println()
	}
	for x := 0; x < len(cells[0])+2; x++ {
		fmt.Print("=")
	}
	fmt.Println()
}
