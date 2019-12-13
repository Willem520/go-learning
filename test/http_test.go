package test

import (
	"go-learning/http"
	"log"
	"testing"
)

func TestHttp(t *testing.T) {
	url := "https://www.baidu.com"
	body := ""
	resp := http.DoGet(&url, &body)
	log.Println("======resp:" + resp)
}
