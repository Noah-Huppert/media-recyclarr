/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type MbBackupApiRestoreOptions struct {
	RestoreServerId bool `json:"RestoreServerId,omitempty"`
	UseFiles string `json:"UseFiles,omitempty"`
}
