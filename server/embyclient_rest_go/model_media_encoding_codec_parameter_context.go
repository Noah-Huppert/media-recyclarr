/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type MediaEncodingCodecParameterContext string

// List of MediaEncoding.CodecParameterContext
const (
	PLAYBACK_MediaEncodingCodecParameterContext MediaEncodingCodecParameterContext = "Playback"
	CONVERSION_MediaEncodingCodecParameterContext MediaEncodingCodecParameterContext = "Conversion"
)
