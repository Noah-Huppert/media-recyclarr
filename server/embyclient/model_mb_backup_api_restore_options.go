/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient

type MbBackupApiRestoreOptions struct {
	RestoreServerId bool `json:"RestoreServerId,omitempty"`
	UseFiles string `json:"UseFiles,omitempty"`
}