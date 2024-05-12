/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient
import (
	"time"
)

type SongInfo struct {
	AlbumArtists []string `json:"AlbumArtists,omitempty"`
	Album string `json:"Album,omitempty"`
	Artists []string `json:"Artists,omitempty"`
	Composers []string `json:"Composers,omitempty"`
	// The name.
	Name string `json:"Name,omitempty"`
	// The metadata language.
	MetadataLanguage string `json:"MetadataLanguage,omitempty"`
	// The metadata country code.
	MetadataCountryCode string `json:"MetadataCountryCode,omitempty"`
	MetadataLanguages []GlobalizationCultureDto `json:"MetadataLanguages,omitempty"`
	ProviderIds *map[string]string `json:"ProviderIds,omitempty"`
	// The year.
	Year int32 `json:"Year,omitempty"`
	IndexNumber int32 `json:"IndexNumber,omitempty"`
	ParentIndexNumber int32 `json:"ParentIndexNumber,omitempty"`
	PremiereDate time.Time `json:"PremiereDate,omitempty"`
	IsAutomated bool `json:"IsAutomated,omitempty"`
	EnableAdultMetadata bool `json:"EnableAdultMetadata,omitempty"`
}