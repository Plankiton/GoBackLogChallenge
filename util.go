package main

import (
    "github.com/Plankiton/SexPistol"

    "io/ioutil"
    "net/http"

    "strings"
)

func Get(u string, v interface{}) (error, int) {

    resp, err := http.Get(u)

    if err == nil {
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if call_back, ok := v.(func (tmp interface{})); ok {
            var data map[string]interface{}
            err = Sex.FromJson(body, &data)
            call_back(data)
        } else {
            err = Sex.FromJson(body, v)
        }

        return err, 200
    }

    return err, 500
}

func Index(s string, s2 string) int {
    return strings.Index(s, s2)
}

func Normalize(t string) string {
    t = strings.ToLower(t)
    t = strings.Replace(t, "_", "+", -1)
    t = strings.Trim(t, " ")

    return Sex.Fmt("https://www.swapi.tech/api/%s", t)
}

