package main

import (
    "github.com/Plankiton/SexPistol"

    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

func Get(u string, v interface{}) (error, int) {
    var resp *http.Response
    var body []byte
    var err error

    resp, err = http.Get("https://swapi.dev/api/"+ url.QueryEscape(u))

    if err != nil {
        body, err = ioutil.ReadAll(resp.Body)
        Sex.SuperPut(u, string(body))
        Sex.FromJson(body, v)
    }

    return err, resp.StatusCode
}

func Index(s string, s2 string) int {
    return strings.Index(s, s2)
}

func Normalize(t string) string {
    t = strings.ToLower(t)
    t = strings.Replace(t, "_", "+", -1)
    t = strings.Trim(t, " ")

    return t
}

