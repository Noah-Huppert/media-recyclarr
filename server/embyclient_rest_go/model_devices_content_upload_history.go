/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type DevicesContentUploadHistory struct {
	DeviceId string `json:"DeviceId,omitempty"`
	FilesUploaded []DevicesLocalFileInfo `json:"FilesUploaded,omitempty"`
}
