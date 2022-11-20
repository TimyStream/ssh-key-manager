package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cheynewallace/tabby"
)

func main() {
	removeLine("./test/authorized_keys")

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

	user := readUserInput("User")
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

func removeSlice(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func readUserInput(userInputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userInputText + ": ")
	answer, err := reader.ReadString('\n')
	checkNilErorr(err)
	return answer
}

func checkNilErorr(err error) {
	if err != nil {
		panic(err)
	}
}
