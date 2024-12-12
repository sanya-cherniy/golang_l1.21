package main

import "fmt"

// Абстрактная структура абстрактной базы данных
type dataBase struct {
}

// Функция запуска
func (db *dataBase) startDB() {
	fmt.Println("databse is running")
}

// Адаптер для базы данных
type dataBaseAdaper struct {
	*dataBase
}

// Адаптированная функция запуска
func (dataBaseAdaper dataBaseAdaper) Run() {
	dataBaseAdaper.dataBase.startDB()
}

// Конструктор адаптера базы данных
func newDataBaseAdaper(dataBase *dataBase) Adapter {
	return &dataBaseAdaper{dataBase: dataBase}
}

// Абстрактная структура абстрактного сервера
type server struct {
}

func (s server) startServer(port string) {
	fmt.Printf("server is running on port: %s", port)
}

type serverAdapter struct {
	*server
	port string
}

// Конструктор адаптера сервера
func (serverAdapter serverAdapter) Run() {
	serverAdapter.server.startServer(serverAdapter.port)
}

func newServerAdapter(server *server, port string) Adapter {
	return &serverAdapter{server: server, port: port}
}

type Adapter interface {
	Run()
}

func main() {
	// Создаем экземпляры адаптируемых структур
	var server server
	var DB dataBase

	//Генерируем массив адаптеров
	infrastructure := [...]Adapter{newDataBaseAdaper(&DB), newServerAdapter(&server, "localhost:8080")}

	// Вызываем у каждого адаптированный метод Run
	for _, elem := range infrastructure {
		elem.Run()
		fmt.Println()
	}

}
