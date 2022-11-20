package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/cheynewallace/tabby"
)

// The Main Function for the Program
func main() {
	fmt.Println("====================================")
	fmt.Println("|| Welcome to the SSH Key Manager ||")
	fmt.Println("====================================")
	printEmtyLine()

	menu("./test/authorized_keys")
}

// The Main Menu function to ask the user what he wants to do
func menu(file string) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Table list of all SSH Keys")
	fmt.Println("2. Add SSH Key")
	fmt.Println("3. Remove SSH Key")
	fmt.Println("4. Exit")
	printEmtyLine()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose [1,2,3,4]: ")
	input, err := reader.ReadString('\n')
	checkNilErorr(err)
	inputNumber, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	checkNilErorr(err)
	switch inputNumber {
	case 1:
		printEmtyLine()
		printFileAsTable(file)
		printEmtyLine()
		printEmtyLine()
		menu(file)
	case 2:
		printEmtyLine()
		appendFile(file)
		printEmtyLine()
		menu(file)
	case 3:
		printEmtyLine()
		removeLine(file)
		printEmtyLine()
		menu(file)
	case 4:
		printEmtyLine()
		fmt.Println("\n Thank you for using the SSH Key Manager. Bye :)")
		os.Exit(0)
	default:
		fmt.Println("Wrong Input")
		printEmtyLine()
		menu(file)
	}
}

func printEmtyLine() {
	fmt.Println()
}

func printFileAsTable(file string) {
	databyte, err := ioutil.ReadFile(file)
	checkNilErorr(err)
	content := string(databyte)

	contentArray := strings.Fields(content)

	t := tabby.New()
	t.AddHeader("ID", "SSH Key Type", "SSH KEY (Truncate to 68 Chars)", "USER")
	counter := 1
	for i := 0; i < len(contentArray)-1; i++ {
		t.AddLine(counter, contentArray[i], truncateText(contentArray[i+1], 68), contentArray[i+2])
		i++
		i++
		counter++
	}
	t.Print()
}

func removeLine(file string) {
	input, err := ioutil.ReadFile(file)
	checkNilErorr(err)
	lines := strings.Split(string(input), "\n")

	content := string(input)
	contentArray := strings.Fields(content)
	counter := 1
	t := tabby.New()
	t.AddHeader("ID", "USER")
	for i := 0; i < len(contentArray)-1; i++ {
		t.AddLine(counter, contentArray[i+2])
		i++
		i++
		counter++
	}
	t.Print()
	printEmtyLine()

	user := readUserInput("User from Table")
	user = strings.TrimSpace(user)
	checkNilErorr(err)

	for i, line := range lines {
		if strings.Contains(line, user) {
			lines[i] = lines[len(lines)-1] // Copy last element to index i.
			lines[len(lines)-1] = ""       // Erase last element (write zero value).
			lines = lines[:len(lines)-1]   // Truncate slice.
			fmt.Println("Test")
		}
	}
	output := strings.Join(lines, "\n")
	outputFile, err := os.Create(file)
	checkNilErorr(err)
	length, err := io.WriteString(outputFile, output)
	checkNilErorr(err)

	fmt.Println("Lenth is: ", length)
	defer outputFile.Close()
}

func readUserInput(userInputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userInputText + ": ")
	answer, err := reader.ReadString('\n')
	checkNilErorr(err)
	return answer
}

func appendFile(file string) {
	databyte, err := ioutil.ReadFile(file)
	checkNilErorr(err)

	content := string(databyte) + "\n" + readUserInput("SSH Key with Comment (Will be the User)")

	appendedFile, err := os.Create(file)
	checkNilErorr(err)
	length, err := io.WriteString(appendedFile, content)
	checkNilErorr(err)

	fmt.Println("length is: ", length)
	defer appendedFile.Close()
}

func truncateText(s string, max int) string {
	return s[:max]
}

func checkNilErorr(err error) {
	if err != nil {
		panic(err)
	}
}
