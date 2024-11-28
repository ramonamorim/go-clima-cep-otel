package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ramonamorim/go-clima-cep-otel/service-a-cep-input/internal/service"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Cep struct {
	Cep string `json:"cep"`
}

func HandleCepRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	var cep Cep
	err = json.Unmarshal(body, &cep)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	ctx, span := otel.Tracer("service-a").Start(r.Context(), "validate-cep")
	span.SetAttributes(attribute.String("cep", cep.Cep))
	defer span.End()

	if !isValidZipcode(cep.Cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	temperature, status, err := service.GetTemperature(cep.Cep, ctx)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	jsonData, err := json.Marshal(temperature)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func isValidZipcode(zipcode string) bool {
	if len(zipcode) != 8 {
		return false
	}
	for _, char := range zipcode {
		if _, err := strconv.Atoi(string(char)); err != nil {
			return false
		}
	}
	return true
}
