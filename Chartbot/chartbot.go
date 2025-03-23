package main

import (
	"fmt"
	"strings"
)

func main() {

	// create reponse detail using arreys

	AnswerResponse := map[string]string{
		"hello":   "hi there  can i assist you today",
		"who created you":"Issa mwanga aka lynxPrazoo",
		"default": "I did not understand you",
		"goodbye": "Goodbye to have nice day",
	}

	fmt.Println("Welcome to Prazoo Chartbot")
	for {
		fmt.Println("Can i assist you: ")
		var requestuser string

		fmt.Scanln(&requestuser)
		requestuser = strings.ToLower(requestuser)

		response, ok := AnswerResponse[requestuser]

		if !ok {
			response = AnswerResponse["default"]
			fmt.Println("chartPrazoo", response)
		}else if ok {
			fmt.Println("chartPrazoo: ",response)
			if ( response == AnswerResponse["goodbye"]){
				break

			}
		}
	}

}
