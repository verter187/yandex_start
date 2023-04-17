package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	Person struct {
		Name        string    `json:"Имя"`
		Email       string    `json:"Почта"`
		DateOfBirth time.Time `json:"-"` // - означает, что это поле не будет сериализовано
	}

	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Attributes struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}

	Data struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		Attributes Attributes
	}

	Response struct {
		Header Header `json:"header"`
		Data   []Data `json:"data"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	byt := []byte(rawResp)
	rs := Response{}
	error := json.Unmarshal(byt, &rs)

	return rs, error
}

func main() {
	// man := &Person{
	// 	Name:  "Aлекс",
	// 	Email: "alex@yandex.ru",
	// }
	// jsMan, err := json.Marshal(man)
	// if err != nil {
	// 	log.Fatalln("unable marshal to json")
	// }
	// fmt.Printf("Man %v", string(jsMan)) // Man {"Имя":"Alex","Почта":"alex@yandex.ru"}
	// fmt.Println()

	// // Анонимные структуры
	// req := struct {
	// 	NameContains string `json:"name_contains"`
	// 	Offset       int    `json:"offset"`
	// 	Limit        int    `json: "limit"`
	// }{
	// 	NameContains: "Ivan",
	// 	Limit:        50,
	// }
	// reqRaw, _ := json.Marshal(req)
	// fmt.Println(string(reqRaw))

	// // Размер struct{} равен 0, при этом объект c имеет адрес. Такую лазейку
	// // можно использовать для оптимизации кода по памяти, а в дальнейшем разберём это на практике.
	// // var c struct{}
	// // или
	// c := struct{}{}

	// fmt.Println(unsafe.Sizeof(c))
	// fmt.Println(unsafe.Pointer(&c))

	rawResp := `{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
} `

	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Data[0].Attributes.ArticleIds[1])

}
