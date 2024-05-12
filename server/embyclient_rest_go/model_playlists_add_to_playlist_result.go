/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type PlaylistsAddToPlaylistResult struct {
	Id string `json:"Id,omitempty"`
	ItemAddedCount int32 `json:"ItemAddedCount,omitempty"`
}
