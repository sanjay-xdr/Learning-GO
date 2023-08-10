package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

func createNotes(authorName string) {

	reader := bufio.NewReader(os.Stdin)
	content := bufio.NewReader(os.Stdin)
	fmt.Println("Enter File Name")
	fileNmae, _ := reader.ReadString('\n')
	newName := strings.TrimSpace(fileNmae)
	name, _ := os.Getwd()

	file, err := os.Create(name + "/" + newName + ".txt")
	if err != nil {
		fmt.Println(err, "I am err")
	}
	dt := time.Now()

	io.WriteString(file, "Author : "+authorName+"\n")
	io.WriteString(file, "Created at: "+dt.Format("02-01-2006 15:04:05")+"\n")
	fmt.Println("Enter the Content of the file")
	userContent, _ := content.ReadString('\n')

	_, fileWriteErr := io.WriteString(file, userContent)
	if err != nil {
		fmt.Println("Something went wrong", fileWriteErr)

	}
	fmt.Println(color.InGreen("File Created Successfully"))

	defer file.Close()

}

func main() {

	fmt.Println(color.OverBlue("Welcome GoNotes"))

	authorName := ReadFile("./config.txt")

	for true {

		fmt.Println(color.InPurple("Write newfile to create new file"))
		fmt.Println(color.InRed("Write quit  to stop this program"))

		userValRead := bufio.NewReader(os.Stdin)
		userVa1l, _ := userValRead.ReadString('\n')

		userVal := strings.TrimSpace(userVa1l)
		switch userVal {
		case "newfile":
			createNotes(authorName)

		case "quit":
			os.Exit(1)
		default:
			fmt.Println("Please Enter a Correct Choice")
		}

	}

}

func ReadFile(filename string) string {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(color.InRed("Something Went Wrong, Make sure you have correct config.txt"))
	}

	jsonData := string(databyte)

	authorName := strings.Split(jsonData, ":")
	return authorName[1]

}

type Author struct {
	Name string `json:"author"`
}
