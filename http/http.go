package http

import (
	"errors"
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

const (
	BodyJson    = "JSON"
	BodyDefault = "DEFAULT"

	RequestGet  = "GET"
	RequestPost = "POST"
	RequestPut  = "PUT"
)

var ContentType = map[string]string{
	BodyJson:    "application/json",
	BodyDefault: "application/x-www-form-urlencoded",
}

func DoGet(url string, bodyType string, auth string, bodyStr string) (string, error) {
	return Do(RequestGet, url, bodyType, auth, bodyStr)
}

func DoPost(url string, bodyType string, auth string, bodyStr string) (string, error) {
	return Do(RequestPost, url, bodyType, auth, bodyStr)
}

func Do(method string, url string, bodyType string, auth string, bodyStr string) (string, error) {
	log.Println("do request => ", method, url, bodyStr)

	client := &http.Client{}
	req, reqErr := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(bodyStr))

	if reqErr != nil {
		return "", reqErr
	}

	switch strings.ToUpper(bodyType) {
	case BodyJson:
		req.Header.Set("Content-Type", ContentType[BodyJson])
	default:
		req.Header.Set("Content-Type", ContentType[BodyDefault])
	}

	if auth != "" {
		req.Header.Set("Authorization", auth)
	}

	resp, respErr := client.Do(req)

	if respErr != nil {
		return "", respErr
	}

	defer resp.Body.Close()

	body, bodyErr := ioutil.ReadAll(resp.Body)

	if bodyErr != nil {
		return "", bodyErr
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status + "," + string(body))
	}

	return string(body), nil
}
