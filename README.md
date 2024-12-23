# go-calculator

Калькулятор на языке Go, сделанный в рамках курса Яндекс Лицея "Программирование на Go".
Модуль вычислений реализован с помощью польской нотации (https://ru.wikipedia.org/wiki/Польская_запись).

## Что поддерживается
Веб-сервис с 1 эндпойнтом с url `/api/v1/calculate`.
Взаимодействие с ним через POST-запрос с телом
`{
"expression": "ВЫРАЖЕНИЕ_ДЛЯ_ВЫЧИСЛЕНИЯ"
}`
Калькулятор поддерживает операции сложения (+), вычитания (-), умножения (*) и деления (/) **только положительных однозначных чисел** (от 0 до 9 включительно). Поддерживаются круглые скобки ( и ) для определения порядка выполнения операций.

## Примеры
Выдается код 200 и ответ с телом
`{
"result": "результат выражения"
}`
, если выражение вычислено успешно.
`curl --location 'localhost:8082/api/v1/calculate' --header 'Content-Type:' --data '{n/json'
"expression": "2+2*4"
}'`

Выдается код 422 и ответ с телом
`{
"error": "Expression is not valid"
}`
, если входные данные не соответствуют требованиям приложения (двузначные цифры, неподдерживаемые операции, незакрытые скобки, тело post-запроса не содержит ключ expression).
`curl --location 'localhost:8082/api/v1/calculate' --header 'Content-Type:' --data '{n/json'
"expression": "2+2*24"
}'`

Выдается код 500 и ответ с телом
`{
"error": "Internal server error"
}`
, если запрос не является POST или иная ошибка приложения (не передали данные вообще).
`curl --location 'localhost:8082/api/v1/calculate' --header 'Content-Type: application/json' --data ''`

## Инструкция по запуску
1. Склонируйте себе репозиторий. Результатом будет папка с кодом под названием go-calculator. 
2. Откройте терминал. Перейдите в папку с пакетом
   `cd go-calculator/rpn/`
3. Запустите программу
   `export PORT=8082 && go run ./cmd/main.go`
   
4. Чтобы выполнить запрос к приложению, откройте второе окно терминала и выполните следующий запрос через curl
   `curl --location 'localhost:8082/api/v1/calculate' --header 'Content-Type: application/json' --data '{
   "expression": "2+2*2"
   }'`
   
