package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RoongJin/pokedex-graphql-sqlite/graph"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()

	MyDB, err := sql.Open("sqlite3", "./Pokedex.db")
	if err != nil {
		log.Fatal(err)
	}

	defer MyDB.Close()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: graph.Database{DBPointer: MyDB},
	}}))

	//DeletePokemon(db, 1, "Bulbasaur", "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.", "Seed", "Grass Poison", "Overgrow")
	//FindPokemonById(db, 1)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
