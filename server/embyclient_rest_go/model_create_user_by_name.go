/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go

type CreateUserByName struct {
	Name string `json:"Name,omitempty"`
	CopyFromUserId string `json:"CopyFromUserId,omitempty"`
	UserCopyOptions []LibraryUserCopyOptions `json:"UserCopyOptions,omitempty"`
}
