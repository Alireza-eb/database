package main 

import 
"bytes"
"encoding/json"
"io/ioutil"
"log"
"net/http"

func main()  {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Alireza",
		"email": "ebrahimi.alireza@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://randomuser.me/api//post", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()
	 body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
}