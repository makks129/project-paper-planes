package suit

import (
	"fmt"
	"strings"
	"testing"
)

func Of(subTests *SubTests) *SubTests {
	if subTests.AfterAll != nil {
		subTests.T.Cleanup(subTests.AfterAll)
	}
	return subTests
}

type SubTests struct {
	T          *testing.T
	BeforeEach func()
	AfterEach  func()
	AfterAll   func()
}

func (s *SubTests) Test(name string, f func(t *testing.T)) {
	if s.AfterEach != nil {
		defer s.AfterEach()
	}
	if s.BeforeEach != nil {
		s.BeforeEach()
	}
	s.T.Run(name, f)
}

func (s *SubTests) Skip(name string, f func(t *testing.T)) {}

func (s *SubTests) SkipLog(name string, f func(t *testing.T)) {
	fmt.Printf("    SKIP  %s/%s\n", s.T.Name(), strings.ReplaceAll(name, " ", "_"))
}
