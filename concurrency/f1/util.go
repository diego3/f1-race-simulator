package f1

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

func random(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func writeJson(fileName string, data interface{}) {
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
