package main

import (
	"flag"
	"log"

	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"

	"github.com/casbin/casbin"
)

var (
	file  = flag.Bool("f", false, "使用文件策略存储.")
	mysql = flag.Bool("mysql", false, "Use a MySQL storage layer.")
)

func main() {
	flag.Parse()

	var a interface{}
	if *mysql {
		// Grab the address, user, and password to the mysql storage
		// from an environmental variable.
		//		fu := os.Getenv("mysqlUsr")
		//		fp := os.Getenv("mysqlPass")
		//		fAddr := os.Getenv("mysqlAddr")

		a = gormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/test", true) // Your driver and data source.
		log.Println("use mysql policy storage!")
	} else if *file {
		// 就两种 非此即彼  如果是多种那么用switch 或者map 或者 if elseif elseif ... 结构来决策
		a = "./config/policy.csv"
		log.Println("use file policy storage! ")
	} else {
		// 默认
		a = "./config/policy.csv"
		log.Println("use default file policy storage! ")
	}

	e := casbin.NewEnforcer("./config/model.conf" /*"./config/policy.csv"*/, a)

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
