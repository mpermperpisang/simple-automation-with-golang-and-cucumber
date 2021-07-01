package supports

import (
	"fmt"

	"github.com/cucumber/godog"
)

var Godogs int

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		Godogs = 0
	})
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.AfterScenario(func(scenario *godog.Scenario, log error) {
		if log != nil {
			fmt.Println(log)
		}
	})
}
