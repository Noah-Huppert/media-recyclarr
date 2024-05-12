/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type ExternalUrl struct {
	// The name.
	Name string `json:"Name,omitempty"`
	// The type of the item.
	Url string `json:"Url,omitempty"`
}
