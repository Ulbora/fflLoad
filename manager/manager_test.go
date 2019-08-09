package manager

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func Test_insertFll(t *testing.T) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "august-gantry-192521", option.WithCredentialsFile("/home/ken/goworkspace/My First Project-682f1b5fa0de.json"))
	fmt.Println("error in test: ", err)
	insertFll(ctx, client)
}
