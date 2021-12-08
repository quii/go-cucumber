package go_cucumber

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	calculator := Calculator{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		calculator = Calculator{}
		return ctx, nil
	})

	ctx.Step(`^I have entered (\d+) into the calculator$`, calculator.EnterNumber)
	ctx.Step(`^I press add$`, calculator.Add)
	ctx.Step(`^I press multiply$`, calculator.Multiply)
	ctx.Step(`^the result should be (\d+)$`, func(expected int) error {
		if calculator.Result() != expected {
			return fmt.Errorf("Expected %d, got %d", expected, calculator.Result())
		}
		return nil
	})
}
