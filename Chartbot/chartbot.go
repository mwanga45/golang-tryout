// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	// create reponse detail using arreys

// 	AnswerResponse := map[string]string{
// 		"hello":   "hi there  can i assist you today",
// 		"who created you":"Issa mwanga aka lynxPrazoo",
// 		"default": "I did not understand you",
// 		"goodbye": "Goodbye to have nice day",
// 	}

// 	fmt.Println("Welcome to Prazoo Chartbot")
// 	for {
// 		fmt.Println("Can i assist you: ")
// 		var requestuser string

// 		fmt.Scanln(&requestuser)
// 		requestuser = strings.ToLower(requestuser)

// 		response, ok := AnswerResponse[requestuser]

// 		if !ok {
// 			response = AnswerResponse["default"]
// 			fmt.Println("chartPrazoo", response)
// 		}else if ok {
// 			fmt.Println("chartPrazoo: ",response)
// 			if ( response == AnswerResponse["goodbye"]){
// 				break

// 			}
// 		}
// 	}

// }
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Response structure holds patterns and possible answers
type Response struct {
	Patterns []string
	Answers  []string
}

var dataset = []Response{
	{
		Patterns: []string{"hello", "hi", "hey"},
		Answers:  []string{"Hello!", "Hi there!", "Hey!"},
	},
	{
		Patterns: []string{"how are you", "how's it going"},
		Answers:  []string{"I'm just a program, but I'm functioning well!", "All systems go!"},
	},
	{
		Patterns: []string{"bye", "goodbye"},
		Answers:  []string{"Goodbye!", "See you later!", "Bye!"},
	},
	{
		Patterns: []string{"help", "problem"},
		Answers:  []string{"I can help with basic questions. Try asking about greetings or say goodbye.", "What can I help you with?"},
	},
	{
        Patterns: []string{"vp"},
		Answers: []string{"Kama kawa niaje"},
	},
}

// Match user input to responses
func findResponse(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))

	for _, resp := range dataset {
		for _, pattern := range resp.Patterns {
			if strings.Contains(input, pattern) {
				return resp.Answers[0] // Return first matching answer
			}
		}
	}
	return "I'm not sure I understand. Can you rephrase that?"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Basic Chatbot (type 'exit' to quit)")

	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if strings.ToLower(userInput) == "exit" {
			break
		}

		fmt.Println("Bot:", findResponse(userInput))
	}
}
