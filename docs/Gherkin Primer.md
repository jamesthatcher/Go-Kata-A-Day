# Gherkin Primer

The natural order of a scenario is _Given..._ _When..._ _Then..._:

* _Given_ describes the preconditions for the scenario and prepares the environment
* _When_ describes the action under test
* _Then_ describes the expected outcomes

The _And_ and _But_ keywords can be used to join several _Given_, _When_, or _Then_ steps together in a more readable way:

```
Given I have a current account with $1000
And I have a savings account with $2000
```

Several related scenarios can often be grounds into a single scenario using a table of examples. For example, the following scenario illustrates how interest is calculated on different types of accounts:

```
Scenario Outline: Earning interest
    Given I have an account of type <account-type> with a balance of <initial-balance>
    When the monthly interest is calculated
    Then I should have earned an annual interest rate of <interest-rate>
    And I should have a new balance of <new-balance>
  
    Examples:
      | initial-balance | account-type | interest-rate | new balance |
      | 10000           | "current"    | 1             | 10008.33    |
      | 10000           | "savings"    | 3             | 10025       |
      | 10000           | "supersaver" | 5             | 10041.67    |

```

Specification by example explained, ShriKant Vashishtha

[Source](https://www.scrumday.in/wp-content/uploads/2017/09/specification-by-examples-writing-executable-specification.pdf)