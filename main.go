package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"oj/apps"
	"os"
)

func main() {

	var appFlag string
	flag.StringVar(&appFlag, "i", "", "cli app")
	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("You need to pipe information throgh oj")
		os.Exit(1)
	}

	var app = NewApp(apps.NewEnv())

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	var jsonOutput = app.parse(string(output))
	fmt.Println(jsonOutput)
}
