package graph

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/RoongJin/pokedex-graphql-sqlite/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	*sql.DB
}

type Resolver struct {
	DB Database
}

func (db Database) AddPokemon(name string, description string, category string, typeOf string, abilities string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO Pokemons(Name, Description, Category, Type, Abilities) values(?,?,?,?,?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(name, description, category, typeOf, abilities)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	fmt.Println(id)

	defer stmt.Close()
	return id, nil
}

func (db Database) UpdatePokemon(id int, name string, description string, category string, typeOf string, abilities string) (int64, error) {
	stmt, err := db.Prepare("update Pokemons set Name=?, Description=?, Category=?, Type=?, Abilities=? where ID=?")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(name, description, category, typeOf, abilities, id)
	if err != nil {
		return -1, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	fmt.Println(affected)

	defer stmt.Close()
	return affected, nil
}

func (db Database) DeletePokemon(id int) (int64, error) {
	stmt, err := db.Prepare("delete from Pokemons where ID=?")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return -1, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	fmt.Println(affected)

	defer stmt.Close()
	return affected, nil
}

func (db Database) GetAllPokemons() ([]*model.Pokemon, error) {
	rows, err := db.Query("select * from Pokemons")
	if err != nil {
		return nil, err
	}

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

func (db Database) FindPokemonById(id int64) (model.Pokemon, error) {
	rows, err := db.Query("select * from Pokemons where ID=?", id)
	if err != nil {
		return model.Pokemon{}, err
	}
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

	if name == "" {
		return model.Pokemon{}, fmt.Errorf("Pokemon with this ID does not exist!")
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
