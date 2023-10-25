package app

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)
var (
	domain string
	method string
	dn string // domainNameの最初を大文字へ変換
	mn string // methodNameの最初を大文字へ変換
	dnmn string
)

func SetCommonVariable(domainName string, methodName string) {
	domain = domainName
	method = methodName
	dn = firstLetterToUpper(domainName)
	mn = firstLetterToUpper(methodName)
	dnmn = dn + mn
}

func firstLetterToUpper(s string) string {
	if s == "" {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	return fmt.Sprintf("%c%s", unicode.ToUpper(r), s[size:])
}

func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// ディレクトリが存在しない場合、ディレクトリを作成
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
