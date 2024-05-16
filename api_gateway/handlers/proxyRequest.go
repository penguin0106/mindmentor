package handlers

import (
	"io"
	"net/http"
)

func proxyRequest(w http.ResponseWriter, url string, method string, body io.Reader) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		http.Error(w, "Ошибка создания запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Не удалось выполнить запрос к удаленному сервису: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копирование заголовков ответа
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
