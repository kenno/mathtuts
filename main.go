package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type studentStruct struct {
	Id       string `json:"id"`
	FullName string `json:"fullname"`
	TutCode  string `json:"tutcode"`
}

const (
	urlEndPoint = "http://mathlms1:8181/students"
)

func main() {
	user := os.Getenv("USER") // return z1234567
	if user != "" {
		i, err := strconv.Atoi(user[1:]) // remove 'z'
		if err != nil {
			log.Fatal(err)
		}
		getContent(i)
	}
}

func getContent(id int) {
	res, err := http.Get(fmt.Sprintf("%s/%d", urlEndPoint, id))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	student := studentStruct{}
	json.NewDecoder(res.Body).Decode(&student)

	if (studentStruct{}) == student {
		fmt.Printf("No tutorial code found for student with ID: %d\n", id)
	} else {
		fmt.Printf("Given name: %s\nLogin: %s\nTutorial code: %s\n",
			student.FullName,
			student.Id,
			student.TutCode)
	}

}
