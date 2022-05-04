package random_text_api

import (
	"io/ioutil"
	"net/http"
)

const apiUrl = "https://baconipsum.com/api/?type=meat-and-filler"

func GetText() string {
	resp, _ := http.DefaultClient.Get(apiUrl)
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}
