package main

import (
	"context"

	"github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	client := embyclient.NewAPIClient(&embyclient.Configuration{})
	client.ItemsServiceApi.GetItems(context.Background(), &embyclient.ItemsServiceApiGetItemsOpts{})
}
