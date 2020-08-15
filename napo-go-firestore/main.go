package main

import (
	"app/app"
	"app/app/v1/injection"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Project Location " + dir)

	args := os.Args[1:]
	if len(args) == 1 {
		if args[0] == "-h" {
			cmdH()
		} else if args[0] == "-start" {
			cmdStart("local")
		} else {
			printErr(args)
		}
	} else if len(args) == 2 {
		if args[0] == "-start" {
			if args[1] == "local" || args[1] == "dev" || args[1] == "prod" {
				cmdStart(args[1])
			} else {
				printErr(args)
			}
		} else {
			printErr(args)
		}
	} else {
		printErr(nil)
	}
}

func printErr(args []string) {
	if len(args) == 1 {
		println("Invalid Arguments [" + args[0] + "] !. " +
			"\nCheck help with -h")
		return
	} else if len(args) == 2 {
		println("Invalid Arguments [" + args[0] + " + " + args[1] + "]!. " +
			"\nCheck help with -h")
		return
	}

	println("Arguments required!. \nCheck help with -h")
}

func cmdH() {
	println("--------------------------------------")
	println("|               Tools                |")
	println("--------------------------------------")
	println("|  -h                 Help           |")
	println("|  -start             Run Local Env  |")
	println("|  -start <env>       Run Env        |")
	println("--------------------------------------")
	println()
	println("--------------------------------------")
	println("|            Environment             |")
	println("--------------------------------------")
	println("|   local             Local          |")
	println("|   dev               Development    |")
	println("|   prod              Production     |")
	println("--------------------------------------")
}

func cmdStart(evn string) {
	//Sertup COnfig
	app.SetupConfig(evn)

	db := app.SetupDbConnection()
	defer db.Close()
	firestoreClient := app.SetupFirestore()
	defer firestoreClient.Close()

	//Inject Dependency for All Services
	services := injection.NewInstanceServiceInjection(db, firestoreClient)

	//setup router
	app.SetupRoute(services)

}
