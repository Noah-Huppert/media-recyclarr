/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

// A class combining both `Emby.Media.Model.Types.ProfileInformation` and `Emby.Media.Model.Types.LevelInformation`.  
type ProfileLevelInformation struct {
	Profile *ProfileInformation `json:"Profile,omitempty"`
	Level *LevelInformation `json:"Level,omitempty"`
}
