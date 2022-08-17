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
	"github.com/EestiChameleon/URLShortenerService/cmd/staticlint/exitanalyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	// определяем map подключаемых правил
	mychecks := []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		exitanalyzer.ExitCheckAnalyzer, // my own analyser - os.Exit in main check
	}
	for _, v := range staticcheck.Analyzers {
		// добавляем в массив нужные проверки
		mychecks = append(mychecks, v.Analyzer) // staticcheck.Analyzers contains only SA checks
	}
	multichecker.Main(
		mychecks...,
	)
}