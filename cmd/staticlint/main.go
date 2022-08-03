/*
Собственный multichecker, состоящий из:
* стандартных статических анализаторов пакета golang.org/x/tools/go/analysis/passes;
* всех анализаторов класса SA пакета staticcheck.io;
* не менее одного анализатора остальных классов пакета staticcheck.io;
* двух или более любых публичных анализаторов на ваш выбор.

Добавлен в multichecker собственный анализатор ExitCheckAnalyzer,
запрещающий использовать прямой вызов os.Exit в функции main пакета main.
*/

package main

import (
	"github.com/EestiChameleon/URLShortenerService/cmd/staticlint/exitanalyzer" // как быть?
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
	"log"
)

func main() {
	mychecks := []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		exitanalyzer.ExitCheckAnalyzer, // my own analyser - os.Exit in main check
	}
	// добавляем анализаторы из staticcheck
	for _, v := range staticcheck.Analyzers { // staticcheck.Analyzers contains only SA checks
		mychecks = append(mychecks, v)
	}
	log.Println(mychecks)
	multichecker.Main(
		mychecks...,
	)
}
