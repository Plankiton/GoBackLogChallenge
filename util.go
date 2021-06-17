package main

import (
    "io/ioutil"
    "net/http"

    "strings"
)

func Get(url string) (map[string]interface{}, error) {
    var json_body map[string] interface{}
    resp, _ := http.Get(url)
    body, err = ioutil.ReadAll(resp.Body)
    Sex.FromJson(body, &json_body)

    return json_body, err
}

func Normalize(t string) string {
    t = strings.Lower(t)
    t = strings.Replace(t, "_", " ")
    t = strings.Trim(t, " ")
}

