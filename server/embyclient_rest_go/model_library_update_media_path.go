/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type LibraryUpdateMediaPath struct {
	Id string `json:"Id,omitempty"`
	PathInfo *MediaPathInfo `json:"PathInfo,omitempty"`
}
