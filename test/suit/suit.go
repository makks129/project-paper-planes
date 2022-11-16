package suit

import "testing"

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

func (s *SubTests) TestIt(name string, f func(t *testing.T)) {
	if s.AfterEach != nil {
		defer s.AfterEach()
	}
	if s.BeforeEach != nil {
		s.BeforeEach()
	}
	s.T.Run(name, f)
}
