Feature: adding numbers
  In order to do math
  As a programmer
  I want to be able to add numbers

  Scenario: Add two numbers
    Given I have entered 50 into the calculator
    And I have entered 70 into the calculator
    When I press add
    Then the result should be 120