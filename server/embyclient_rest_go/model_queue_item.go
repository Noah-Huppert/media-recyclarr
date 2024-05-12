/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type QueueItem struct {
	Id int64 `json:"Id,omitempty"`
	PlaylistItemId string `json:"PlaylistItemId,omitempty"`
}
