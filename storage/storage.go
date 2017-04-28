package storage

import (
	"log"

	"github.com/fatih/structs"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Insert data in Neo4j
func Insert(stmt bolt.Stmt, data interface{}) error {

	// Parses data struct to a map before sending it to ExeNeo()
	if err != nil {
		return err
	}

	numResult, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("CREATED ROWS: %d\n", numResult)
	return nil
}







	}
}
