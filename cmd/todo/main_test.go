package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Runnig test....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}
