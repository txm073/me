package main

import (
	"fmt"

	me "example.com/me/internal/backend"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := me.ProjectDB.Load()
	check(err)
	fmt.Printf("Loaded %d project(s)\n", len(me.ProjectDB.Projects))
	me.CreateHandlers()
	me.RunServer()
	check(me.ProjectDB.Save())
}
