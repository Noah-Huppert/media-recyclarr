/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type GeneralCommand struct {
	Name string `json:"Name,omitempty"`
	ControllingUserId string `json:"ControllingUserId,omitempty"`
	Arguments map[string]string `json:"Arguments,omitempty"`
}
