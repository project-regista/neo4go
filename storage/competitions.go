package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// TODO: Create queries
const (
	competitionQuery         = "MERGE (n:NODE {foo: {foo}, bar: {bar}})"
	competitionsQuery        = "MERGE (n:NODE {foo: {foo}, bar: {bar}})"
	competitionsSeasonsQuery = "MERGE (n:NODE {foo: {foo}, bar: {bar}})"
)

// CompetitionStmt prepare a neo4j statement for single competition data
func CompetitionStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(competitionQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create competition statement: %s", err)
	}
	return stmt, nil
}

// CompetitionsStmt prepare a neo4j statement for a list competitions
func CompetitionsStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(competitionsQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create competitions statement: %s", err)
	}
	return stmt, nil
}

// CompetitionsSeasonsStmt prepare a neo4j statement for a list of competitions w/ season data
func CompetitionsSeasonsStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(competitionsSeasonsQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create competitionsSeasons statement: %s", err)
	}
	return stmt, nil
}
