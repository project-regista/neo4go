package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const matchQuery = "foo"

// MatchStmt prepare a Neo4j statement for a single match
func MatchStmt(conn bolt.Conn) (bolt.Stmt, error) {
	stmt, err := conn.PrepareNeo(matchQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create season statement: %s", err)
	}
	return stmt, nil
}
