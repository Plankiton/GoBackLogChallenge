<img src="img/image1.png" style="background-color: white;border-radius: 5px"/>

# Disclaimer

O código era pra ficar mais simples que isso, no inicio criei uma função pra fazer a requisição e parsear a resposta, além de todas as requisições sincronamente, uma atrás da outra e por fim, retornar o valor final. Mas tinha um problema, a resposta tava durando mais de um minuto pra chegar.

Em alguns casos já que a resposta normal do servidor demora mais de 4s, por isso resolvi que tinha que dar uma forma de deixar pelo menos um pouco mais rápido.

Então, modifiquei a função Get pra receber um callback e argumentos para tornar as coisas assincronas e tentar executar mais de uma requisição ao mesmo tempo, isso aumentou levemente a velocidade e a complexidade do código (além de não tem muitas boas práticas quando o assunto é thread safe, mas ta funcionando), então pode ter ficado um pouco difícil de ler pra algumas pessoas.

> PS: A API te bloqueia temporariamente depois de um certo número de requests, portanto use com moderação

# Desafio - GO StarWars

Star Wars é uma grande saga de filmes de Ficção Científica que está dentro da cultura pop desde os anos 70 até os dias de hoje.

 desafio consiste em criar rotas que irão receber parâmetro com o do filme e retornar todos os dados do filme e dos os nomes de personagens, naves, planetas, veículos, espécies e suas características. Exemplo:

**Request:**

*http://localhost:8000/A_NEW_HOPE*

**Response:**

```json
{
    "title": "A New Hope",
    "episode_id": 4,
    "opening_crawl": "It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....",
    "director": "George Lucas",
    "producer": "Gary Kurtz, Rick McCallum",
    "release_date": "1977-05-25",
    "characters": [
        {
            "name": "Luke Skywalker",
            "height": "172",
            "mass": "77",
            "hair_color": "blond",
            "skin_color": "fair",
            "eye_color": "blue",
            "birth_year": "19BBY",
            "gender": "male",
            "homeworld": "Tatooine"
        },
        {
            "name": "C-3PO",
            "height": "167",
            "mass": "75",
            "hair_color": "n/a",
            "skin_color": "gold",
            "eye_color": "yellow",
            "birth_year": "112BBY",
            "gender": "n/a",
            "homeworld": "Tatooine"
        }
        ...
    ],
    "planets": [
        {
            "name": "Tatooine",
            "rotation_period": "23",
            "orbital_period": "304",
            "diameter": "10465",
            "climate": "arid",
            "gravity": "1 standard",
            "terrain": "desert",
            "surface_water": "1",
            "population": "200000"
        },
        ...
    ],
    "starships": [
        {
            "name": "CR90 corvette",
            "model": "CR90 corvette",
            "manufacturer": "Corellian Engineering Corporation",
            "cost_in_credits": "3500000",
            "length": "150",
            "max_atmosphering_speed": "950",
            "crew": "30-165",
            "passengers": "600",
            "cargo_capacity": "3000000",
            "consumables": "1 year",
            "hyperdrive_rating": "2.0",
            "MGLT": "60",
            "starship_class": "corvette",
            "pilots": []
        },
        ...
    ],
    "vehicles": [
        {
            "name": "Sand Crawler",
            "model": "Digger Crawler",
            "manufacturer": "Corellia Mining Corporation",
            "cost_in_credits": "150000",
            "length": "36.8 ",
            "max_atmosphering_speed": "30",
            "crew": "46",
            "passengers": "30",
            "cargo_capacity": "50000",
            "consumables": "2 months",
            "vehicle_class": "wheeled",
            "pilots": []
        },
        ...
    ],
    "species": [
        {
            "name": "Human",
            "classification": "mammal",
            "designation": "sentient",
            "average_height": "180",
            "skin_colors": "caucasian, black, asian, hispanic",
            "hair_colors": "blonde, brown, black, red",
            "eye_colors": "brown, blue, green, hazel, grey, amber",
            "average_lifespan": "120",
            "homeworld": "Coruscant",
            "language": "Galactic Basic"

        },
        ...
    ]
}
```

## Requisitos Obrigatórios

- [X] Deve utilizar a api https://swapi.dev/ diretamente pelo **endpoint**
- [X] Deve ser disponibilizado uma rota para receber o parâmetro por **params**
- [X] Toda a estrutura deve ser desenvolvida apenas utilizando a linguagem GO e suas bibliotecas existentes/oficiais, disponíveis em **https://pkg.go.tech/**
- [ ] As requisições não podem passar de 15ms para o retorno
    
    > A API swapi leva mais de 3s para responder sozinha, e a minha faz diversas requisições com base em sublinks oq impossibilita isso:
    > ![image-20210618192035264](/home/plankiton/.config/Typora/typora-user-images/image-20210618192035264.png)
- [X] Pode ser utilizado todo tipo de consulta e exemplos da internet
- [X] O projeto deve estar disponibilizado no GitHub
- [X] **Para a criação do servidor web deverá ser utilizado o SexPistol ([https://github.com/plankiton/SexPistol](https://www.google.com/url?q=https://github.com/plankiton/SexPistol&sa=D&source=editors&ust=1623963347448000&usg=AOvVaw22ahpgc6kY4d3_QZg6yb6j))**
- [X] **Deverá funcionar apenas para os 6 primeiros filmes da Saga Star Wars**

## Requisitos Opcionais

- [x] Criação de um ambiente docker para rodar a aplicação, externalizando a porta para ser acessada via **localhost**
- [ ] Criação de um Insomnia/Postman para teste através das plataformas
- [ ] Criação de um painel web para apresentar os dados ao clicar em botões com os nomes dos filmesG
