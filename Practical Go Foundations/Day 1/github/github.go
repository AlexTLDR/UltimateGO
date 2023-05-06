package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* JSON <-> Go
true/false - true/false
string - string
null - nil
number - float64(default number type), float32, int8, int16, int32, int64, int, uint8, ...
array - []any ([]interface{})
object - map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/

type Reply struct {
	Name string
	//Public_Repos int
	NumRepos int `json:"public_repos"`
}

func main() {
	s := "AlexTLDR"
	fmt.Println(githubInfo(s))
}

func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		//log.Fatalf("error: %s", err)
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	//fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	/*
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatalf("error: can't copy - %s", err)
		}
	*/
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		//log.Fatalf("error: can't decode - %s", err)
		return "", 0, err
	}
	//fmt.Println(r)
	//fmt.Printf("%#v\n", r)
	return r.Name, r.NumRepos, err
}
