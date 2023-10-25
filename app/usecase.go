package app

import (
	"fmt"
	"io/ioutil"
)

func GenerateUsecaseFile(defaultPath string) {
	// ディレクトリを生成
	dir := defaultPath + "/usecase/" + domain + "/" + method
	// usecase IFファイルを生成
	generateUsecase(dir)
	// usecase inputDataファイルを生成
	generateInputData(dir)
	// usecase outputDataファイルを生成
	generateOutputData(dir)
}

func generateUsecase(dir string) {
	// ファイル名を生成（domainNameの最初を大文字へ変換）
	fileName := dnmn + "UseCase.java"
	// ファイルの中身を生成
	content := generateUsecaseIFContent()
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
	fmt.Println("Generated usecase IF file:", filePath)
}

func generateUsecaseIFContent() string {
	javaContent :=
		fmt.Sprintf(`package com.example.todoapi.usecase.%s.%s;`, domain, method) + "\n\n" +
		"import com.example.todoapi.usecase.core.UseCase;" + "\n\n" +
		fmt.Sprintf("public interface %sUseCase extends UseCase<%s,%s> {",
			dnmn, dnmn + "InputData", dnmn + "OutputData") + "\n" +
		fmt.Sprintf("    %sOutputData handle(%sInputData inputData);", dnmn, dnmn) + "\n" +
		"}"
	return javaContent
}

func generateInputData(dir string) {
	// ファイル名を生成（domainの最初を大文字へ変換）
	fileName := dnmn + "InputData.java"
	// ファイルの中身を生成
	content := generateInputDataContent()
	filePath := dir + "/" + fileName
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Generated usecase inputData file:", filePath)
}

func generateInputDataContent() string {
	javaContent :=
		fmt.Sprintf(`package com.example.todoapi.usecase.%s.%s;`, domain, method) + "\n\n" +
		"import com.example.todoapi.usecase.core.InputData;" + "\n\n" +
		fmt.Sprintf("public class %sInputData implements InputData<%sOutputData> {",
			dnmn, dnmn) + "\n" + "}"
	return javaContent
}

func generateOutputData(dir string) {
	// ファイル名を生成（domainの最初を大文字へ変換）
	fileName := dnmn + "OutputData.java"
	// ファイルの中身を生成
	content := generateOutputDataContent()
	filePath := dir + "/" + fileName
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Generated usecase outputData file:", filePath)
}

func generateOutputDataContent() string {
	javaContent :=
		fmt.Sprintf(`package com.example.todoapi.usecase.%s.%s;`, domain, method) + "\n\n" +
		"import com.example.todoapi.usecase.core.OutputData;" + "\n\n" +
		fmt.Sprintf("public class %sOutputData implements OutputData {",
			dnmn) + "\n" + "}"
	return javaContent
}

