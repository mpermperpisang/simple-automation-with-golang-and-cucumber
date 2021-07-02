package main

import (
	"github.com/cucumber/godog"
	"github.com/simple-automation-with-golang-and-cucumber/features/suites"
	"github.com/simple-automation-with-golang-and-cucumber/features/supports"
)

func GodogTestSuiteContext(s *godog.TestSuiteContext) {
	supports.InitializeTestSuite(s)
}

func GodogScenarioContext(s *godog.ScenarioContext) {
	supports.InitializeScenario(s)
	suites.AutomationScenarioStep(s)
}
