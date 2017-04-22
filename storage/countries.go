package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// TODO: Create queries
const (
	countryQuery              = "MERGE (n:NODE {foo: {foo}, bar: {bar}})"
	countriesQuery            = "MERGE (n:NODE {foo: {foo}, bar: {bar}})"
	countriesCompetitionQuery = "MERGE (n:NODE {foo: {foo}, bar: {bar}})"
)

// CountryStmt prepare a neo4j statement for single contry data
func CountryStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(countryQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create country statement: %s", err)
	}
	return stmt, nil
}

// CountriesStmt prepare a neo4j statement for a list countries
func CountriesStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(countriesQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create countries statement: %s", err)
	}
	return stmt, nil
}

// CountriesCompetitionsStmt prepare a neo4j statement for a list of countries w/ competition data
func CountriesCompetitionsStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(countriesCompetitionQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create countriesCompetitions statement: %s", err)
	}
	return stmt, nil
}
