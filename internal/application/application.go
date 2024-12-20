package application

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kirshir/Calculator_server/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type CaclulationRequest struct {
	Expression string `json:"expression"`
}

type CaclulationResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func HandleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) //405
		return
	}

	request := new(CaclulationRequest)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `{"error":"Invalid request format"}`, http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		var status int
		var errMessage string

		switch {
		case errors.Is(err, calculation.ErrInvalidExpression):
			status = http.StatusUnprocessableEntity
			errMessage = calculation.ErrInvalidExpression.Error()
		case errors.Is(err, calculation.ErrInvalidBrackets):
			status = http.StatusUnprocessableEntity
			errMessage = calculation.ErrInvalidBrackets.Error()
		case errors.Is(err, calculation.ErrInvalidCharacter):
			status = http.StatusUnprocessableEntity
			errMessage = calculation.ErrInvalidCharacter.Error()
		case errors.Is(err, calculation.ErrDivisonByZero):
			status = http.StatusUnprocessableEntity
			errMessage = calculation.ErrDivisonByZero.Error()
		default:
			status = http.StatusInternalServerError
			errMessage = "Internal server error"
		}

		response := CaclulationResponse{Error: errMessage}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := CaclulationResponse{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Метод запуска консольного приложения приложения
func (a *Application) Run() error {
	for {
		log.Println("input expression or exit")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}

		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}

		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

// Метод запуска сервера
func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", HandleCalculate)
	log.Printf("Server is running on port %s\n", a.config.Addr)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
