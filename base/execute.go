package base

import (
	"os/exec"
	"log"
	"fmt"
	"strings"
	"bytes"
)

func lookPath()  {
	path, err := exec.LookPath("fortune")

	if err != nil{
		log.Fatal("installing fortune is in your future")
	}else{
		fmt.Printf("fortune is available at %s\n", path)
	}
}

func command()  {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %q\n", out.String())
}


