# test for untypical

### Тестовое задание:
"Необходимо создать Http сервис - key-value хранилище. 
### Сервис должен содержать четыре метода в апи:
 - Upsert (вставка либо обновление)
 - Delete
 - Get
 - List

 Хранить данные можно просто в оперативной памяти при помощи map."

### Для запуска:
* `$ go run cmd/test_untypical/main.go`

### Примеры запросов:
По умолчанию запускается на порту :8080, можно изменить в конфиге config/config.yaml

//key и value могут быть как числами, так  и строкой
* Upsert: `curl -X POST --data '{"key":"xyz","value":"valueXYZ"}' "http://localhost:8080/upsert"`


* Get: `curl -X GET http://localhost:8080/get?key=xyz"`
  

* List: `curl -X GET "http://localhost:8080/list"`
  

* Delete: `curl -X DELETE "http://localhost:8080/delete?key=xyz"`
