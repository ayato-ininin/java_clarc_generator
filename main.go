package main

import (
	"fmt"
	"jaca_clarc_generator/app"
)

func main() {
	var defaultPath string
	var domainName string
	var methodName string
	fmt.Println("Enter default path (e.g., /Users/xxx/java/todo-api/src/main/java/com/example/todoapi):")
	fmt.Scanln(&defaultPath)
	fmt.Println("Enter domain name (e.g., task):")
	fmt.Scanln(&domainName)
	fmt.Println("Enter method name (e.g., create):")
	fmt.Scanln(&methodName)
	app.SetCommonVariable(domainName, methodName)

	// Generate controllerファイル
	app.GenerateControllerFile(defaultPath)
	// Generate domain serviceファイル
	app.GenerateDomainServiceFile(defaultPath)
	// Generate usecaseファイル
	app.GenerateUsecaseFile(defaultPath)
	// Generate stubファイル
	app.GenerateStubFile(defaultPath)
}
