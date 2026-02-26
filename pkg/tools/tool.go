package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendSMS(templateID, code, phone string) (string, error) {
	url := fmt.Sprintf("https://push.spug.cc/send/%s?code=%s&targets=%s",
		templateID, code, phone)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
