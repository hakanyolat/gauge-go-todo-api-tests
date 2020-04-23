package condition

import "github.com/tebeka/selenium"

func ElementIsLocated(by, selector string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(by, selector)
		return err == nil, nil
	}
}

func ElementsIsLocated(by, selector string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElements(by, selector)
		return err == nil, nil
	}
}
