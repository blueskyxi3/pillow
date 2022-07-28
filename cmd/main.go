package main

import (
	"context"
	"fmt"
	"github.com/blueskyxi3/pillow/pkg/pillow"
	"log"
)

var db *pillow.DB

var rawJSON = `
{
   "selector": {
      "orderNo": {
         "$regex": "pr2022*"
      }
   }
}
`

var rawJSON2 = `
		{
		    "selector": {
					"$and": [
		        	{
		            "port": { "$gt": 15 }
		        	},
		        	{
		            "dbname": {
		              "$in": ["cdrprd"]
		          	}
		        	}
		    	]
				}
		}`

var raw1JSON = `
		{
		    "selector": {
					"$and": [
		        	{
		            "_id": { "$gt": null }
		        	},
		        	{
		            "year": {
		              "$in": [2007, 2004]
		          	}
		        	}
		    	]
				}
		}`

func main() {
	//createDoc()
	//searchDoc()

	searchByJSON(rawJSON)
}

func searchByJSON(query string) {
	output, err := db.QueryWithJSON(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	rows, _ := output["docs"].([]interface{})
	for _, row := range rows {
		log.Printf(" [===>] Record: %v\n", row)
		if rec, ok := row.(map[string]interface{}); ok {
			for key, val := range rec {
				log.Printf(" [========>] %v = %v", key, val)
			}
		} else {
			fmt.Printf("record not a map[string]interface{}: %v\n", row)
		}
	}
	// log.Printf("%v\n", output)
}
func searchDoc() {
	resMap, err := db.ListDocuments(context.TODO())
	if err != nil {
		panic(err)
	}
	totalRows, _ := resMap["total_rows"].(float64)
	if totalRows > 0 {
		rows, _ := resMap["rows"].([]interface{})
		for _, row := range rows {
			log.Printf(" [===>] Record: %v\n", row)
			if rec, ok := row.(map[string]interface{}); ok {
				for key, val := range rec {
					log.Printf(" [========>] %s = %s", key, val)
				}
			} else {
				fmt.Printf("record not a map[string]interface{}: %v\n", row)
			}
		}
	}
}
func init() {
	log.Println("---init db---")
	dsn := "http://admin:admin@localhost:5984/"
	client, err := pillow.New(dsn)
	if err != nil {
		panic(err)
	}
	db = client.Database(context.TODO(), "order")
}

func createDoc() {
	document := map[string]interface{}{
		"orderNo":    "tenants:john-doe",
		"first_name": "John",
		"last_name":  "Doe",
	}

	_, err := db.CreateDocument(context.TODO(), document)
	if err != nil {
		panic(err)
	}

	log.Printf("doc %v created \n", document)
}
