[![CodeFactor](https://www.codefactor.io/repository/github/sashashrek/gt_askme/badge)](https://www.codefactor.io/repository/github/sashashrek/gt_askme)

# GT_askMe
Запуск различных тасков и различного их количества

### Суть
Данная утилита позволяет запустить несколько программ (тасков). Настройка происходит через конфигурационный файл

### Принцип работы
Всё начинается с настройки. В данном случае это конфигурационный файл _config.xml_. Его структура примерно такая:
```xml
<?xml version="1.0"?>
<Config>
    <ActionDefault>all</ActionDefault>
    <Tasks>
        <Task>
            <Name>Out</Name>
            <Command>echo Hello!</Command>
            <SystemFlag>Y</SystemFlag>
        </Task>
        <Task>
            <Name>Test</Name>
            <Command>notepad</Command>
            <SystemFlag>N</SystemFlag>
        </Task>
    </Tasks>
</Config>
```

- Тег "**ActionDefault**": указывает какие таски должны запускаться. Если указанно **all**, то запустить все таски в блоке **Tasks**. Если нужно запустить какие-то отдельные таски, то в этом теге нужно указать имя этого таска. 
  Если требуется запустить несколько тасков - их названия должны прописываться через ";" без пробелов: ```<ActionDefault>Test;Out</ActionDefault>```
- Список для запуска тасков указан в блоке **Tasks**. Каждый таск имеет следующую структуру:
  ```xml
  <Task>
    <Name>Out</Name>
    <Command>echo Hello!</Command>
    <SystemFlag>Y</SystemFlag>
  </Task>
  ```
  где: **Name** - название таска, не должно содержать ";"; **Command** - команда, которую нужно запустить; **SystemFlag** - если Y, то запустить команду из-под системного терминала. Если N - просто выполнить команду
  
  Блоков **Task** может быть сколько угодно
  
  Сама утилита запускается без флагов, вся настройка происходит через конфигурационный файл
