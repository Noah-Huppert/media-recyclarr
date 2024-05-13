/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyseerrclient

import (
	"encoding/json"
)

// checks if the SonarrSeries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SonarrSeries{}

// SonarrSeries struct for SonarrSeries
type SonarrSeries struct {
	Title *string `json:"title,omitempty"`
	SortTitle *string `json:"sortTitle,omitempty"`
	SeasonCount *float32 `json:"seasonCount,omitempty"`
	Status *string `json:"status,omitempty"`
	Overview *string `json:"overview,omitempty"`
	Network *string `json:"network,omitempty"`
	AirTime *string `json:"airTime,omitempty"`
	Images []SonarrSeriesImagesInner `json:"images,omitempty"`
	RemotePoster *string `json:"remotePoster,omitempty"`
	Seasons []SonarrSeriesSeasonsInner `json:"seasons,omitempty"`
	Year *float32 `json:"year,omitempty"`
	Path *string `json:"path,omitempty"`
	ProfileId *float32 `json:"profileId,omitempty"`
	LanguageProfileId *float32 `json:"languageProfileId,omitempty"`
	SeasonFolder *bool `json:"seasonFolder,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	UseSceneNumbering *bool `json:"useSceneNumbering,omitempty"`
	Runtime *float32 `json:"runtime,omitempty"`
	TvdbId *float32 `json:"tvdbId,omitempty"`
	TvRageId *float32 `json:"tvRageId,omitempty"`
	TvMazeId *float32 `json:"tvMazeId,omitempty"`
	FirstAired *string `json:"firstAired,omitempty"`
	LastInfoSync NullableString `json:"lastInfoSync,omitempty"`
	SeriesType *string `json:"seriesType,omitempty"`
	CleanTitle *string `json:"cleanTitle,omitempty"`
	ImdbId *string `json:"imdbId,omitempty"`
	TitleSlug *string `json:"titleSlug,omitempty"`
	Certification *string `json:"certification,omitempty"`
	Genres []string `json:"genres,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Added *string `json:"added,omitempty"`
	Ratings []SonarrSeriesRatingsInner `json:"ratings,omitempty"`
	QualityProfileId *float32 `json:"qualityProfileId,omitempty"`
	Id NullableFloat32 `json:"id,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	AddOptions []SonarrSeriesAddOptionsInner `json:"addOptions,omitempty"`
}

// NewSonarrSeries instantiates a new SonarrSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarrSeries() *SonarrSeries {
	this := SonarrSeries{}
	return &this
}

// NewSonarrSeriesWithDefaults instantiates a new SonarrSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarrSeriesWithDefaults() *SonarrSeries {
	this := SonarrSeries{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SonarrSeries) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SonarrSeries) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SonarrSeries) SetTitle(v string) {
	o.Title = &v
}

// GetSortTitle returns the SortTitle field value if set, zero value otherwise.
func (o *SonarrSeries) GetSortTitle() string {
	if o == nil || IsNil(o.SortTitle) {
		var ret string
		return ret
	}
	return *o.SortTitle
}

// GetSortTitleOk returns a tuple with the SortTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetSortTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SortTitle) {
		return nil, false
	}
	return o.SortTitle, true
}

// HasSortTitle returns a boolean if a field has been set.
func (o *SonarrSeries) HasSortTitle() bool {
	if o != nil && !IsNil(o.SortTitle) {
		return true
	}

	return false
}

// SetSortTitle gets a reference to the given string and assigns it to the SortTitle field.
func (o *SonarrSeries) SetSortTitle(v string) {
	o.SortTitle = &v
}

// GetSeasonCount returns the SeasonCount field value if set, zero value otherwise.
func (o *SonarrSeries) GetSeasonCount() float32 {
	if o == nil || IsNil(o.SeasonCount) {
		var ret float32
		return ret
	}
	return *o.SeasonCount
}

// GetSeasonCountOk returns a tuple with the SeasonCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetSeasonCountOk() (*float32, bool) {
	if o == nil || IsNil(o.SeasonCount) {
		return nil, false
	}
	return o.SeasonCount, true
}

// HasSeasonCount returns a boolean if a field has been set.
func (o *SonarrSeries) HasSeasonCount() bool {
	if o != nil && !IsNil(o.SeasonCount) {
		return true
	}

	return false
}

// SetSeasonCount gets a reference to the given float32 and assigns it to the SeasonCount field.
func (o *SonarrSeries) SetSeasonCount(v float32) {
	o.SeasonCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SonarrSeries) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SonarrSeries) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SonarrSeries) SetStatus(v string) {
	o.Status = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *SonarrSeries) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *SonarrSeries) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *SonarrSeries) SetOverview(v string) {
	o.Overview = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *SonarrSeries) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *SonarrSeries) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *SonarrSeries) SetNetwork(v string) {
	o.Network = &v
}

// GetAirTime returns the AirTime field value if set, zero value otherwise.
func (o *SonarrSeries) GetAirTime() string {
	if o == nil || IsNil(o.AirTime) {
		var ret string
		return ret
	}
	return *o.AirTime
}

// GetAirTimeOk returns a tuple with the AirTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetAirTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AirTime) {
		return nil, false
	}
	return o.AirTime, true
}

// HasAirTime returns a boolean if a field has been set.
func (o *SonarrSeries) HasAirTime() bool {
	if o != nil && !IsNil(o.AirTime) {
		return true
	}

	return false
}

// SetAirTime gets a reference to the given string and assigns it to the AirTime field.
func (o *SonarrSeries) SetAirTime(v string) {
	o.AirTime = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *SonarrSeries) GetImages() []SonarrSeriesImagesInner {
	if o == nil || IsNil(o.Images) {
		var ret []SonarrSeriesImagesInner
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetImagesOk() ([]SonarrSeriesImagesInner, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *SonarrSeries) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []SonarrSeriesImagesInner and assigns it to the Images field.
func (o *SonarrSeries) SetImages(v []SonarrSeriesImagesInner) {
	o.Images = v
}

// GetRemotePoster returns the RemotePoster field value if set, zero value otherwise.
func (o *SonarrSeries) GetRemotePoster() string {
	if o == nil || IsNil(o.RemotePoster) {
		var ret string
		return ret
	}
	return *o.RemotePoster
}

// GetRemotePosterOk returns a tuple with the RemotePoster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetRemotePosterOk() (*string, bool) {
	if o == nil || IsNil(o.RemotePoster) {
		return nil, false
	}
	return o.RemotePoster, true
}

// HasRemotePoster returns a boolean if a field has been set.
func (o *SonarrSeries) HasRemotePoster() bool {
	if o != nil && !IsNil(o.RemotePoster) {
		return true
	}

	return false
}

// SetRemotePoster gets a reference to the given string and assigns it to the RemotePoster field.
func (o *SonarrSeries) SetRemotePoster(v string) {
	o.RemotePoster = &v
}

// GetSeasons returns the Seasons field value if set, zero value otherwise.
func (o *SonarrSeries) GetSeasons() []SonarrSeriesSeasonsInner {
	if o == nil || IsNil(o.Seasons) {
		var ret []SonarrSeriesSeasonsInner
		return ret
	}
	return o.Seasons
}

// GetSeasonsOk returns a tuple with the Seasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetSeasonsOk() ([]SonarrSeriesSeasonsInner, bool) {
	if o == nil || IsNil(o.Seasons) {
		return nil, false
	}
	return o.Seasons, true
}

// HasSeasons returns a boolean if a field has been set.
func (o *SonarrSeries) HasSeasons() bool {
	if o != nil && !IsNil(o.Seasons) {
		return true
	}

	return false
}

// SetSeasons gets a reference to the given []SonarrSeriesSeasonsInner and assigns it to the Seasons field.
func (o *SonarrSeries) SetSeasons(v []SonarrSeriesSeasonsInner) {
	o.Seasons = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *SonarrSeries) GetYear() float32 {
	if o == nil || IsNil(o.Year) {
		var ret float32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetYearOk() (*float32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *SonarrSeries) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given float32 and assigns it to the Year field.
func (o *SonarrSeries) SetYear(v float32) {
	o.Year = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SonarrSeries) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SonarrSeries) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SonarrSeries) SetPath(v string) {
	o.Path = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *SonarrSeries) GetProfileId() float32 {
	if o == nil || IsNil(o.ProfileId) {
		var ret float32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *SonarrSeries) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given float32 and assigns it to the ProfileId field.
func (o *SonarrSeries) SetProfileId(v float32) {
	o.ProfileId = &v
}

// GetLanguageProfileId returns the LanguageProfileId field value if set, zero value otherwise.
func (o *SonarrSeries) GetLanguageProfileId() float32 {
	if o == nil || IsNil(o.LanguageProfileId) {
		var ret float32
		return ret
	}
	return *o.LanguageProfileId
}

// GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetLanguageProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LanguageProfileId) {
		return nil, false
	}
	return o.LanguageProfileId, true
}

// HasLanguageProfileId returns a boolean if a field has been set.
func (o *SonarrSeries) HasLanguageProfileId() bool {
	if o != nil && !IsNil(o.LanguageProfileId) {
		return true
	}

	return false
}

// SetLanguageProfileId gets a reference to the given float32 and assigns it to the LanguageProfileId field.
func (o *SonarrSeries) SetLanguageProfileId(v float32) {
	o.LanguageProfileId = &v
}

// GetSeasonFolder returns the SeasonFolder field value if set, zero value otherwise.
func (o *SonarrSeries) GetSeasonFolder() bool {
	if o == nil || IsNil(o.SeasonFolder) {
		var ret bool
		return ret
	}
	return *o.SeasonFolder
}

// GetSeasonFolderOk returns a tuple with the SeasonFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetSeasonFolderOk() (*bool, bool) {
	if o == nil || IsNil(o.SeasonFolder) {
		return nil, false
	}
	return o.SeasonFolder, true
}

// HasSeasonFolder returns a boolean if a field has been set.
func (o *SonarrSeries) HasSeasonFolder() bool {
	if o != nil && !IsNil(o.SeasonFolder) {
		return true
	}

	return false
}

// SetSeasonFolder gets a reference to the given bool and assigns it to the SeasonFolder field.
func (o *SonarrSeries) SetSeasonFolder(v bool) {
	o.SeasonFolder = &v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *SonarrSeries) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *SonarrSeries) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *SonarrSeries) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetUseSceneNumbering returns the UseSceneNumbering field value if set, zero value otherwise.
func (o *SonarrSeries) GetUseSceneNumbering() bool {
	if o == nil || IsNil(o.UseSceneNumbering) {
		var ret bool
		return ret
	}
	return *o.UseSceneNumbering
}

// GetUseSceneNumberingOk returns a tuple with the UseSceneNumbering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetUseSceneNumberingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSceneNumbering) {
		return nil, false
	}
	return o.UseSceneNumbering, true
}

// HasUseSceneNumbering returns a boolean if a field has been set.
func (o *SonarrSeries) HasUseSceneNumbering() bool {
	if o != nil && !IsNil(o.UseSceneNumbering) {
		return true
	}

	return false
}

// SetUseSceneNumbering gets a reference to the given bool and assigns it to the UseSceneNumbering field.
func (o *SonarrSeries) SetUseSceneNumbering(v bool) {
	o.UseSceneNumbering = &v
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *SonarrSeries) GetRuntime() float32 {
	if o == nil || IsNil(o.Runtime) {
		var ret float32
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetRuntimeOk() (*float32, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *SonarrSeries) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given float32 and assigns it to the Runtime field.
func (o *SonarrSeries) SetRuntime(v float32) {
	o.Runtime = &v
}

// GetTvdbId returns the TvdbId field value if set, zero value otherwise.
func (o *SonarrSeries) GetTvdbId() float32 {
	if o == nil || IsNil(o.TvdbId) {
		var ret float32
		return ret
	}
	return *o.TvdbId
}

// GetTvdbIdOk returns a tuple with the TvdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetTvdbIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TvdbId) {
		return nil, false
	}
	return o.TvdbId, true
}

// HasTvdbId returns a boolean if a field has been set.
func (o *SonarrSeries) HasTvdbId() bool {
	if o != nil && !IsNil(o.TvdbId) {
		return true
	}

	return false
}

// SetTvdbId gets a reference to the given float32 and assigns it to the TvdbId field.
func (o *SonarrSeries) SetTvdbId(v float32) {
	o.TvdbId = &v
}

// GetTvRageId returns the TvRageId field value if set, zero value otherwise.
func (o *SonarrSeries) GetTvRageId() float32 {
	if o == nil || IsNil(o.TvRageId) {
		var ret float32
		return ret
	}
	return *o.TvRageId
}

// GetTvRageIdOk returns a tuple with the TvRageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetTvRageIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TvRageId) {
		return nil, false
	}
	return o.TvRageId, true
}

// HasTvRageId returns a boolean if a field has been set.
func (o *SonarrSeries) HasTvRageId() bool {
	if o != nil && !IsNil(o.TvRageId) {
		return true
	}

	return false
}

// SetTvRageId gets a reference to the given float32 and assigns it to the TvRageId field.
func (o *SonarrSeries) SetTvRageId(v float32) {
	o.TvRageId = &v
}

// GetTvMazeId returns the TvMazeId field value if set, zero value otherwise.
func (o *SonarrSeries) GetTvMazeId() float32 {
	if o == nil || IsNil(o.TvMazeId) {
		var ret float32
		return ret
	}
	return *o.TvMazeId
}

// GetTvMazeIdOk returns a tuple with the TvMazeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetTvMazeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TvMazeId) {
		return nil, false
	}
	return o.TvMazeId, true
}

// HasTvMazeId returns a boolean if a field has been set.
func (o *SonarrSeries) HasTvMazeId() bool {
	if o != nil && !IsNil(o.TvMazeId) {
		return true
	}

	return false
}

// SetTvMazeId gets a reference to the given float32 and assigns it to the TvMazeId field.
func (o *SonarrSeries) SetTvMazeId(v float32) {
	o.TvMazeId = &v
}

// GetFirstAired returns the FirstAired field value if set, zero value otherwise.
func (o *SonarrSeries) GetFirstAired() string {
	if o == nil || IsNil(o.FirstAired) {
		var ret string
		return ret
	}
	return *o.FirstAired
}

// GetFirstAiredOk returns a tuple with the FirstAired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetFirstAiredOk() (*string, bool) {
	if o == nil || IsNil(o.FirstAired) {
		return nil, false
	}
	return o.FirstAired, true
}

// HasFirstAired returns a boolean if a field has been set.
func (o *SonarrSeries) HasFirstAired() bool {
	if o != nil && !IsNil(o.FirstAired) {
		return true
	}

	return false
}

// SetFirstAired gets a reference to the given string and assigns it to the FirstAired field.
func (o *SonarrSeries) SetFirstAired(v string) {
	o.FirstAired = &v
}

// GetLastInfoSync returns the LastInfoSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSeries) GetLastInfoSync() string {
	if o == nil || IsNil(o.LastInfoSync.Get()) {
		var ret string
		return ret
	}
	return *o.LastInfoSync.Get()
}

// GetLastInfoSyncOk returns a tuple with the LastInfoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSeries) GetLastInfoSyncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastInfoSync.Get(), o.LastInfoSync.IsSet()
}

// HasLastInfoSync returns a boolean if a field has been set.
func (o *SonarrSeries) HasLastInfoSync() bool {
	if o != nil && o.LastInfoSync.IsSet() {
		return true
	}

	return false
}

// SetLastInfoSync gets a reference to the given NullableString and assigns it to the LastInfoSync field.
func (o *SonarrSeries) SetLastInfoSync(v string) {
	o.LastInfoSync.Set(&v)
}
// SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil
func (o *SonarrSeries) SetLastInfoSyncNil() {
	o.LastInfoSync.Set(nil)
}

// UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
func (o *SonarrSeries) UnsetLastInfoSync() {
	o.LastInfoSync.Unset()
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *SonarrSeries) GetSeriesType() string {
	if o == nil || IsNil(o.SeriesType) {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetSeriesTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesType) {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *SonarrSeries) HasSeriesType() bool {
	if o != nil && !IsNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *SonarrSeries) SetSeriesType(v string) {
	o.SeriesType = &v
}

// GetCleanTitle returns the CleanTitle field value if set, zero value otherwise.
func (o *SonarrSeries) GetCleanTitle() string {
	if o == nil || IsNil(o.CleanTitle) {
		var ret string
		return ret
	}
	return *o.CleanTitle
}

// GetCleanTitleOk returns a tuple with the CleanTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetCleanTitleOk() (*string, bool) {
	if o == nil || IsNil(o.CleanTitle) {
		return nil, false
	}
	return o.CleanTitle, true
}

// HasCleanTitle returns a boolean if a field has been set.
func (o *SonarrSeries) HasCleanTitle() bool {
	if o != nil && !IsNil(o.CleanTitle) {
		return true
	}

	return false
}

// SetCleanTitle gets a reference to the given string and assigns it to the CleanTitle field.
func (o *SonarrSeries) SetCleanTitle(v string) {
	o.CleanTitle = &v
}

// GetImdbId returns the ImdbId field value if set, zero value otherwise.
func (o *SonarrSeries) GetImdbId() string {
	if o == nil || IsNil(o.ImdbId) {
		var ret string
		return ret
	}
	return *o.ImdbId
}

// GetImdbIdOk returns a tuple with the ImdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetImdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImdbId) {
		return nil, false
	}
	return o.ImdbId, true
}

// HasImdbId returns a boolean if a field has been set.
func (o *SonarrSeries) HasImdbId() bool {
	if o != nil && !IsNil(o.ImdbId) {
		return true
	}

	return false
}

// SetImdbId gets a reference to the given string and assigns it to the ImdbId field.
func (o *SonarrSeries) SetImdbId(v string) {
	o.ImdbId = &v
}

// GetTitleSlug returns the TitleSlug field value if set, zero value otherwise.
func (o *SonarrSeries) GetTitleSlug() string {
	if o == nil || IsNil(o.TitleSlug) {
		var ret string
		return ret
	}
	return *o.TitleSlug
}

// GetTitleSlugOk returns a tuple with the TitleSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetTitleSlugOk() (*string, bool) {
	if o == nil || IsNil(o.TitleSlug) {
		return nil, false
	}
	return o.TitleSlug, true
}

// HasTitleSlug returns a boolean if a field has been set.
func (o *SonarrSeries) HasTitleSlug() bool {
	if o != nil && !IsNil(o.TitleSlug) {
		return true
	}

	return false
}

// SetTitleSlug gets a reference to the given string and assigns it to the TitleSlug field.
func (o *SonarrSeries) SetTitleSlug(v string) {
	o.TitleSlug = &v
}

// GetCertification returns the Certification field value if set, zero value otherwise.
func (o *SonarrSeries) GetCertification() string {
	if o == nil || IsNil(o.Certification) {
		var ret string
		return ret
	}
	return *o.Certification
}

// GetCertificationOk returns a tuple with the Certification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetCertificationOk() (*string, bool) {
	if o == nil || IsNil(o.Certification) {
		return nil, false
	}
	return o.Certification, true
}

// HasCertification returns a boolean if a field has been set.
func (o *SonarrSeries) HasCertification() bool {
	if o != nil && !IsNil(o.Certification) {
		return true
	}

	return false
}

// SetCertification gets a reference to the given string and assigns it to the Certification field.
func (o *SonarrSeries) SetCertification(v string) {
	o.Certification = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SonarrSeries) GetGenres() []string {
	if o == nil || IsNil(o.Genres) {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SonarrSeries) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *SonarrSeries) SetGenres(v []string) {
	o.Genres = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SonarrSeries) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SonarrSeries) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SonarrSeries) SetTags(v []string) {
	o.Tags = v
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *SonarrSeries) GetAdded() string {
	if o == nil || IsNil(o.Added) {
		var ret string
		return ret
	}
	return *o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetAddedOk() (*string, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *SonarrSeries) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given string and assigns it to the Added field.
func (o *SonarrSeries) SetAdded(v string) {
	o.Added = &v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *SonarrSeries) GetRatings() []SonarrSeriesRatingsInner {
	if o == nil || IsNil(o.Ratings) {
		var ret []SonarrSeriesRatingsInner
		return ret
	}
	return o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetRatingsOk() ([]SonarrSeriesRatingsInner, bool) {
	if o == nil || IsNil(o.Ratings) {
		return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *SonarrSeries) HasRatings() bool {
	if o != nil && !IsNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given []SonarrSeriesRatingsInner and assigns it to the Ratings field.
func (o *SonarrSeries) SetRatings(v []SonarrSeriesRatingsInner) {
	o.Ratings = v
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise.
func (o *SonarrSeries) GetQualityProfileId() float32 {
	if o == nil || IsNil(o.QualityProfileId) {
		var ret float32
		return ret
	}
	return *o.QualityProfileId
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetQualityProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.QualityProfileId) {
		return nil, false
	}
	return o.QualityProfileId, true
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *SonarrSeries) HasQualityProfileId() bool {
	if o != nil && !IsNil(o.QualityProfileId) {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given float32 and assigns it to the QualityProfileId field.
func (o *SonarrSeries) SetQualityProfileId(v float32) {
	o.QualityProfileId = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSeries) GetId() float32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret float32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSeries) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SonarrSeries) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableFloat32 and assigns it to the Id field.
func (o *SonarrSeries) SetId(v float32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SonarrSeries) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SonarrSeries) UnsetId() {
	o.Id.Unset()
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSeries) GetRootFolderPath() string {
	if o == nil || IsNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSeries) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *SonarrSeries) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *SonarrSeries) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *SonarrSeries) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *SonarrSeries) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetAddOptions returns the AddOptions field value if set, zero value otherwise.
func (o *SonarrSeries) GetAddOptions() []SonarrSeriesAddOptionsInner {
	if o == nil || IsNil(o.AddOptions) {
		var ret []SonarrSeriesAddOptionsInner
		return ret
	}
	return o.AddOptions
}

// GetAddOptionsOk returns a tuple with the AddOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeries) GetAddOptionsOk() ([]SonarrSeriesAddOptionsInner, bool) {
	if o == nil || IsNil(o.AddOptions) {
		return nil, false
	}
	return o.AddOptions, true
}

// HasAddOptions returns a boolean if a field has been set.
func (o *SonarrSeries) HasAddOptions() bool {
	if o != nil && !IsNil(o.AddOptions) {
		return true
	}

	return false
}

// SetAddOptions gets a reference to the given []SonarrSeriesAddOptionsInner and assigns it to the AddOptions field.
func (o *SonarrSeries) SetAddOptions(v []SonarrSeriesAddOptionsInner) {
	o.AddOptions = v
}

func (o SonarrSeries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SonarrSeries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.SortTitle) {
		toSerialize["sortTitle"] = o.SortTitle
	}
	if !IsNil(o.SeasonCount) {
		toSerialize["seasonCount"] = o.SeasonCount
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.AirTime) {
		toSerialize["airTime"] = o.AirTime
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.RemotePoster) {
		toSerialize["remotePoster"] = o.RemotePoster
	}
	if !IsNil(o.Seasons) {
		toSerialize["seasons"] = o.Seasons
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.LanguageProfileId) {
		toSerialize["languageProfileId"] = o.LanguageProfileId
	}
	if !IsNil(o.SeasonFolder) {
		toSerialize["seasonFolder"] = o.SeasonFolder
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !IsNil(o.UseSceneNumbering) {
		toSerialize["useSceneNumbering"] = o.UseSceneNumbering
	}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	if !IsNil(o.TvdbId) {
		toSerialize["tvdbId"] = o.TvdbId
	}
	if !IsNil(o.TvRageId) {
		toSerialize["tvRageId"] = o.TvRageId
	}
	if !IsNil(o.TvMazeId) {
		toSerialize["tvMazeId"] = o.TvMazeId
	}
	if !IsNil(o.FirstAired) {
		toSerialize["firstAired"] = o.FirstAired
	}
	if o.LastInfoSync.IsSet() {
		toSerialize["lastInfoSync"] = o.LastInfoSync.Get()
	}
	if !IsNil(o.SeriesType) {
		toSerialize["seriesType"] = o.SeriesType
	}
	if !IsNil(o.CleanTitle) {
		toSerialize["cleanTitle"] = o.CleanTitle
	}
	if !IsNil(o.ImdbId) {
		toSerialize["imdbId"] = o.ImdbId
	}
	if !IsNil(o.TitleSlug) {
		toSerialize["titleSlug"] = o.TitleSlug
	}
	if !IsNil(o.Certification) {
		toSerialize["certification"] = o.Certification
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if !IsNil(o.QualityProfileId) {
		toSerialize["qualityProfileId"] = o.QualityProfileId
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if !IsNil(o.AddOptions) {
		toSerialize["addOptions"] = o.AddOptions
	}
	return toSerialize, nil
}

type NullableSonarrSeries struct {
	value *SonarrSeries
	isSet bool
}

func (v NullableSonarrSeries) Get() *SonarrSeries {
	return v.value
}

func (v *NullableSonarrSeries) Set(val *SonarrSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarrSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarrSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarrSeries(val *SonarrSeries) *NullableSonarrSeries {
	return &NullableSonarrSeries{value: val, isSet: true}
}

func (v NullableSonarrSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarrSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


