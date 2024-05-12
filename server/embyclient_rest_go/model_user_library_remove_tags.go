/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type UserLibraryRemoveTags struct {
	Tags []NameIdPair `json:"Tags,omitempty"`
}
