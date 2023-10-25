package app

import (
	"fmt"
	"io/ioutil"
)

func GenerateDomainServiceFile(defaultPath string) {
	// ディレクトリを生成
	dir := defaultPath + "/domain/applicationService/" + domain
	// domainファイルを生成
	generateDomainService(dir)
}

func generateDomainService(dir string) {
	// ファイル名を生成（domainNameの最初を大文字へ変換）
	fileName := dnmn + "Interactor.java"
	// ファイルの中身を生成
	content := generateDomainServiceContent()
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
	fmt.Println("Generated domain interactor file:", filePath)
}

func generateDomainServiceContent() string {
	javaContent :=
		fmt.Sprintf(`package com.example.todoapi.domain.applicationService.%s;`, domain) + "\n\n" +
		fmt.Sprintf("import com.example.todoapi.domain.model.%s.%sRepository;", domain, dn) + "\n" +
		fmt.Sprintf("import com.example.todoapi.usecase.%s.%s.%sInputData;", domain, method, dnmn) + "\n" +
		fmt.Sprintf("import com.example.todoapi.usecase.%s.%s.%sOutputData;", domain, method, dnmn) + "\n" +
		fmt.Sprintf("import com.example.todoapi.usecase.%s.%s.%sUseCase;", domain, method, dnmn) + "\n" +
		"import org.springframework.beans.factory.annotation.Autowired;" + "\n" +
		"import org.springframework.context.annotation.Profile;" + "\n" +
		"import org.springframework.stereotype.Service;" + "\n\n" +
		"@Service" + "\n" +
		"@Profile(\"prod\")" + "\n" +
		fmt.Sprintf("public class %sInteractor implements %sUseCase {", dnmn, dnmn) + "\n\n" +
		fmt.Sprintf("	private final %sRepository %sRepository;", dn, domain) + "\n" +
		"	@Autowired" + "\n" +
		fmt.Sprintf("	public %sInteractor(%sRepository %sRepository) {", dnmn, dn, domain) + "\n" +
		fmt.Sprintf("		this.%sRepository = %sRepository;", domain, domain) + "\n" +
		"	}" + "\n" +
		"	@Override" + "\n" +
		fmt.Sprintf("	public %sOutputData handle(%sInputData inputData) {", dnmn, dnmn) + "\n" +
		fmt.Sprintf("		return new %sOutputData();", dnmn) + "\n" +
		"	}" + "\n" +
		"}"
	return javaContent
}
