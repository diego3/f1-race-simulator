package gamelogic

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

func Random(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func WriteJson(fileName string, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		log.Println("Error trying to marshall json: ", err)
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Println("Error trying to create file: ", err)
	}
	f.WriteString(string(bytes))
}
