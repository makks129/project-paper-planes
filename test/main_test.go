package main

import (
	"os"
	"testing"

	"github.com/makks129/project-paper-planes/src/utils"
	"github.com/makks129/project-paper-planes/test/mock"
)

func TestMain(m *testing.M) {
	defer teardown()
	setup()
	res := m.Run()
	os.Exit(res)
}

func setup() {
	utils.Log("<<< SETUP >>>")
	mock.StartMysqlContainer()
	mock.WaitForDB()
	mock.CreateDB()
}

func teardown() {
	utils.Log("<<< TEARDOWN >>>")
	mock.DropDB()
	mock.StopMysqlContainer()
}
