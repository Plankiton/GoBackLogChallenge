package main
import (
    "github.com/Plankiton/SexPistol"
)

func main() {
    new(Sex.Pistol).
    Add("/{title}", TitleSearch).
    Run()
}

func TitleSearch (r Sex.Request) (Sex.Json, int) {
    title := Normalize(r.PathVars["title"])

    var res map[string]interface{}
    if err, status := Get(Sex.Fmt("films/?search=%s", title), &res); err != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: Sex.Fmt("%v", err),
            Data: res,
        }, status
    }

    if _, ok := res["result"]; !ok {
        return Sex.Bullet {
            Type: "Error",
            Message: "Error on load of data from https://swapi.dev",
            Data: res,
        }, 500
    }

    if res["count"].(float64) > 1 || res["count"].(float64) < 1 {
        return Sex.Bullet {
            Type: "Error",
            Message: "Movie not found",
        }, Sex.StatusNotFound
    }

    movie := res["result"].([]interface{})[0].(map[string]interface{})
    for prop, value := range movie {
        if links, ok := value.([]interface{}); ok {
            new_link_list := make([]interface{}, len(links))

            for l, _link := range links {
                link := _link.(string)
                Get(link[Index(link, prop):], new_link_list[l])
            }

            movie[prop] = new_link_list
        }
    }

    return res, Sex.StatusOK
}
