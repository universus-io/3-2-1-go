package calculator

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type ExpHandler struct {
	calculator *Calculator
}

func NewExpHandler(calculator *Calculator) *ExpHandler {
	return &ExpHandler{calculator}
}

func (h *ExpHandler) parseUrlPath(path string) float64 {

	s := strings.Replace(path, "/exp/", "", 1)

	x, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return math.NaN()
	}

	return x
}

func (h *ExpHandler) handle(path string) string {

	x := h.parseUrlPath(path)

	log.Printf("[ExpHandler] x=%f", x)

	var response string

	if math.IsNaN(x) {

		response = "{Error, unable to parse '" + path + "'"

	} else {

		result := h.calculator.Exp(x)

		response = "{Ok, result=" + strconv.FormatFloat(result, 'f', 8, 64) + "}"

	}

	log.Printf("[ExpHandler] response=%s", response)

	return response
}

func (h *ExpHandler) Handle(writer http.ResponseWriter, request *http.Request) {

	log.Print("[ExpHandler] handle() >")

	path := request.URL.Path

	response := h.handle(path)

	_, _ = fmt.Fprint(writer, response)

	log.Print("[ExpHandler] handle() <")
}
