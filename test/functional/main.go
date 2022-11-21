package main

import (
	"flag"

	"github.com/makks129/project-paper-planes/src/utils"
	"github.com/makks129/project-paper-planes/test/functional/mock"
)

func main() {
	wordPtr := flag.String("phase", "", "values: setup, teardown")
	flag.Parse()
	switch *wordPtr {
	case "setup":
		setup()
	case "teardown":
		teardown()
	default:
		panic("Unknown phase")
	}
}

func setup() {
	utils.Log("=============== SETUP ===============")
	mock.StartMysqlContainer()
	mock.WaitForDB()
	mock.CreateDB()
	utils.Log("=====================================")
}

func teardown() {
	utils.Log("============== TEARDOWN ==============")
	mock.DropDB()
	mock.StopMysqlContainer()
	utils.Log("======================================")
}
