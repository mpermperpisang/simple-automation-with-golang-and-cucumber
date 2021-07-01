@example
Feature: Example Automation
  Simple automation with Golang and Cucumber

  @simple-godog-1
  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

  @simple-godog-2
  Scenario Outline: Eat godogs with different quantity
    Given there are <existQty> godogs
    When I eat <eatenQty>
    Then there should be <remainQty> remaining

    Examples:
      | existQty | eatenQty | remainQty |
      | 10       | 5        | 5         |
      | 5        | 5        | 0         |

  @simple-godog-3
  Scenario Outline: Eat godogs with different quantity taken from env
    Given there are <existQty> godogs
    When I eat <eatenQty>
    Then there should be <remainQty> remaining

    Examples:
      | existQty      | eatenQty | remainQty |
      | ENV:GODOGS_10 | 5        | 5         |
      | ENV:GODOGS_5  | 5        | 0         |

  @simple-godog-4 @exclude
  Scenario: Eat godogs with different quantity with data table
    Given there are 12 godogs
    When I eat godogs
      | qty |
      | 2   |
      | 3   |
    Then there should be 7 remaining
