/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type EncodingContext string

// List of EncodingContext
const (
	STREAMING_EncodingContext EncodingContext = "Streaming"
	STATIC_EncodingContext EncodingContext = "Static"
)
