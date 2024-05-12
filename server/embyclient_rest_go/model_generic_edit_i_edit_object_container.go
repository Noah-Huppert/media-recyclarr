/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type GenericEditIEditObjectContainer struct {
	Object *interface{} `json:"Object,omitempty"`
	DefaultObject *interface{} `json:"DefaultObject,omitempty"`
	TypeName string `json:"TypeName,omitempty"`
}
