package graph

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/RoongJin/pokedex-graphql-sqlite/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	DBPointer *sql.DB
}

type Resolver struct {
	DB Database
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AddPokemon(db *sql.DB, name string, description string, category string, typeOf string, abilities string) int64 {
	stmt, err := db.Prepare("INSERT INTO Pokemons(Name, Description, Category, Type, Abilities) values(?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(name, description, category, typeOf, abilities)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	defer stmt.Close()
	return id
}

func UpdatePokemon(db *sql.DB, id int, name string, description string, category string, typeOf string, abilities string) int64 {
	stmt, err := db.Prepare("update Pokemons set Name=?, Description=?, Category=?, Type=?, Abilities=? where ID=?")
	checkErr(err)

	res, err := stmt.Exec(name, description, category, typeOf, abilities, id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)

	defer stmt.Close()
	return affected
}

func DeletePokemon(db *sql.DB, id int) int64 {
	stmt, err := db.Prepare("delete from Pokemons where ID=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)

	defer stmt.Close()
	return affected
}

func GetAllPokemons(db *sql.DB) ([]*model.Pokemon, error) {
	rows, err := db.Query("select * from Pokemons")
	checkErr(err)

	var pokeList []*model.Pokemon
	for rows.Next() {
		var name string
		var desc string
		var category string
		var types string
		var abilities string
		var dummy string
		err = rows.Scan(&name, &desc, &category, &types, &abilities, &dummy)
		fmt.Println("Name: " + name)
		fmt.Println("Description: " + desc)
		fmt.Println("Category: " + category)
		fmt.Println("Type: " + types)
		fmt.Println("Abilities: " + abilities)

		t := strings.Split(types, " ")
		a := strings.Split(abilities, " ")

		poke := model.Pokemon{
			ID:          dummy,
			Name:        name,
			Description: desc,
			Category:    category,
			Type:        t,
			Abilities:   a,
		}
		pokeList = append(pokeList, &poke)
	}

	return pokeList, nil
}

func FindPokemonById(db *sql.DB, id int64) (model.Pokemon, error) {
	rows, err := db.Query("select * from Pokemons where ID=?", id)
	checkErr(err)
	var name string
	var desc string
	var category string
	var types string
	var abilities string
	var dummy string

	for rows.Next() {
		err = rows.Scan(&name, &desc, &category, &types, &abilities, &dummy)
		fmt.Println("Name: " + name)
		fmt.Println("Description: " + desc)
		fmt.Println("Category: " + category)
		fmt.Println("Type: " + types)
		fmt.Println("Abilities: " + abilities)
	}

	dm := model.Pokemon{}
	if name == "" {
		return dm, fmt.Errorf("Pokemon with this ID does not exist!")
	}

	t := strings.Split(types, " ")
	a := strings.Split(abilities, " ")

	poke := model.Pokemon{
		ID:          dummy,
		Name:        name,
		Description: desc,
		Category:    category,
		Type:        t,
		Abilities:   a,
	}

	defer rows.Close()
	return poke, nil
}
