// package handlers_test

// import (
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"go.sprint-1/02_net_http/1_example_1/handlers"
// )

// func TestStatusHandler(t *testing.T) {
// 	// определяем структуру теста
// 	type want struct {
// 		code        int
// 		response    string
// 		contentType string
// 	}
// 	// создаём массив тестов: имя и желаемый результат
// 	tests := []struct {
// 		name string
// 		want want
// 	}{
// 		// определяем все тесты
// 		{
// 			name: "positive test #1",
// 			want: want{
// 				code:        200,
// 				response:    `{"status":"ok"}`,
// 				contentType: "application/json",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		// запускаем каждый тест
// 		t.Run(tt.name, func(t *testing.T) {
// 			request := httptest.NewRequest(http.MethodGet, "/status", nil)

// 			// создаём новый Recorder
// 			w := httptest.NewRecorder()
// 			// определяем хендлер
// 			h := http.HandlerFunc(handlers.StatusHandler)
// 			// запускаем сервер
// 			h.ServeHTTP(w, request)
// 			res := w.Result()

// 			// проверяем код ответа
// 			if res.StatusCode != tt.want.code {
// 				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
// 			}

// 			// получаем и проверяем тело запроса
// 			defer res.Body.Close()
// 			resBody, err := io.ReadAll(res.Body)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			if string(resBody) != tt.want.response {
// 				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
// 			}

// 			// заголовок ответа
// 			if res.Header.Get("Content-Type") != tt.want.contentType {
// 				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
// 			}
// 		})
// 	}
// }
