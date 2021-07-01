package stepdefinitions

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"github.com/simple-automation-with-golang-and-cucumber/features/helper"
	"github.com/simple-automation-with-golang-and-cucumber/features/supports"
)

func ThereAreGodogs(available string) error {
	var err error

	if strings.Contains(available, "ENV:") {
		qty := strings.ReplaceAll(available, "ENV:", "")
		available = os.Getenv(qty)
	}

	supports.Godogs, err = strconv.Atoi(available)
	helper.LogPanicln(err)

	return nil
}

func IEat(num int) error {
	if supports.Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, supports.Godogs)
	}

	supports.Godogs -= num

	return nil
}

func IEatALot(num *godog.Table) error {
	for index := 1; index < len(num.Rows); index++ {
		qty, _ := strconv.Atoi(num.Rows[index].Cells[0].Value)

		IEat(qty)
	}

	return nil
}

func ThereShouldBeRemaining(remaining int) error {
	if supports.Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, supports.Godogs)
	}

	return nil
}
