package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Hello struct{
	Name string `json:"name"`
}

type BattleRequest struct {
	MonsterA uint32
	MonsterB  uint32
	Address string
	BattleLevel uint32

}

func main() {

	//values := map[string]string{"name": "username"}
	//jsonValue, _ := json.Marshal(values)

	//m := make(map[string]interface{})
	s := &Hello{
		Name: "aaaa",
	}
	//
	bod := &bytes.Buffer{}
	json.NewEncoder(bod).Encode(s)

	request, _ := http.NewRequest("POST", "http://192.168.50.16:809/hello", bod)
	request.Header.Set("Content-Type", "application/json")

	client:=&http.Client{}
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf(string(body))

}
