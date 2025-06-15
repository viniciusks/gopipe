package main

import (
	"github.com/viniciusks/gopipe/cmd"
)

var (
    tag = "dev"
    commit  = "none"
    date    = "unknown"
)

func main() {
	cmd.SetVersionVars(tag, commit, date)
    cmd.Execute()
}