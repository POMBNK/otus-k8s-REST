# Kuberrest

Для запуска необходимо иметь установленный **minikube**, **kubectl**, **helm**. Также для удобства используйте утилиту **make**

Порядок запуска:

1. Склонировать репозиторий
```shell
    git clone https://github.com/POMBNK/otus-k8s-REST && cd otus-k8s-REST
``` 
2. Добавить в файл hosts 127.0.0.1 arch.homework, затем для развертывания приложения выполнить команду. После выполнения откроется дашборд
```shell
make start
```

3. Открыть доступ через loopback interface
```shell
  make tun
```
4. Перейти на http://arch.homework/swagger/ Здесь вы найдете документацию к API, а также подставленные примеры запросов, достаточно только их вызвать через UI


5. По окончанию работы, очистить поды + джобы
```shell
make done
```