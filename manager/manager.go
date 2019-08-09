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

func insertFll(ctx context.Context, client *bigquery.Client) bool {
	fmt.Println("client:", client)	
	u := client.Dataset("ffl").Table("ffl_list").Uploader()	
	fmt.Println(ctx)
	fmt.Println(u)	
	var f Ffl
	f.ID = "2"
	f.Name = "two"
	// if err := u.Put(ctx, f); err != nil {
	if err := u.Put(ctx, f); err != nil {		
		fmt.Println("err:", err)
	}	
	return true
}
