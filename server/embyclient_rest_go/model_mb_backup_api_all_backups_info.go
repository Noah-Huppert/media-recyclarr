/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type MbBackupApiAllBackupsInfo struct {
	FullBackupInfo *MbBackupBackupInfo `json:"FullBackupInfo,omitempty"`
	LightBackups []MbBackupBackupInfo `json:"LightBackups,omitempty"`
}
