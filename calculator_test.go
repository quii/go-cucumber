package go_cucumber

import (
	"fmt"
	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	calculator := Calculator{}

	ctx.Step(`^I have entered (\d+) into the calculator$`, calculator.EnterNumber)
	ctx.Step(`^I press add$`, calculator.Add)
	ctx.Step(`^the result should be (\d+)$`, func(expected int) error {
		if calculator.Result() != expected {
			return fmt.Errorf("Expected %d, got %d", expected, calculator.Result())
		}
		return nil
	})
}
