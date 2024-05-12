/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type CodecType string

// List of CodecType
const (
	VIDEO_CodecType CodecType = "Video"
	VIDEO_AUDIO_CodecType CodecType = "VideoAudio"
	AUDIO_CodecType CodecType = "Audio"
)
