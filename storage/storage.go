package storage

import (
	"log"
	"reflect"

	"github.com/fatih/structs"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Insert data in Neo4j
func Insert(stmt bolt.Stmt, data interface{}) error {

	m := structs.Map(data)

	flag := false
	for _, v := range m {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			flag = true
		}
	}
	if flag {
		m = flatMap(m, "")
	}

	// Parses data struct to a map before sending it to ExeNeo()
	result, err := stmt.ExecNeo(m)
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

func flatMap(value map[string]interface{}, prefix string) map[string]interface{} {
	m := make(map[string]interface{})
	flattener(value, prefix, m)
	return m
}

func flattener(value interface{}, prefix string, m map[string]interface{}) {

	base := ""

	if prefix != "" {
		base = prefix + "_"
	}

	original := reflect.ValueOf(value)
	kind := original.Kind()
	t := original.Type()

	switch kind {

	case reflect.Map:
		for _, childKey := range original.MapKeys() {
			childValue := original.MapIndex(childKey)
			flattener(childValue.Interface(), base+childKey.String(), m)
		}

	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			childValue := original.Field(i)
			childKey := t.Field(i).Name
			flattener(childValue.Interface(), base+childKey, m)
		}

	default:
		if prefix != "" {
			m[prefix] = value
		}
	}
}
