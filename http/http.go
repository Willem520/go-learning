package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
author：willem
description：http请求
date：2019-12-10 15:42
*/

func DoGet(url *string, bodyStr *string) string {
	return Do("GET", url, bodyStr)
}

func DoPost(url *string, bodyStr *string) string {
	return Do("POST", url, bodyStr)
}

func Do(method string, url *string, bodyStr *string) string {
	client := &http.Client{}

	req, reqErr := http.NewRequest(strings.ToUpper(method), *url, strings.NewReader(*bodyStr))

	if reqErr != nil {
		log.Fatalln("error")
		return ""
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := client.Do(req)

	if respErr != nil {
		log.Fatalln("error")
		return ""
	}

	defer resp.Body.Close()

	body, bodyErr := ioutil.ReadAll(resp.Body)

	if bodyErr != nil {
		log.Fatalln(bodyErr)
		return ""
	}

	return string(body)
}
