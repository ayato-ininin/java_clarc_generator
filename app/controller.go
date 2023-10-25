package app

import (
	"fmt"
	"io/ioutil"
)

func GenerateControllerFile(defaultPath string) {
	// ディレクトリを生成
	dir := defaultPath + "/controller/" + domain
	// controllerファイルを生成
	generateController(dir)
}

func generateController(dir string) {
	// ファイル名を生成（domainNameの最初を大文字へ変換）
	fileName := dn + "Controller.java"
	// ファイルの中身を生成
	content := generateControllerContent()
	filePath := dir + "/" + fileName
	// ディレクトリが存在するか確認し、存在しない場合は作成
	if err := EnsureDir(dir); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Generated controller file:", filePath)
}

func generateControllerContent() string {
	javaContent :=
		fmt.Sprintf(`package com.example.todoapi.controller.%s;`, domain) + "\n\n" +
		fmt.Sprintf("public class %sController {", dn) + "\n" +
		"}"
	return javaContent
}
