package main

import "github.com/MrChaos1993/gen"

type MainMethod interface {
	// FindLocal
	//
	// select * from @@table where id=@id
	FindLocal(id int) gen.T
}

func main() {}
