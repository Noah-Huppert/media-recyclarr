/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type SyncedItemProgress struct {
	Progress float64 `json:"Progress,omitempty"`
	Status *SyncJobItemStatus `json:"Status,omitempty"`
}
