package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {

	url := launcher.New().
		// Proxy("127.0.0.1:8080").     // set flag "--proxy-server=127.0.0.1:8080"
		// Delete("use-mock-keychain"). // delete flag "--use-mock-keychain"
		MustLaunch()

	// Launch a new browser with default options, and connect to it.
	browser := rod.New().ControlURL(url).MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	fmt.Println("a")
	// Create a new page
	page := browser.MustPage("https://www.baidu.com")

	fmt.Println("b")
	// We use css selector to get the search input element and input "git"
	page.MustElement("input").MustInput("git").MustType(input.Enter)

	fmt.Println("c")
	// Wait until css selector get the element then get the text content of it.
	text := page.MustElement(".codesearch-results p").MustText()

	fmt.Println(text)

	// Get all input elements. Rod supports query elements by css selector, xpath, and regex.
	// For more detailed usage, check the query_test.go file.
	fmt.Println("Found", len(page.MustElements("input")), "input elements")

	// Eval js on the page
	page.MustEval(`() => console.log("hello world")`)

	// Pass parameters as json objects to the js function. This MustEval will result 3
	fmt.Println("1 + 2 =", page.MustEval(`(a, b) => a + b`, 1, 2).Int())

	// When eval on an element, "this" in the js is the current DOM element.
	fmt.Println(page.MustElement("title").MustEval(`() => this.innerText`).String())

}
