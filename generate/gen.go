// This program generates version.go.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

var tpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// THIS FILE IS AUTOGENERATED BY THE MAKEFILE. DO NOT EDIT.
package version

import "fmt"

var (
	version = "{{ .Version }}"
	buildDate = "{{ .BuildDate }}"
	commitHash = "{{ .CommitHash }}"
	author = "{{ .Author }}"
)

// PrintVersion returns the current version information
func PrintVersion() {
	fmt.Printf("Version %s, Date: %s (Commit: %s), Author: %s\n", version, buildDate, commitHash, author)
}
`))

func main() {

	version, err := ioutil.ReadFile("../VERSION.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	cmd, err := exec.Run("git rev-parse --verify HEAD", nil, nil, exec.DevNull, exec.Pipe, exec.MergeWithStdout)

	cmdOutput, _ := ioutil.ReadAll(cmd.Stdout)

	f, err := os.Create("version/version_new.go")
	die(err)
	defer f.Close()

	tpl.Execute(f, struct {
		Version    string
		BuildDate  string
		CommitHash string
		Author     string
	}{
		Version:    string(version),
		BuildDate:  time.Now().Format(time.RFC850),
		CommitHash: string(cmdOutput),
		Author:     "Alain Lefebvre",
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}