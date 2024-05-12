/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type ImageSavingConvention string

// List of ImageSavingConvention
const (
	LEGACY_ImageSavingConvention ImageSavingConvention = "Legacy"
	COMPATIBLE_ImageSavingConvention ImageSavingConvention = "Compatible"
)
