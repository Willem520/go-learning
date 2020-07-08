package test

import (
	"go-learning/http"
	"log"
	"testing"
)

func TestHttp(t *testing.T) {
	url := "https://www.baidu.com"
	body := ""
	resp, _ := http.DoGet(url, "json", "", body)
	log.Println("======resp:" + resp)
}
