package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/* JSON <-> Go
true/false - true/false
string - string
null - nil
number - float64(default number type), float32, int8, int16, int32, int64, int, uint8, ...
array - []any (known before go 1.9 as the ([]interface{}))
object - map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/

type Reply struct {
	Name         string
	Public_Repos int
}

func main() {
	resp, err := http.Get("https://api.github.com/users/AlexTLDR")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	/*
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatalf("error: can't copy - %s", err)
		}
	*/
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Println(r)
}
