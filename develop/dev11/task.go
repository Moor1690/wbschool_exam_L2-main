/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Event представляет событие в календаре.
// Оно содержит ID, название и дату события.
type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

// Response представляет стандартный ответ сервера.
// Он может содержать результат операции или сообщение об ошибке.
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Функция для записи JSON ответа
func writeJSONResponse(w http.ResponseWriter, code int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// Обработчики запросов
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != "POST" {
		writeJSONResponse(w, http.StatusMethodNotAllowed, Response{Error: "Only POST method is allowed"})
		return
	}

	// Парсинг параметров из тела запроса
	if err := r.ParseForm(); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid request"})
		return
	}

	// Создание структуры события из параметров
	event := Event{
		Title: r.FormValue("title"),
		Date:  r.FormValue("date"),
	}

	// Валидация параметров события (примерная реализация)
	if event.Title == "" || event.Date == "" {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Missing title or date"})
		return
	}

	event.ID = nextEventID
	nextEventID++
	eventsStore[event.ID] = event

	// Возврат ответа
	writeJSONResponse(w, http.StatusOK, Response{Result: "Event created with ID: " + strconv.Itoa(event.ID)})
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != "POST" {
		writeJSONResponse(w, http.StatusMethodNotAllowed, Response{Error: "Only POST method is allowed"})
		return
	}

	// Парсинг тела запроса
	if err := r.ParseForm(); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid request"})
		return
	}

	// Получение и валидация параметров
	idStr := r.FormValue("id")
	title := r.FormValue("title")
	date := r.FormValue("date")

	// Преобразование id в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid ID"})
		return
	}

	// Проверка остальных параметров
	if title == "" || date == "" {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Missing title or date"})
		return
	}

	fmt.Println("updateEventHandler ", id)

	if _, exists := eventsStore[id]; !exists {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Event not found"})
		return
	}
	eventsStore[id] = Event{ID: id, Title: title, Date: date}

	// Возврат успешного ответа
	writeJSONResponse(w, http.StatusOK, Response{Result: "Event updated"})
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != "POST" {
		writeJSONResponse(w, http.StatusMethodNotAllowed, Response{Error: "Only POST method is allowed"})
		return
	}

	// Парсинг тела запроса
	if err := r.ParseForm(); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid request"})
		return
	}

	// Получение и валидация параметра id
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid ID"})
		return
	}

	fmt.Println("deleteEventHandler ", id)

	if _, exists := eventsStore[id]; !exists {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Event not found"})
		return
	}
	delete(eventsStore, id)

	// Возврат успешного ответа
	writeJSONResponse(w, http.StatusOK, Response{Result: "Event deleted"})
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг даты из запроса
	dateStr := r.URL.Query().Get("date")
	date, err := parseDate(dateStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid date format"})
		return
	}

	var eventsList []Event
	for _, event := range eventsStore {
		eventDate, _ := parseDate(event.Date)
		if eventDate == date {
			eventsList = append(eventsList, event)
		}
	}

	// Сериализация списка событий и отправка в ответе
	writeJSONResponse(w, http.StatusOK, Response{Result: fmt.Sprintf("Events for day: %v", eventsList)})
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	startDate, err := parseDate(dateStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid date format"})
		return
	}
	endDate := startDate.AddDate(0, 0, 7)

	var eventsList []Event
	for _, event := range eventsStore {
		eventDate, _ := parseDate(event.Date)
		if eventDate.After(startDate) && eventDate.Before(endDate) {
			eventsList = append(eventsList, event)
		}
	}

	writeJSONResponse(w, http.StatusOK, Response{Result: fmt.Sprintf("Events for week: %v", eventsList)})
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	startDate, err := parseDate(dateStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, Response{Error: "Invalid date format"})
		return
	}
	endDate := startDate.AddDate(0, 1, 0)

	var eventsList []Event
	for _, event := range eventsStore {
		eventDate, _ := parseDate(event.Date)
		if eventDate.After(startDate) && eventDate.Before(endDate) {
			eventsList = append(eventsList, event)
		}
	}

	writeJSONResponse(w, http.StatusOK, Response{Result: fmt.Sprintf("Events for month: %v", eventsList)})
}

var eventsStore = make(map[int]Event)
var nextEventID = 1

func main() {
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	http.Handle("/", loggingMiddleware(http.DefaultServeMux))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
