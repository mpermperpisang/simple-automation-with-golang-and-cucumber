package suites

import (
	"github.com/cucumber/godog"
	"github.com/simple-automation-with-golang-and-cucumber/features/stepdefinitions"
)

// AutomationScenarioStep : steps collection
func AutomationScenarioStep(s *godog.ScenarioContext) {
	s.Step(`^there are (.*) godogs$`, stepdefinitions.ThereAreGodogs)
	s.Step(`^I eat (\d+)$`, stepdefinitions.IEat)
	s.Step(`^I eat godogs$`, stepdefinitions.IEatALot)
	s.Step(`^there should be (-?\d+) remaining$`, stepdefinitions.ThereShouldBeRemaining)
}
