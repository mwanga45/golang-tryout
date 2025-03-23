package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please type your content:")
	contentReader := bufio.NewReader(os.Stdin)
	content, err := contentReader.ReadString('\n')
	content = strings.TrimSpace(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Please enter the name of your file:")
	//create variable that will hold are variable from input (user)
	filenameReader := bufio.NewReader(os.Stdin)
	// create variable that will tell complier where it will end reading
	filename, err := filenameReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Trim the newline from the filename
	filename = strings.TrimSpace(filename)

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// using difer to close the file is good way of close  
	defer file.Close()
    
	// Write content to the file
	bytesWritten, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("The length of your string is:", bytesWritten)
	Readfilecontent(filename)
}


// Create the function that will be able to read the content of file inside the file

func Readfilecontent(filename string){
     databyte,err := os.ReadFile(filename)

	 if err != nil{
		panic(err)
	 }

	 fmt.Println("the content within file is \n: ",string(databyte))
	
}