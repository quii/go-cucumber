Feature: calculator
  In order to do math
  As a programmer
  I want to be able to use a calculator

  Scenario: Add two numbers
    Given I have entered 50 into the calculator
    And I have entered 70 into the calculator
    When I press add
    Then the result should be 120

  Scenario: Multiply two numbers
    Given I have entered 2 into the calculator
    And I have entered 3 into the calculator
    When I press multiply
    Then the result should be 6
