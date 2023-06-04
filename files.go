package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

/*Чтение и парсинг файла конфигурации*/
func ReadConf(conf *Config, path string) error {
	file_conf, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file_conf.Close()
	err = xml.NewDecoder(file_conf).Decode(conf)
	if err != nil {
		return err
	}
	return nil
}

/*Запись лога*/
func AddFileLog(log []byte, path string) error {
	file_log, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file_log.Close()
	writer := bufio.NewWriter(file_log)
	for _, buf := range log {
		err = writer.WriteByte(buf)
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

/*Конкэт времени и текста*/
func ConcatLog(text string) []byte {
	return []byte(fmt.Sprintf("%s >> %s", time.Now().Format("02.01.2006 15:04:05"), text))
}
