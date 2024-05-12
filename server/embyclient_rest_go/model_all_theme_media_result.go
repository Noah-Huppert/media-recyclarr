/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type AllThemeMediaResult struct {
	ThemeVideosResult *ThemeMediaResult `json:"ThemeVideosResult,omitempty"`
	ThemeSongsResult *ThemeMediaResult `json:"ThemeSongsResult,omitempty"`
	SoundtrackSongsResult *ThemeMediaResult `json:"SoundtrackSongsResult,omitempty"`
}
