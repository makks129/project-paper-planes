package main

import (
	"testing"

	"github.com/makks129/project-paper-planes/test/mock"
)

// go test -run Test -v ./test/...

func TestMain(m *testing.M) {
	defer teardown()
	setup()

	m.Run()
}

func setup() {
	mock.StartMysqlContainer()
	mock.WaitForDB()
	mock.CreateDB()
}

func teardown() {
	mock.DropDB()
	mock.StopMysqlContainer()
}
