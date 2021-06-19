package main
import "github.com/Plankiton/SexPistol"

func main() {
    new(Sex.Pistol).
    Add("/{title}", TitleSearch).
    Run()
}

func TitleSearch (r Sex.Request) (Sex.Json, int) {
    title := r.PathVars["title"]

    var res map[string]interface{}
    if err, status := Get( Normalize(Sex.Fmt("films/?search=%s", title)), &res); err != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: Sex.Fmt("%v", err),
            Data: res,
        }, status
    }

    if _, ok := res["result"]; !ok {
        return Sex.Bullet {
            Type: "Error",
            Message: "Error on load of data from https://www.swapi.tech",
            Data: res,
        }, 500
    }

    result := res["result"].([]interface{})
    if len(result) > 1 || len(result) < 1 {
        return Sex.Bullet {
            Type: "Error",
            Message: "Movie not found",
        }, Sex.StatusNotFound
    }


    props := 0
    lock := []byte{}
    movie := result[0].(map[string]interface{})["properties"].(map[string]interface{})
    for prop, value := range movie {
        if links, ok := value.([]interface{}); ok {
            movie[prop] = []map[string]interface{}{}

            for _, _link := range links {
                link := _link.(string)

                go Get(link, func (tmp interface{}, v...interface{}) {
                    if v != nil {
                        prop := v[0].(string)
                        if tmp != nil {
                            res := tmp.(map[string]interface{})
                            result := res["result"]
                            if fixed_list, ok := result.(map[string]interface{})["properties"].(map[string]interface{}); ok {
                                for p, v := range fixed_list {
                                    if _, ok := v.([]interface{}); ok {
                                        fixed_list[p] = nil
                                    }
                                }
                                movie[prop] = append(movie[prop].([]map[string]interface{}), fixed_list)
                            }
                        }
                        lock = append(lock, 1)
                    }
                }, prop)
            }
            props ++
        }
    }
    for len(lock) < props {}

    return movie, Sex.StatusOK
}
