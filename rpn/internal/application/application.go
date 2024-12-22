package application

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/nat-kpt/rpn/pkg/rpn"
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

// Функция запуска приложения
// тут будем читать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("application was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := rpn.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed with error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type ResponseOK struct {
	Result string `json:"result"`
}

type ResponseNotOK struct {
	Error string `json:"error"`
}
var answerOk = ResponseOK{}
var answerErr = ResponseNotOK{}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	// ошибку 500 бросаем, когда не пост метод или ничего не передали
	if err != nil || r.Method != "POST" {
		answerErr.Error = "Internal server error"
		responseBytes, _ := json.Marshal(answerErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responseBytes)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	result, err := rpn.Calc(request.Expression)
	if err != nil {
		// 422 кидаем, когда калькулятор не поддерживает такие выражения
		// либо передали неправильную json
		answerErr.Error = "Expression is not valid"
		responseBytes, _ := json.Marshal(answerErr)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(responseBytes)
		http.Error(w, "", http.StatusUnprocessableEntity)
	} else {
		// 200 когда все ок
		answerOk.Result = strconv.FormatFloat(result, 'f', 2, 64)
		responseBytes, _ := json.Marshal(answerOk)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
