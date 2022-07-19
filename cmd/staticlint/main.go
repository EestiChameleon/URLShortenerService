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
	"encoding/json"
	"github.com/EestiChameleon/URLShortenerService/cmd/staticlint/exitanalyzer" // как быть?
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
	"os"
	"path/filepath"
)

// Config — имя файла конфигурации.
const Config = `config.json`

// ConfigData описывает структуру файла конфигурации.
type ConfigData struct {
	Staticcheck []string
}

func main() {
	// Read and unmarshal the staticcheck config
	appfile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(filepath.Join(filepath.Dir(appfile), Config))
	if err != nil {
		panic(err)
	}
	var cfg ConfigData
	if err = json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	mychecks := []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		shift.Analyzer,
		structtag.Analyzer,
		exitanalyzer.ExitCheckAnalyzer, // my own analyser - os.Exit in main check
	}
	checks := make(map[string]bool)
	for _, v := range cfg.Staticcheck {
		checks[v] = true
	}
	// добавляем анализаторы из staticcheck, которые указаны в файле конфигурации
	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	multichecker.Main(
		mychecks...,
	)
}
