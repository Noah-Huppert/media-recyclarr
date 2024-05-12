/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient

type ApiListingProviderTypeInfo struct {
	Name string `json:"Name,omitempty"`
	Id string `json:"Id,omitempty"`
	SetupUrl string `json:"SetupUrl,omitempty"`
}