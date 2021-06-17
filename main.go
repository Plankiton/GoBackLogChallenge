package main
import "github.com/Plankiton/SexPistol"

func main() {
    new(Sex.Pistol).
    Add("/{title}", TitleSearch).
    Run()
}

func TitleSearch (r Sex.Request) (Sex.Json, int) {
    title := Normalize(r.PathVars["title"])

    resp, err := Get("https://swapi.dev/api/films")
    if err != nil {
        return Sex.Bullet {
            Type: "Error",
            Message: "internal server error",
        }, 500
    }

    movies := resp["results"].([]interface{})
    for _, m := range movies {
        movie := m.(map[string]interface{})
        if Normalize(movie["title"].(string)) != title {
            continue
        }

        for prop, value := range movie {
            if urls, ok := value.([]string); ok {
                cat := make([]map[string]interface{}, len(urls))
                for u, url := range urls {
                    data, err := Get(url)
                    if err != nil {
                        continue
                    }

                    cat[u] = data
                }

                movie[prop] = cat
            }
        }

        return movie, Sex.StatusOK
    }


    return Sex.Bullet {
        Type: "Error",
        Message: "movie not found",
    }, 404
}
