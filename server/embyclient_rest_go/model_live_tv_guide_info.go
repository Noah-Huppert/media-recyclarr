/*
 * Emby Server REST API
 *
 * Explore the Emby Server API
 *
 */
package embyclient_rest_go
import (
	"time"
)

type LiveTvGuideInfo struct {
	// The start date.
	StartDate time.Time `json:"StartDate,omitempty"`
	// The end date.
	EndDate time.Time `json:"EndDate,omitempty"`
}
