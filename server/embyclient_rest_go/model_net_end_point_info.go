/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type NetEndPointInfo struct {
	IsLocal bool `json:"IsLocal,omitempty"`
	IsInNetwork bool `json:"IsInNetwork,omitempty"`
}
