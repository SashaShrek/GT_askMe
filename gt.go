package main

import (
	"encoding/xml"
	"fmt"
)

type (
	Task struct {
		Name    string `xml:"Name"`
		Command string `xml:"Command"`
		SysFlg  string `xml:"SystemFlag"`
	}

	Tasks struct {
		Tsk []Task `xml:"Task"`
	}

	Config struct {
		XMLName xml.Name `xml:"Config"`
		ActDef  string   `xml:"ActionDefault"`
		Tsks    Tasks    `xml:"Tasks"`
	}
)

var (
	pathLog  string = "output.log"
	pathConf string = "config.xml"
)

func main() {
	var text string
	text = "Читаю файл конфигурации...\n"
	fmt.Println(text)
	err := AddFileLog(ConcatLog(text), pathLog)
	if err != nil {
		fmt.Printf("При записи лога произошла ошибка: %s\n", err.Error())
		return
	}
	var conf Config
	err = ReadConf(&conf, pathConf)
	if err != nil {
		fmt.Printf("Произошла ошибка при чтении файла конфигурации: %s", err.Error())
		return
	}
	if ScanTasks(conf.Tsks.Tsk) == false {
		fmt.Println("Некорректное имя таска")
		return
	}

	text = fmt.Sprintf("Прочитано. Запускаю указанные задачи...\n")
	fmt.Printf("\n")
	err = AddFileLog(ConcatLog(text), pathLog)
	if err != nil {
		fmt.Printf("При записи лога произошла ошибка: %s\n", err.Error())
		return
	}
	if conf.ActDef == "all" {
		fmt.Printf("Всего: %d задач\n", len(conf.Tsks.Tsk))
		RunTasks(conf.Tsks.Tsk, &text)
	} else {
		check := CheckTask(ParseTasks(conf.ActDef), conf.Tsks.Tsk)
		RunTasks(check, &text)
	}
	fmt.Println("Нажми Enter для завершения...")
	fmt.Scanf("%s")
}
