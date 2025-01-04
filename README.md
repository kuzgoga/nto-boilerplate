# NTO

## Задачи
- [x] Migrator
- [ ] Crud service generator (based on gorm gen)
- [ ] Easy init without git repository (like create-nto-app)
- [ ] Excel export
- [ ] Models validation logic (callbacks and other stuff)
- [ ] Models [linter](https://git.gogacoder.ru/NTO/gormlint)
- [ ] Auto-generated frontend  

## Установка

Для разработки необходимы следующие инструменты:
* [Node.js](https://nodejs.org/en)
* [Golang](https://go.dev/dl/) 
* VSCode или Goland
* [UPX](https://github.com/upx/upx/releases/latest) - для сжатия.  
**Необходимо зазеркалировать с Github и добавить в PATH!!!**
* [Git](https://git-scm.com/)
* Wails3:  
Установка:
`go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest`  
**Необходимо зазеркалировать с Github и добавить в PATH!!!**
* MinGw-64: https://jmeubank.github.io/tdm-gcc/  
**Необходимо зазеркалировать с Github!!!**

## Разработка
### Hot Reload
Для запуска приложения в режиме разработчика используйте эту команду в директории проекта:
```
wails3 dev
```  
DevServer также можно открыть по адресу http://localhost:34115.  
Этот инструмент предоставляет возможность вызывать Go код прямо из инструментов разработчика.

### Генерация TS биндингов
Для обновления API для TypeScript используйте команду:
```
wails3 generate bindings -ts
```

## Сборка

Для финальной сборки запустите эту команду в директории проекта:
```
go env -w CGO_ENABLED=1
wails3 build -clean -upx -v 2 -webview2 embed
```
**Перед релизом не забыть**:
* поместить все нужные asset'ы в папку assets
* изменить версию схемы БД (пока не нужно)
* приложить сопроводительную записку.

## Работа без GitHub
Настройте прокси для скачивания зависимостей через прокси:
```
go env -w GOPROXY="https://proxy.golang.org,direct"
```

