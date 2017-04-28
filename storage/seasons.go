package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const seasonQuery = `MERGE (s:Season{id: {ID}}) ON CREATE SET s.name = {Name}, s.active = {Active} ON MATCH SET s.name = {Name}, s.active = {Active} MERGE (c:Competition{id: 													{Competition_ID}}) ON CREATE SET c.cup = {Competition_Cup}, c.name = {Competition_Name}, c.active = {Competition_Active} ON MATCH SET c.cup = {Competition_Cup}, 										c.name = {Competition_Name}, c.active = {Competition_Active} MERGE (c)-[:HAS_SEASON]->(s)`

// SeasonStmt prepare a neo4j statement for single season data
func SeasonStmt(conn bolt.Conn) (bolt.Stmt, error) {
	stmt, err := conn.PrepareNeo(seasonQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create season statement: %s", err)
	}
	return stmt, nil
}
