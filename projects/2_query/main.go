package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			http.Error(w, "Parameter 'name' is required", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Hello,%s!", name)
		fmt.Fprintln(w, response)
	})

	fmt.Println("Сервер запущен на порту :9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
