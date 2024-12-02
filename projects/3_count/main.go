package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	counter int        // Переменная-счетчик
	mu      sync.Mutex // Мьютекс для потокобезопасных операций
)

func main() {
	http.HandleFunc("/count", countHandler)

	// Запускаем сервер на порту 3333
	fmt.Println("Сервер запущен на порту :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// Обработчик для маршрута /count
func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // Обрабатываем GET запрос
		mu.Lock()
		fmt.Fprintf(w, "Текущий счетчик: %d", counter)
		mu.Unlock()

	case http.MethodPost: // Обрабатываем POST запрос
		if err := r.ParseForm(); err != nil { // Парсим данные формы
			http.Error(w, "Ошибка обработки формы", http.StatusBadRequest)
			return
		}
		fmt.Print("hello")

		// Получаем значение из ключа "count"
		countValue := r.FormValue("count")
		// Преобразуем строку в число
		countInt, err := strconv.Atoi(countValue)
		if err != nil { // Если преобразование прошло с ошибкой
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}

		// Увеличиваем счетчик
		mu.Lock()
		counter += countInt
		mu.Unlock()

		fmt.Fprintf(w, "Счетчик увеличен на %d. Текущий счетчик: %d", countInt, counter)
	default: // Если метод не поддерживается
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
