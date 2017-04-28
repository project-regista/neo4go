package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const countryQuery = `MERGE (c:Country{id: {ID} })
											ON CREATE SET c.name = {Name}, c.isoCode = {IsoCode}, c.flag = {Flag}
											ON MATCH SET c.name = {Name}, c.isoCode = {IsoCode}, c.flag = {Flag}`

// CountryStmt prepare a neo4j statement for single contry data
func CountryStmt(conn bolt.Conn) (bolt.Stmt, error) {
	stmt, err := conn.PrepareNeo(countryQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create country statement: %s", err)
	}
	return stmt, nil
}
