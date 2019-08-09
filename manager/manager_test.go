package manager

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/bigquery"
)

func Test_insertFll(t *testing.T) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "august-gantry-192521")
	fmt.Println("error: ", err)
	insertFll(client)
}
