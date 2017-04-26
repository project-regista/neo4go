package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// TODO: Create queries
const (
	competitionQuery = `MERGE (c:Competition{id: {ID} }) 
								ON CREATE SET c.cup = {Cup}, c.name = {Name}, c.active = {Active}
								ON MATCH SET c.cup = {Cup}, c.name = {Name}, c.active = {Active}`

	competitionsQuery = `MERGE (c:Competition{id: {ID} }) 
								ON CREATE SET c.cup = {Cup}, c.name = {Name}, c.active = {Active}
								ON MATCH SET c.cup = {Cup}, c.name = {Name}, c.active = {Active}`

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
func CompetitionsStmt(conn bolt.Conn, n int) (bolt.PipelineStmt, error) {
	s := make([]string, 0, 0)
	for i := 0; i < n; i++ {
		s = append(s, competitionsQuery)
	}

	pipeStmt, err := conn.PreparePipeline(s...)
	if err != nil {
		return nil, fmt.Errorf("Failed to create competitions statement: %s", err)
	}
	return pipeStmt, nil
}

// CompetitionsSeasonsStmt prepare a neo4j statement for a list of competitions w/ season data
func CompetitionsSeasonsStmt(conn bolt.Conn) (bolt.Stmt, error) {

	stmt, err := conn.PrepareNeo(competitionsSeasonsQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create competitionsSeasons statement: %s", err)
	}
	return stmt, nil
}
