package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		err          error
		fileContents string
	)

	if len(os.Args) < 2 {
		fmt.Printf("no arguments supplied in pre-commit check\n")
		os.Exit(1)
	}

	fileName := os.Args[1]

	if fileContents, err = readFile(fileName); err != nil {
		fmt.Printf("Error reading file for security check: %s\n", err.Error())
		os.Exit(1)
	}

	if isEnvFile(fileName) {
		fmt.Printf("You are not allowed to commit .env files. This poses a security risk.\n")
		os.Exit(1)
	}

	if containsApiKey(fileContents) {
		fmt.Printf("It looks like you are attempting to set an API key in %s. This is not allowed.\n", fileName)
		os.Exit(1)
	}

	if containsShopifyKeys(fileContents) {
		fmt.Printf("It looks like you are attempting to set Shopify credentials in %s. This is not allowed.\n", fileName)
		os.Exit(1)
	}

	if containsGithubToken(fileContents) {
		fmt.Printf("It looks like you are attemting to set a GITHUB_TOKEN in %s. This is not allowed.\n", fileName)
		os.Exit(1)
	}

	os.Exit(0)
}

func readFile(fileName string) (string, error) {
	var (
		err error
		b   []byte
	)

	if b, err = os.ReadFile(fileName); err != nil {
		return "", err
	}

	return string(b), nil
}

func isEnvFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".env")
}

func containsApiKey(fileContents string) bool {
	r, _ := regexp.Compile("(?i)api_key\\s*(=|:=|==|===)\\s*[\"'`]{1}.*?[\"'`]")
	return r.MatchString(fileContents)
}

func containsShopifyKeys(fileContents string) bool {
	r, _ := regexp.Compile("(?i)shppa_.*?\\s*(=|:=|==|===)\\s*[\"'`]{1}.*?[\"'`]")
	return r.MatchString(fileContents)
}

func containsGithubToken(fileContents string) bool {
	r, _ := regexp.Compile("(?i)github_token\\s*(=|:=|==|===)\\s*[\"'`]{1}.*?[\"'`]")
	return r.MatchString(fileContents)
}
