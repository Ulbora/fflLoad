package manager

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

//Ffl ffl
type Ffl struct {
	ID   string `bigquery:"id"`
	Name string `bigquery:"name"`
}

func insertFll(client *bigquery.Client) bool {
	fmt.Println("client:", client)
	ctx := context.Background()
	// [START bigquery_table_insert_rows]
	//u := client.Dataset("august-gantry-192521:ffl").Table("ffl_list").Uploader()
	//u := client.Dataset("ffl").Table("ffl_list").Uploader()
	// items := []*Item{
	// 	// Item implements the ValueSaver interface.
	// 	{Name: "Phred Phlyntstone", Age: 32},
	// 	{Name: "Wylma Phlyntstone", Age: 29},
	fmt.Println(ctx)
	//fmt.Println(u)
	// }
	var f Ffl
	f.ID = "2"
	f.Name = "two"
	// if err := u.Put(ctx, f); err != nil {
	// if err := u.Put(ctx, f); err != nil {
	// 	//return err
	// 	fmt.Println("err:", err)
	// }
	// [END bigquery_table_insert_rows]
	return true
}
