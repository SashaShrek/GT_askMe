package main

import (
	"fmt"
	"os/exec"
	"strings"
)

/*Подготовка старта тасков*/
func RunTasks(tsk []Task, text *string) {
	for _, task := range tsk {
		err := StartTask(task.Command, task.Name, task.SysFlg)
		if err != nil {
			*text = fmt.Sprintf("Произошла ошибка при запуске таска. %s\n", err.Error())
			fmt.Printf("%s", *text)
			err = AddFileLog(ConcatLog(*text), pathLog)
			if err != nil {
				fmt.Printf("При записи лога произошла ошибка: %s\n", err.Error())
				return
			}
			continue
		}
	}
}

/*Запуск таска*/
func StartTask(task string, name_task string, param string) error {
	var cmd *exec.Cmd
	var err error
	fmt.Printf("Запускаю задачу %s\n", name_task)
	err = AddFileLog(ConcatLog(fmt.Sprintf("%s\n", name_task)), pathLog)
	if err != nil {
		return err
	}
	if param == "Y" {
		cmd = exec.Command("cmd", "/C", task)
		stdout, err := cmd.CombinedOutput()
		//err = cmd.Run()
		if err != nil {
			return err
		}
		fmt.Printf("%s", stdout)
	} else {
		cmd = exec.Command(task)
		err = cmd.Start()
		if err != nil {
			return err
		}
	}
	fmt.Printf("Задача %s успешно запущена\n\n", name_task)
	return nil
}

/*Парсинг структуры, тег ActionDefault*/
func ParseTasks(tsks string) []string {
	return strings.Split(tsks, ";")
}

/*Проверка и подготовка тасков, указанных к запуску в теге ActionDefault*/
func CheckTask(tsks []string, tsk []Task) []Task {
	var tasks []Task
	var task Task
	var value bool
	for _, elem := range tsks {
		task, value = SearchTask(elem, tsk)
		if value == false {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}

/*?*/
func ScanTasks(tsks []Task) bool {
	for _, val := range tsks {
		if val.Name == "" || len(strings.Split(val.Name, ";")) > 1 {
			return false
		}
	}
	return true
}

/*Поиск таска по имени*/
func SearchTask(name_task string, tsk []Task) (Task, bool) {
	var task Task
	for _, elem := range tsk {
		if elem.Name == name_task {
			task = elem
			break
		}
	}
	if len(task.Name) == 0 {
		return task, false
	}
	return task, true
}
