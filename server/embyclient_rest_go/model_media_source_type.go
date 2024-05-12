/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type MediaSourceType string

// List of MediaSourceType
const (
	DEFAULT__MediaSourceType MediaSourceType = "Default"
	GROUPING_MediaSourceType MediaSourceType = "Grouping"
	PLACEHOLDER_MediaSourceType MediaSourceType = "Placeholder"
)
