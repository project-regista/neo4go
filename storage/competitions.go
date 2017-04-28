package storage

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const competitionQuery = `MERGE (c:Competition{id: {ID} })
							ON CREATE SET c.cup = {Cup}, c.name = {Name}, c.active = {Active}
							ON MATCH SET c.cup = {Cup}, c.name = {Name}, c.active = {Active}
							MERGE (cc:Country{id: {Country_ID} })
							ON CREATE SET cc.name = {Country_Name}, cc.isoCode = {Country_IsoCode}, cc.flag = {Country_Flag}
							ON MATCH SET cc.name = {Country_Name}, cc.isoCode = {Country_IsoCode}, cc.flag = {Country_Flag}
							MERGE (cc)-[:ORGANISES]->(c)`

// CompetitionStmt prepare a neo4j statement for single competition data
func CompetitionStmt(conn bolt.Conn) (bolt.Stmt, error) {
	stmt, err := conn.PrepareNeo(competitionQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to create competition statement: %s", err)
	}
	return stmt, nil
}
