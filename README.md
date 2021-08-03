# fibonacci
fibonacci REST gRPC with redis

## Порядок запуска:
1. запустить образ redis в docker :
docker run —name rdb -p 6379:6379 redis
2. запустить main.go - параллельно запускаются 2 сервера Rest и gRPC.

# Доступ по Rest:
через Postman http://127.0.0.1:8000/fibo POST-запрос типа JSON {"first": n1,"second": n2}
![Postman](https://github.com/alexjurev/fibonacci/raw/main/postman.png)

# Доступ по gRPC:
через утилиту Evans в командной строке 
evans fibo/proto/fibonacci.proto -p 8080
далее вызов функции FiboCount
![Evans](https://github.com/alexjurev/fibonacci/raw/main/evans.png)

