package storage

import (
	"log"

	"fmt"

	"github.com/fatih/structs"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Insert data in Neo4j
func Insert(stmt bolt.Stmt, data interface{}) error {

	// Parses data struct to a map before sending it to ExeNeo()
	result, err := stmt.ExecNeo(structs.Map(data))
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

// Insert a Pipeline data in Neo4j
func InsertPipeline(pipeStmt bolt.PipelineStmt, data interface{}, conn bolt.Conn) error {

	fmt.Println(pipeStmt)

	// Prepare Pipeline with queries to execute
	// pipeline, err := conn.PreparePipeline(pipeStmt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Execute queries
	// _, err = pipeline.ExecPipeline(nilSlice(pipeStmt)...)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Close pipeline
	// if err = pipeline.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	return nil
}

// Return a slice of nil values with queries length.
// Workaround for creating n nil params for ExecPipeline method
func nilSlice(s []string) []map[string]interface{} {
	tmp := make([]map[string]interface{}, len(s), len(s))
	for i := range tmp {
		tmp[i] = nil
	}
	return tmp
}
