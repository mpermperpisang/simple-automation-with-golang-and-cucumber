@example
Feature: Example Automation
  Simple automation with Golang and Cucumber

  Background: Available Godogs
    Given there are 12 godogs

  @simple-godog-1
  Scenario: Eat 5 out of 12
    When I eat 5
    Then there should be 7 remaining

  @simple-godog-2
  Scenario: Eat godogs with different quantity with data table
    When I eat godogs
      | qty |
      | 2   |
      | 3   |
    Then there should be 7 remaining
