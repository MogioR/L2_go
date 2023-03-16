package rest

import (
	"fmt"
	"log"
	"net/http"
	resthendler "task11/internal/transport/rest/v1"
	"time"
)

var (
	mux *http.ServeMux
)

// Logger MiddleWare-функция для логирования запросов
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

// Регистрация хендлеров
func RegisterHendlers() {
	mux = http.NewServeMux()

	mux.HandleFunc("/create_event", Logger(resthendler.CreateEvent))
	mux.HandleFunc("/update_event", Logger(resthendler.UpdateEvent))
	mux.HandleFunc("/delete_event", Logger(resthendler.DeleteEvent))
	mux.HandleFunc("/events_for_day", Logger(resthendler.EventsForDay))
	mux.HandleFunc("/events_for_week", Logger(resthendler.EventsForWeek))
	mux.HandleFunc("/events_for_month", Logger(resthendler.EventsForMonth))
}

// Запуск сервера
func StartServer(port int) {
	log.Println("start server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
