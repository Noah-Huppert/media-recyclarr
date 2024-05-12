/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go
// PackageVersionClass : Enum PackageVersionClass  
type PackageVersionClass string

// List of PackageVersionClass
const (
	RELEASE_PackageVersionClass PackageVersionClass = "Release"
	BETA_PackageVersionClass PackageVersionClass = "Beta"
	DEV_PackageVersionClass PackageVersionClass = "Dev"
)
