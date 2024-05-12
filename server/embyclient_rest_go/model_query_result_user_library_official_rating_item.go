/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type QueryResultUserLibraryOfficialRatingItem struct {
	Items []UserLibraryOfficialRatingItem `json:"Items,omitempty"`
	TotalRecordCount int32 `json:"TotalRecordCount,omitempty"`
}
