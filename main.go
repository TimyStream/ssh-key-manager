package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cheynewallace/tabby"
)

func main() {
	fmt.Println("====================================")
	fmt.Println("|| Welcome to the SSH Key Manager ||")
	fmt.Println("====================================\n")

	printFileAsTable("~/.ssh/authorized_keys")
}

func printFileAsTable(file string) {
	databyte, err := ioutil.ReadFile(file)
	checkNilErorr(err)
	content := string(databyte)

	contentArray := strings.Fields(content)

	t := tabby.New()
	t.AddHeader("ID", "SSH Key Type", "SSH KEY", "USER")
	counter := 1
	for i := 0; i < len(contentArray)-1; i++ {
		t.AddLine(counter, contentArray[i], truncateText(contentArray[i+1], 68), contentArray[i+2])
		i++
		i++
		counter++
	}
	t.Print()
}

func readFile(file string) {
	databyte, err := ioutil.ReadFile(file)
	checkNilErorr(err)

	fmt.Println("Output of file: \n", string(databyte))
}

func appendFile(file string) {
	databyte, err := ioutil.ReadFile(file)
	checkNilErorr(err)
	appendedFile, err := os.Create(file)
	checkNilErorr(err)

	content := string(databyte) + "\n" + "second key"

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
