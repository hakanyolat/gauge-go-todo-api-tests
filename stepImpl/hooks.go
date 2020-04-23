package stepImpl

import (
	"fmt"
	"github.com/getgauge-contrib/gauge-go/gauge"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/hakanyolat/gauge-go-todo-api-tests/config"
	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver
var conf config.Config

var _ = gauge.BeforeSuite(func() {
	conf = config.GetConfig()

	var err error

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, conf.SeleniumHubUrl)

	if err != nil {
		T.Fail(fmt.Errorf("Failed to open session: %s\n", err.Error()))
		return
	}

	err = driver.Get(conf.TestUrl)

	if err != nil {
		T.Fail(fmt.Errorf("Failed to load page: %s\n", err.Error()))
		return
	}
}, []string{}, AND)

var _ = gauge.AfterSuite(func() {
	panic(driver.Quit())
}, []string{}, AND)
