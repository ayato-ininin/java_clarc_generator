package app

import (
	"fmt"
	"io/ioutil"
)

func GenerateStubFile(defaultPath string) {
	// ディレクトリを生成
	stubDir := defaultPath + "/stubs/" + domain
	stubJsonDir := defaultPath + "/stubs/jsons"
	// stubファイルを生成
	generateStub(stubDir)
	// stub jsonファイルを生成（中身空）
	generateStubJson(stubJsonDir)
}

func generateStub(dir string) {
	// ファイル名を生成（domainNameの最初を大文字へ変換）
	fileName := "Stub" + dnmn + "Interactor.java"
	// ファイルの中身を生成
	content := generateStubContent()
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
	fmt.Println("Generated stub interactor file:", filePath)
}

func generateStubContent() string {
	javaContent :=
		fmt.Sprintf(`package com.example.todoapi.stubs.%s;`, domain) + "\n\n" +
		"import com.example.todoapi.objectLoader.JsonsLoader;" + "\n" +
		fmt.Sprintf("import com.example.todoapi.usecase.%s.%s.%sInputData;", domain, method, dnmn) + "\n" +
		fmt.Sprintf("import com.example.todoapi.usecase.%s.%s.%sOutputData;", domain, method, dnmn) + "\n" +
		fmt.Sprintf("import com.example.todoapi.usecase.%s.%s.%sUseCase;", domain, method, dnmn) + "\n" +
		"import org.springframework.context.annotation.Profile;" + "\n" +
		"import org.springframework.stereotype.Service;" + "\n\n" +
		"@Service" + "\n" +
		"@Profile(\"dev\")" + "\n" +
		"public class Stub" + dnmn + "Interactor implements " + dnmn + "UseCase {" + "\n" +
		"	private JsonsLoader jsonsLoader;" + "\n" +
		"	public Stub" + dnmn + "Interactor(JsonsLoader jsonsLoader) {" + "\n" +
		"		this.jsonsLoader = jsonsLoader;" + "\n" +
		"	}" + "\n\n" +
		"	@Override" + "\n" +
		"	public " + dnmn + "OutputData handle(" + dnmn + "InputData inputData) {" + "\n" +
		"		return jsonsLoader.generate(" + dnmn + "OutputData.class);" + "\n" +
		"	}" + "\n" +
		"}"
	return javaContent
}

func generateStubJson(dir string) {
	// ファイル名を生成（domainNameの最初を大文字へ変換）
	fileName := dnmn + "OutputData.jsons"
	// ファイルの中身を生成
	filePath := dir + "/" + fileName
	// ディレクトリが存在するか確認し、存在しない場合は作成
	if err := EnsureDir(dir); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	err := ioutil.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Generated stub jsons file:", filePath)
}
