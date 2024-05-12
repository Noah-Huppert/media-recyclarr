/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type LibraryRenameVirtualFolder struct {
	Id string `json:"Id,omitempty"`
	NewName string `json:"NewName,omitempty"`
}
