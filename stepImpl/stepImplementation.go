package stepImpl

import (
	"errors"
	"fmt"
	"github.com/getgauge-contrib/gauge-go/gauge"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/hakanyolat/gauge-go-todo-api-tests/condition"
	"github.com/tebeka/selenium"
	"strings"
	"time"
)

func findTodoElementInList(todo string) (selenium.WebElement, error) {
	var err error
	var els []selenium.WebElement

	if err = driver.Wait(condition.ElementsIsLocated(selenium.ByCSSSelector, ".todo-title")); err != nil {
		return nil, err
	}

	if els, err = driver.FindElements(selenium.ByCSSSelector, ".todo-title"); err != nil {
		return nil, err
	} else {
		for _, el := range els {
			if text, _ := el.Text(); text == todo {
				return el, nil
			}
		}
		return nil, errors.New("todo not found in list")
	}

}

var _ = gauge.Step("Clear all todos", func() {
	_ = driver.SetImplicitWaitTimeout(time.Second)

	els, _ := driver.FindElements(selenium.ByCSSSelector, ".remove-item")
	for _, el := range els {
		_ = el.Click()
	}
})

var _ = gauge.Step("Add todo <todo>", func(todo string) {
	el, _ := driver.FindElement(selenium.ByID, "todo-input")
	_ = el.SendKeys(todo + selenium.EnterKey)
})

var _ = gauge.Step("See <todo> in list", func(todo string) {
	//_ = driver.SetImplicitWaitTimeout(time.Second)
	if _, err := findTodoElementInList(todo); err != nil {
		T.Fail(err)
	}
})

var _ = gauge.Step("Click <todo>", func(todo string) {
	if el, err := findTodoElementInList(todo); err != nil {
		T.Fail(err)
	} else {
		_ = el.Click()
	}
})

var _ = gauge.Step("Remove todo <todo>", func(todo string) {
	var err error
	var todoEl selenium.WebElement
	var parentEl selenium.WebElement
	var removeEl selenium.WebElement

	if todoEl, err = findTodoElementInList(todo); err != nil {
		T.Fail(err)
	}

	if parentEl, err = todoEl.FindElement(selenium.ByXPATH, "./.."); err != nil {
		T.Fail(err)
	}

	if removeEl, err = parentEl.FindElement(selenium.ByCSSSelector, ".remove-item"); err != nil {
		T.Fail(err)
	}

	_ = removeEl.Click()
})

var _ = gauge.Step("Check <todo> is completed", func(todo string) {
	var el selenium.WebElement
	var err error
	var classAttr string

	if el, err = findTodoElementInList(todo); err != nil {
		T.Fail(err)
	}

	if classAttr, err = el.GetAttribute("class"); err != nil {
		T.Fail(err)
	}

	if !strings.Contains(classAttr, "todo-completed") {
		T.Fail(fmt.Errorf("%s is not completed", todo))
	}
})

var _ = gauge.Step("Check <todo> is not completed", func(todo string) {
	var el selenium.WebElement
	var err error
	var classAttr string

	if el, err = findTodoElementInList(todo); err != nil {
		T.Fail(err)
	}

	if classAttr, err = el.GetAttribute("class"); err != nil {
		T.Fail(err)
	}

	if strings.Contains(classAttr, "todo-completed") {
		T.Fail(fmt.Errorf("%s is completed", todo))
	}
})

var _ = gauge.Step("Check todo list is empty", func() {
	_ = driver.SetImplicitWaitTimeout(time.Second)

	if els, err := driver.FindElements(selenium.ByCSSSelector, ".todo-title"); err != nil {
		T.Fail(err)
	} else if len(els) > 0 {
		T.Fail(fmt.Errorf("todo list is not empty: %d item found", len(els)))
	}
})
