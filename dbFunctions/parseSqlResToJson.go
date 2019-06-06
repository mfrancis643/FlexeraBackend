package dbFunctions

import (
	"database/sql"
	"encoding/json"
	"log"
)

func parseSqlResToJson(rows *sql.Rows) string  {

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	//Make table skeleton
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	//Work through empty table cells and populate from "rows" variable
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		//Add new row to data
		tableData = append(tableData, entry)
	}

	//Convert map format to Json format
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(jsonData), nil)
	return string(jsonData)
}