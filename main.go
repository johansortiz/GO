// main.go

package main

import (
	"log"
	"net/http"

	. "github.com/johansortiz/SuperHero/handlers"
	. "github.com/johansortiz/SuperHero/models"
	. "github.com/johansortiz/SuperHero/repository"
)

func main() {
	repo := NewSuperheroRepository()

	// Agregar los superhéroes al repositorio
	repo.Add(Superhero{
		Name: "Spider-Man",
		Biography: Biography{
			FullName: "Peter Parker",
		},
		Powerstats: Powerstats{
			Intelligence: 90,
			Strength:     55,
			Speed:        67,
			Durability:   75,
			Power:        74,
			Combat:       85,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
		},
	})
	repo.Add(Superhero{
		Name: "Iron Man",
		Biography: Biography{
			FullName: "Tony Stark",
		},
		Powerstats: Powerstats{
			Intelligence: 100,
			Strength:     85,
			Speed:        58,
			Durability:   85,
			Power:        100,
			Combat:       64,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
		},
	})
	repo.Add(Superhero{
		Name: "Black Widow",
		Biography: Biography{
			FullName: "Natasha Romanoff",
		},
		Powerstats: Powerstats{
			Intelligence: 75,
			Strength:     13,
			Speed:        33,
			Durability:   30,
			Power:        36,
			Combat:       100,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
		},
	})
	repo.Add(Superhero{
		Name: "Thor",
		Biography: Biography{
			FullName: "Thor Odinson",
		},
		Powerstats: Powerstats{
			Intelligence: 69,
			Strength:     100,
			Speed:        83,
			Durability:   100,
			Power:        100,
			Combat:       100,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
		},
	})
	repo.Add(Superhero{
		Name: "Wolverine",
		Biography: Biography{
			FullName: "John Logan",
		},
		Powerstats: Powerstats{
			Intelligence: 63,
			Strength:     32,
			Speed:        50,
			Durability:   100,
			Power:        89,
			Combat:       100,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
		},
	})

	handler := NewSuperheroHandler(repo)

	http.Handle("/api/superhero", handler)

	log.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
