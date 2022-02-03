package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(res *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(res.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
