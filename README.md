# Todo API Acceptance Tests

This repo includes acceptance tests for todo application.

### Requirements

* Backend Project [Todo API with Golang](https://github.com/hakanyolat/go-todo-api)
* Frontend Project [Vue client for Todo API](https://github.com/hakanyolat/todo-api-vue-client) (Test URL)
* [Gauge](https://gauge.org/)
* Gauge go plugin (can be installed using ```gauge --install go```)
* [Chrome driver](http://chromedriver.storage.googleapis.com/index.html) (On mac: ```brew install chromedriver```)
* Selenium server standalone (Available in resources directory of this repository. If you want a different version of it, it can be downloaded from [here](http://selenium-release.storage.googleapis.com/index.html))

> Note: Make sure the browser (Google Chrome) installed on your computer is the version compatible with the webdriver.

### Installation

1. Download this project with ```go get```.

```bash
$ go get github.com/hakanyolat/gauge-go-todo-api-tests
```

2. Install the dependencies with ```go mod```.

```bash
$ go mod vendor
```

### Configuration

Configurations are defined in the ```.env``` file.

```
GAUGE_TEST_URL=http://127.0.0.1:8080
GAUGE_SELENIUM_HUB_URL=http://192.168.56.1:4444/wd/hub
```

Alternatively, you can set the environment value as in the example below.

```bash
$ export GAUGE_TEST_URL=http://127.0.0.1:8080
$ export GAUGE_SELENIUM_HUB_URL=http://192.168.56.1:4444/wd/hub
```

> Note: Make sure [GOPATH](https://github.com/golang/go/wiki/GOPATH) is present in your Environment values.

### Running Tests

1. Launch selenium grid hub

```bash
$ java -jar <path_to_selenium_server_jar> -role hub
```

2. Launch a selenium grid node

```bash
$ java -jar <path_to_selenium_server_jar> -role node
```

3. Execute tests (This uses Chrome as default browser for specs execution)

```bash
$ gauge run specs
```