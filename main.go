package main

import "github.com/minj131/jobcoin/cmd"

func main() {
	err := cmd.Jobcoin()
	if err != nil {
		panic(err)
	}
}
