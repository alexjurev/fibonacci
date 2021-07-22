package apiserver

import (
	"encoding/json"
	"fibo/pkg"
	"net/http"

)

type InputValues struct {
	FirstValue int `json:"first"`
	SecondValue int `json:"second"`
	
}

type Result struct {
	FirstValue []int 
	SecondValue []int 
}

func FiboCounter(w http.ResponseWriter, r *http.Request) {
	inputValues := InputValues{}
	result := Result{}	
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&inputValues)
	
	if r.Method == "POST" {
		result.FirstValue, result.SecondValue = pkg.FibSlice(inputValues.FirstValue, inputValues.SecondValue)
		}
		json.NewEncoder(w).Encode(result)
	}
	
