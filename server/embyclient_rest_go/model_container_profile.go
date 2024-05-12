/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type ContainerProfile struct {
	Type_ *DlnaProfileType `json:"Type,omitempty"`
	Conditions []ProfileCondition `json:"Conditions,omitempty"`
	Container string `json:"Container,omitempty"`
}
