package main

import (
	"testing"

	"github.com/makks129/project-paper-planes/test/mock"
)

// go test -run Test -v ./test

func TestMain(m *testing.M) {
	defer setupAndTeardown()()
	m.Run()
}

func setupAndTeardown() func() {
	// Setup
	mock.StartMysqlContainer()

	// Teardown
	return func() {
		mock.StopMysqlContainer()
	}
}
