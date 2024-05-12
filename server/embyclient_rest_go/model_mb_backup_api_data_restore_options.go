/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type MbBackupApiDataRestoreOptions struct {
	Users []MbBackupApiUserRestoreInfo `json:"Users,omitempty"`
}
