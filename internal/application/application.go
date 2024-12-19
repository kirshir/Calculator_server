package application

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/kirshir/Calculator_server/pkg/calculation"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

// Функция запуска консольного приложения приложения
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
