/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type NameLongIdPair struct {
	// The name.
	Name string `json:"Name,omitempty"`
	// The identifier.
	Id int64 `json:"Id,omitempty"`
}
