package main

import (
	"log"

	"github.com/casbin/casbin"
)

func main() {
	e := casbin.NewEnforcer("./config/model.conf", "./config/policy.csv")

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	if e.Enforce(sub, obj, act) == true {
		// permit alice to read data1
		log.Println("permit alice to read data1")
	} else {
		// deny the request, show an error

		log.Fatalln("  alice has no rights to read data1")
	}
	//    /*
	//	roles := e.GetAllRoles( /*"alice"*/ )
	//	log.Println(roles)
}
