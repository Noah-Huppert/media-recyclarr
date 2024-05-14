# ModelLibraryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableArchiveMediaFiles** | Pointer to **bool** |  | [optional] 
**EnablePhotos** | Pointer to **bool** |  | [optional] 
**EnableRealtimeMonitor** | Pointer to **bool** |  | [optional] 
**EnableMarkerDetection** | Pointer to **bool** |  | [optional] 
**EnableMarkerDetectionDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**IntroDetectionFingerprintLength** | Pointer to **int32** |  | [optional] 
**EnableChapterImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractChapterImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**DownloadImagesInAdvance** | Pointer to **bool** |  | [optional] 
**PathInfos** | Pointer to [**[]ModelModelMediaPathInfo**](ModelModelMediaPathInfo.md) |  | [optional] 
**IgnoreHiddenFiles** | Pointer to **bool** |  | [optional] 
**IgnoreFileExtensions** | Pointer to **[]string** |  | [optional] 
**SaveLocalMetadata** | Pointer to **bool** |  | [optional] 
**SaveMetadataHidden** | Pointer to **bool** |  | [optional] 
**SaveLocalThumbnailSets** | Pointer to **bool** |  | [optional] 
**ImportPlaylists** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSeriesGrouping** | Pointer to **bool** |  | [optional] 
**ShareEmbeddedMusicAlbumImages** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedTitles** | Pointer to **bool** |  | [optional] 
**EnableAudioResume** | Pointer to **bool** |  | [optional] 
**AutoGenerateChapters** | Pointer to **bool** |  | [optional] 
**AutomaticRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PlaceholderMetadataRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**PreferredImageLanguage** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**MetadataSavers** | Pointer to **[]string** |  | [optional] 
**DisabledLocalMetadataReaders** | Pointer to **[]string** |  | [optional] 
**LocalMetadataReaderOrder** | Pointer to **[]string** |  | [optional] 
**DisabledLyricsFetchers** | Pointer to **[]string** |  | [optional] 
**SaveLyricsWithMedia** | Pointer to **bool** |  | [optional] 
**LyricsDownloadMaxAgeDays** | Pointer to **int32** |  | [optional] 
**LyricsFetcherOrder** | Pointer to **[]string** |  | [optional] 
**LyricsDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**DisabledSubtitleFetchers** | Pointer to **[]string** |  | [optional] 
**SubtitleFetcherOrder** | Pointer to **[]string** |  | [optional] 
**SkipSubtitlesIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipSubtitlesIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**SubtitleDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**SubtitleDownloadMaxAgeDays** | Pointer to **int32** |  | [optional] 
**RequirePerfectSubtitleMatch** | Pointer to **bool** |  | [optional] 
**SaveSubtitlesWithMedia** | Pointer to **bool** |  | [optional] 
**ForcedSubtitlesOnly** | Pointer to **bool** |  | [optional] 
**HearingImpairedSubtitlesOnly** | Pointer to **bool** |  | [optional] 
**TypeOptions** | Pointer to [**[]ModelModelTypeOptions**](ModelModelTypeOptions.md) |  | [optional] 
**CollapseSingleItemFolders** | Pointer to **bool** |  | [optional] 
**EnableAdultMetadata** | Pointer to **bool** |  | [optional] 
**ImportCollections** | Pointer to **bool** |  | [optional] 
**MinCollectionItems** | Pointer to **int32** |  | [optional] 
**MusicFolderStructure** | Pointer to **string** |  | [optional] 
**MinResumePct** | Pointer to **int32** |  | [optional] 
**MaxResumePct** | Pointer to **int32** |  | [optional] 
**MinResumeDurationSeconds** | Pointer to **int32** |  | [optional] 
**ThumbnailImagesIntervalSeconds** | Pointer to **int32** |  | [optional] 
**SampleIgnoreSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelLibraryOptions

`func NewModelLibraryOptions() *ModelLibraryOptions`

NewModelLibraryOptions instantiates a new ModelLibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLibraryOptionsWithDefaults

`func NewModelLibraryOptionsWithDefaults() *ModelLibraryOptions`

NewModelLibraryOptionsWithDefaults instantiates a new ModelLibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableArchiveMediaFiles

`func (o *ModelLibraryOptions) GetEnableArchiveMediaFiles() bool`

GetEnableArchiveMediaFiles returns the EnableArchiveMediaFiles field if non-nil, zero value otherwise.

### GetEnableArchiveMediaFilesOk

`func (o *ModelLibraryOptions) GetEnableArchiveMediaFilesOk() (*bool, bool)`

GetEnableArchiveMediaFilesOk returns a tuple with the EnableArchiveMediaFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArchiveMediaFiles

`func (o *ModelLibraryOptions) SetEnableArchiveMediaFiles(v bool)`

SetEnableArchiveMediaFiles sets EnableArchiveMediaFiles field to given value.

### HasEnableArchiveMediaFiles

`func (o *ModelLibraryOptions) HasEnableArchiveMediaFiles() bool`

HasEnableArchiveMediaFiles returns a boolean if a field has been set.

### GetEnablePhotos

`func (o *ModelLibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *ModelLibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *ModelLibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *ModelLibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *ModelLibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *ModelLibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *ModelLibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *ModelLibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableMarkerDetection

`func (o *ModelLibraryOptions) GetEnableMarkerDetection() bool`

GetEnableMarkerDetection returns the EnableMarkerDetection field if non-nil, zero value otherwise.

### GetEnableMarkerDetectionOk

`func (o *ModelLibraryOptions) GetEnableMarkerDetectionOk() (*bool, bool)`

GetEnableMarkerDetectionOk returns a tuple with the EnableMarkerDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarkerDetection

`func (o *ModelLibraryOptions) SetEnableMarkerDetection(v bool)`

SetEnableMarkerDetection sets EnableMarkerDetection field to given value.

### HasEnableMarkerDetection

`func (o *ModelLibraryOptions) HasEnableMarkerDetection() bool`

HasEnableMarkerDetection returns a boolean if a field has been set.

### GetEnableMarkerDetectionDuringLibraryScan

`func (o *ModelLibraryOptions) GetEnableMarkerDetectionDuringLibraryScan() bool`

GetEnableMarkerDetectionDuringLibraryScan returns the EnableMarkerDetectionDuringLibraryScan field if non-nil, zero value otherwise.

### GetEnableMarkerDetectionDuringLibraryScanOk

`func (o *ModelLibraryOptions) GetEnableMarkerDetectionDuringLibraryScanOk() (*bool, bool)`

GetEnableMarkerDetectionDuringLibraryScanOk returns a tuple with the EnableMarkerDetectionDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarkerDetectionDuringLibraryScan

`func (o *ModelLibraryOptions) SetEnableMarkerDetectionDuringLibraryScan(v bool)`

SetEnableMarkerDetectionDuringLibraryScan sets EnableMarkerDetectionDuringLibraryScan field to given value.

### HasEnableMarkerDetectionDuringLibraryScan

`func (o *ModelLibraryOptions) HasEnableMarkerDetectionDuringLibraryScan() bool`

HasEnableMarkerDetectionDuringLibraryScan returns a boolean if a field has been set.

### GetIntroDetectionFingerprintLength

`func (o *ModelLibraryOptions) GetIntroDetectionFingerprintLength() int32`

GetIntroDetectionFingerprintLength returns the IntroDetectionFingerprintLength field if non-nil, zero value otherwise.

### GetIntroDetectionFingerprintLengthOk

`func (o *ModelLibraryOptions) GetIntroDetectionFingerprintLengthOk() (*int32, bool)`

GetIntroDetectionFingerprintLengthOk returns a tuple with the IntroDetectionFingerprintLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroDetectionFingerprintLength

`func (o *ModelLibraryOptions) SetIntroDetectionFingerprintLength(v int32)`

SetIntroDetectionFingerprintLength sets IntroDetectionFingerprintLength field to given value.

### HasIntroDetectionFingerprintLength

`func (o *ModelLibraryOptions) HasIntroDetectionFingerprintLength() bool`

HasIntroDetectionFingerprintLength returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *ModelLibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *ModelLibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *ModelLibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *ModelLibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *ModelLibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *ModelLibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *ModelLibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *ModelLibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetDownloadImagesInAdvance

`func (o *ModelLibraryOptions) GetDownloadImagesInAdvance() bool`

GetDownloadImagesInAdvance returns the DownloadImagesInAdvance field if non-nil, zero value otherwise.

### GetDownloadImagesInAdvanceOk

`func (o *ModelLibraryOptions) GetDownloadImagesInAdvanceOk() (*bool, bool)`

GetDownloadImagesInAdvanceOk returns a tuple with the DownloadImagesInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadImagesInAdvance

`func (o *ModelLibraryOptions) SetDownloadImagesInAdvance(v bool)`

SetDownloadImagesInAdvance sets DownloadImagesInAdvance field to given value.

### HasDownloadImagesInAdvance

`func (o *ModelLibraryOptions) HasDownloadImagesInAdvance() bool`

HasDownloadImagesInAdvance returns a boolean if a field has been set.

### GetPathInfos

`func (o *ModelLibraryOptions) GetPathInfos() []ModelModelMediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *ModelLibraryOptions) GetPathInfosOk() (*[]ModelModelMediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *ModelLibraryOptions) SetPathInfos(v []ModelModelMediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *ModelLibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetIgnoreHiddenFiles

`func (o *ModelLibraryOptions) GetIgnoreHiddenFiles() bool`

GetIgnoreHiddenFiles returns the IgnoreHiddenFiles field if non-nil, zero value otherwise.

### GetIgnoreHiddenFilesOk

`func (o *ModelLibraryOptions) GetIgnoreHiddenFilesOk() (*bool, bool)`

GetIgnoreHiddenFilesOk returns a tuple with the IgnoreHiddenFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreHiddenFiles

`func (o *ModelLibraryOptions) SetIgnoreHiddenFiles(v bool)`

SetIgnoreHiddenFiles sets IgnoreHiddenFiles field to given value.

### HasIgnoreHiddenFiles

`func (o *ModelLibraryOptions) HasIgnoreHiddenFiles() bool`

HasIgnoreHiddenFiles returns a boolean if a field has been set.

### GetIgnoreFileExtensions

`func (o *ModelLibraryOptions) GetIgnoreFileExtensions() []string`

GetIgnoreFileExtensions returns the IgnoreFileExtensions field if non-nil, zero value otherwise.

### GetIgnoreFileExtensionsOk

`func (o *ModelLibraryOptions) GetIgnoreFileExtensionsOk() (*[]string, bool)`

GetIgnoreFileExtensionsOk returns a tuple with the IgnoreFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFileExtensions

`func (o *ModelLibraryOptions) SetIgnoreFileExtensions(v []string)`

SetIgnoreFileExtensions sets IgnoreFileExtensions field to given value.

### HasIgnoreFileExtensions

`func (o *ModelLibraryOptions) HasIgnoreFileExtensions() bool`

HasIgnoreFileExtensions returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *ModelLibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *ModelLibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *ModelLibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *ModelLibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetSaveMetadataHidden

`func (o *ModelLibraryOptions) GetSaveMetadataHidden() bool`

GetSaveMetadataHidden returns the SaveMetadataHidden field if non-nil, zero value otherwise.

### GetSaveMetadataHiddenOk

`func (o *ModelLibraryOptions) GetSaveMetadataHiddenOk() (*bool, bool)`

GetSaveMetadataHiddenOk returns a tuple with the SaveMetadataHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMetadataHidden

`func (o *ModelLibraryOptions) SetSaveMetadataHidden(v bool)`

SetSaveMetadataHidden sets SaveMetadataHidden field to given value.

### HasSaveMetadataHidden

`func (o *ModelLibraryOptions) HasSaveMetadataHidden() bool`

HasSaveMetadataHidden returns a boolean if a field has been set.

### GetSaveLocalThumbnailSets

`func (o *ModelLibraryOptions) GetSaveLocalThumbnailSets() bool`

GetSaveLocalThumbnailSets returns the SaveLocalThumbnailSets field if non-nil, zero value otherwise.

### GetSaveLocalThumbnailSetsOk

`func (o *ModelLibraryOptions) GetSaveLocalThumbnailSetsOk() (*bool, bool)`

GetSaveLocalThumbnailSetsOk returns a tuple with the SaveLocalThumbnailSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalThumbnailSets

`func (o *ModelLibraryOptions) SetSaveLocalThumbnailSets(v bool)`

SetSaveLocalThumbnailSets sets SaveLocalThumbnailSets field to given value.

### HasSaveLocalThumbnailSets

`func (o *ModelLibraryOptions) HasSaveLocalThumbnailSets() bool`

HasSaveLocalThumbnailSets returns a boolean if a field has been set.

### GetImportPlaylists

`func (o *ModelLibraryOptions) GetImportPlaylists() bool`

GetImportPlaylists returns the ImportPlaylists field if non-nil, zero value otherwise.

### GetImportPlaylistsOk

`func (o *ModelLibraryOptions) GetImportPlaylistsOk() (*bool, bool)`

GetImportPlaylistsOk returns a tuple with the ImportPlaylists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPlaylists

`func (o *ModelLibraryOptions) SetImportPlaylists(v bool)`

SetImportPlaylists sets ImportPlaylists field to given value.

### HasImportPlaylists

`func (o *ModelLibraryOptions) HasImportPlaylists() bool`

HasImportPlaylists returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *ModelLibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *ModelLibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *ModelLibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *ModelLibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetShareEmbeddedMusicAlbumImages

`func (o *ModelLibraryOptions) GetShareEmbeddedMusicAlbumImages() bool`

GetShareEmbeddedMusicAlbumImages returns the ShareEmbeddedMusicAlbumImages field if non-nil, zero value otherwise.

### GetShareEmbeddedMusicAlbumImagesOk

`func (o *ModelLibraryOptions) GetShareEmbeddedMusicAlbumImagesOk() (*bool, bool)`

GetShareEmbeddedMusicAlbumImagesOk returns a tuple with the ShareEmbeddedMusicAlbumImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareEmbeddedMusicAlbumImages

`func (o *ModelLibraryOptions) SetShareEmbeddedMusicAlbumImages(v bool)`

SetShareEmbeddedMusicAlbumImages sets ShareEmbeddedMusicAlbumImages field to given value.

### HasShareEmbeddedMusicAlbumImages

`func (o *ModelLibraryOptions) HasShareEmbeddedMusicAlbumImages() bool`

HasShareEmbeddedMusicAlbumImages returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *ModelLibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *ModelLibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *ModelLibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *ModelLibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableAudioResume

`func (o *ModelLibraryOptions) GetEnableAudioResume() bool`

GetEnableAudioResume returns the EnableAudioResume field if non-nil, zero value otherwise.

### GetEnableAudioResumeOk

`func (o *ModelLibraryOptions) GetEnableAudioResumeOk() (*bool, bool)`

GetEnableAudioResumeOk returns a tuple with the EnableAudioResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioResume

`func (o *ModelLibraryOptions) SetEnableAudioResume(v bool)`

SetEnableAudioResume sets EnableAudioResume field to given value.

### HasEnableAudioResume

`func (o *ModelLibraryOptions) HasEnableAudioResume() bool`

HasEnableAudioResume returns a boolean if a field has been set.

### GetAutoGenerateChapters

`func (o *ModelLibraryOptions) GetAutoGenerateChapters() bool`

GetAutoGenerateChapters returns the AutoGenerateChapters field if non-nil, zero value otherwise.

### GetAutoGenerateChaptersOk

`func (o *ModelLibraryOptions) GetAutoGenerateChaptersOk() (*bool, bool)`

GetAutoGenerateChaptersOk returns a tuple with the AutoGenerateChapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGenerateChapters

`func (o *ModelLibraryOptions) SetAutoGenerateChapters(v bool)`

SetAutoGenerateChapters sets AutoGenerateChapters field to given value.

### HasAutoGenerateChapters

`func (o *ModelLibraryOptions) HasAutoGenerateChapters() bool`

HasAutoGenerateChapters returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *ModelLibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *ModelLibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *ModelLibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *ModelLibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPlaceholderMetadataRefreshIntervalDays

`func (o *ModelLibraryOptions) GetPlaceholderMetadataRefreshIntervalDays() int32`

GetPlaceholderMetadataRefreshIntervalDays returns the PlaceholderMetadataRefreshIntervalDays field if non-nil, zero value otherwise.

### GetPlaceholderMetadataRefreshIntervalDaysOk

`func (o *ModelLibraryOptions) GetPlaceholderMetadataRefreshIntervalDaysOk() (*int32, bool)`

GetPlaceholderMetadataRefreshIntervalDaysOk returns a tuple with the PlaceholderMetadataRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderMetadataRefreshIntervalDays

`func (o *ModelLibraryOptions) SetPlaceholderMetadataRefreshIntervalDays(v int32)`

SetPlaceholderMetadataRefreshIntervalDays sets PlaceholderMetadataRefreshIntervalDays field to given value.

### HasPlaceholderMetadataRefreshIntervalDays

`func (o *ModelLibraryOptions) HasPlaceholderMetadataRefreshIntervalDays() bool`

HasPlaceholderMetadataRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ModelLibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ModelLibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ModelLibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ModelLibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetPreferredImageLanguage

`func (o *ModelLibraryOptions) GetPreferredImageLanguage() string`

GetPreferredImageLanguage returns the PreferredImageLanguage field if non-nil, zero value otherwise.

### GetPreferredImageLanguageOk

`func (o *ModelLibraryOptions) GetPreferredImageLanguageOk() (*string, bool)`

GetPreferredImageLanguageOk returns a tuple with the PreferredImageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredImageLanguage

`func (o *ModelLibraryOptions) SetPreferredImageLanguage(v string)`

SetPreferredImageLanguage sets PreferredImageLanguage field to given value.

### HasPreferredImageLanguage

`func (o *ModelLibraryOptions) HasPreferredImageLanguage() bool`

HasPreferredImageLanguage returns a boolean if a field has been set.

### GetContentType

`func (o *ModelLibraryOptions) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ModelLibraryOptions) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ModelLibraryOptions) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ModelLibraryOptions) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ModelLibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ModelLibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ModelLibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ModelLibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *ModelLibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *ModelLibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *ModelLibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *ModelLibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### GetDisabledLocalMetadataReaders

`func (o *ModelLibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *ModelLibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *ModelLibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *ModelLibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *ModelLibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *ModelLibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *ModelLibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *ModelLibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### GetDisabledLyricsFetchers

`func (o *ModelLibraryOptions) GetDisabledLyricsFetchers() []string`

GetDisabledLyricsFetchers returns the DisabledLyricsFetchers field if non-nil, zero value otherwise.

### GetDisabledLyricsFetchersOk

`func (o *ModelLibraryOptions) GetDisabledLyricsFetchersOk() (*[]string, bool)`

GetDisabledLyricsFetchersOk returns a tuple with the DisabledLyricsFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLyricsFetchers

`func (o *ModelLibraryOptions) SetDisabledLyricsFetchers(v []string)`

SetDisabledLyricsFetchers sets DisabledLyricsFetchers field to given value.

### HasDisabledLyricsFetchers

`func (o *ModelLibraryOptions) HasDisabledLyricsFetchers() bool`

HasDisabledLyricsFetchers returns a boolean if a field has been set.

### GetSaveLyricsWithMedia

`func (o *ModelLibraryOptions) GetSaveLyricsWithMedia() bool`

GetSaveLyricsWithMedia returns the SaveLyricsWithMedia field if non-nil, zero value otherwise.

### GetSaveLyricsWithMediaOk

`func (o *ModelLibraryOptions) GetSaveLyricsWithMediaOk() (*bool, bool)`

GetSaveLyricsWithMediaOk returns a tuple with the SaveLyricsWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLyricsWithMedia

`func (o *ModelLibraryOptions) SetSaveLyricsWithMedia(v bool)`

SetSaveLyricsWithMedia sets SaveLyricsWithMedia field to given value.

### HasSaveLyricsWithMedia

`func (o *ModelLibraryOptions) HasSaveLyricsWithMedia() bool`

HasSaveLyricsWithMedia returns a boolean if a field has been set.

### GetLyricsDownloadMaxAgeDays

`func (o *ModelLibraryOptions) GetLyricsDownloadMaxAgeDays() int32`

GetLyricsDownloadMaxAgeDays returns the LyricsDownloadMaxAgeDays field if non-nil, zero value otherwise.

### GetLyricsDownloadMaxAgeDaysOk

`func (o *ModelLibraryOptions) GetLyricsDownloadMaxAgeDaysOk() (*int32, bool)`

GetLyricsDownloadMaxAgeDaysOk returns a tuple with the LyricsDownloadMaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsDownloadMaxAgeDays

`func (o *ModelLibraryOptions) SetLyricsDownloadMaxAgeDays(v int32)`

SetLyricsDownloadMaxAgeDays sets LyricsDownloadMaxAgeDays field to given value.

### HasLyricsDownloadMaxAgeDays

`func (o *ModelLibraryOptions) HasLyricsDownloadMaxAgeDays() bool`

HasLyricsDownloadMaxAgeDays returns a boolean if a field has been set.

### GetLyricsFetcherOrder

`func (o *ModelLibraryOptions) GetLyricsFetcherOrder() []string`

GetLyricsFetcherOrder returns the LyricsFetcherOrder field if non-nil, zero value otherwise.

### GetLyricsFetcherOrderOk

`func (o *ModelLibraryOptions) GetLyricsFetcherOrderOk() (*[]string, bool)`

GetLyricsFetcherOrderOk returns a tuple with the LyricsFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsFetcherOrder

`func (o *ModelLibraryOptions) SetLyricsFetcherOrder(v []string)`

SetLyricsFetcherOrder sets LyricsFetcherOrder field to given value.

### HasLyricsFetcherOrder

`func (o *ModelLibraryOptions) HasLyricsFetcherOrder() bool`

HasLyricsFetcherOrder returns a boolean if a field has been set.

### GetLyricsDownloadLanguages

`func (o *ModelLibraryOptions) GetLyricsDownloadLanguages() []string`

GetLyricsDownloadLanguages returns the LyricsDownloadLanguages field if non-nil, zero value otherwise.

### GetLyricsDownloadLanguagesOk

`func (o *ModelLibraryOptions) GetLyricsDownloadLanguagesOk() (*[]string, bool)`

GetLyricsDownloadLanguagesOk returns a tuple with the LyricsDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricsDownloadLanguages

`func (o *ModelLibraryOptions) SetLyricsDownloadLanguages(v []string)`

SetLyricsDownloadLanguages sets LyricsDownloadLanguages field to given value.

### HasLyricsDownloadLanguages

`func (o *ModelLibraryOptions) HasLyricsDownloadLanguages() bool`

HasLyricsDownloadLanguages returns a boolean if a field has been set.

### GetDisabledSubtitleFetchers

`func (o *ModelLibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *ModelLibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *ModelLibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *ModelLibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *ModelLibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *ModelLibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *ModelLibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *ModelLibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ModelLibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ModelLibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *ModelLibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *ModelLibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *ModelLibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *ModelLibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *ModelLibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *ModelLibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### GetSubtitleDownloadMaxAgeDays

`func (o *ModelLibraryOptions) GetSubtitleDownloadMaxAgeDays() int32`

GetSubtitleDownloadMaxAgeDays returns the SubtitleDownloadMaxAgeDays field if non-nil, zero value otherwise.

### GetSubtitleDownloadMaxAgeDaysOk

`func (o *ModelLibraryOptions) GetSubtitleDownloadMaxAgeDaysOk() (*int32, bool)`

GetSubtitleDownloadMaxAgeDaysOk returns a tuple with the SubtitleDownloadMaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadMaxAgeDays

`func (o *ModelLibraryOptions) SetSubtitleDownloadMaxAgeDays(v int32)`

SetSubtitleDownloadMaxAgeDays sets SubtitleDownloadMaxAgeDays field to given value.

### HasSubtitleDownloadMaxAgeDays

`func (o *ModelLibraryOptions) HasSubtitleDownloadMaxAgeDays() bool`

HasSubtitleDownloadMaxAgeDays returns a boolean if a field has been set.

### GetRequirePerfectSubtitleMatch

`func (o *ModelLibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *ModelLibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *ModelLibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *ModelLibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *ModelLibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *ModelLibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *ModelLibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *ModelLibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetForcedSubtitlesOnly

`func (o *ModelLibraryOptions) GetForcedSubtitlesOnly() bool`

GetForcedSubtitlesOnly returns the ForcedSubtitlesOnly field if non-nil, zero value otherwise.

### GetForcedSubtitlesOnlyOk

`func (o *ModelLibraryOptions) GetForcedSubtitlesOnlyOk() (*bool, bool)`

GetForcedSubtitlesOnlyOk returns a tuple with the ForcedSubtitlesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSubtitlesOnly

`func (o *ModelLibraryOptions) SetForcedSubtitlesOnly(v bool)`

SetForcedSubtitlesOnly sets ForcedSubtitlesOnly field to given value.

### HasForcedSubtitlesOnly

`func (o *ModelLibraryOptions) HasForcedSubtitlesOnly() bool`

HasForcedSubtitlesOnly returns a boolean if a field has been set.

### GetHearingImpairedSubtitlesOnly

`func (o *ModelLibraryOptions) GetHearingImpairedSubtitlesOnly() bool`

GetHearingImpairedSubtitlesOnly returns the HearingImpairedSubtitlesOnly field if non-nil, zero value otherwise.

### GetHearingImpairedSubtitlesOnlyOk

`func (o *ModelLibraryOptions) GetHearingImpairedSubtitlesOnlyOk() (*bool, bool)`

GetHearingImpairedSubtitlesOnlyOk returns a tuple with the HearingImpairedSubtitlesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHearingImpairedSubtitlesOnly

`func (o *ModelLibraryOptions) SetHearingImpairedSubtitlesOnly(v bool)`

SetHearingImpairedSubtitlesOnly sets HearingImpairedSubtitlesOnly field to given value.

### HasHearingImpairedSubtitlesOnly

`func (o *ModelLibraryOptions) HasHearingImpairedSubtitlesOnly() bool`

HasHearingImpairedSubtitlesOnly returns a boolean if a field has been set.

### GetTypeOptions

`func (o *ModelLibraryOptions) GetTypeOptions() []ModelModelTypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *ModelLibraryOptions) GetTypeOptionsOk() (*[]ModelModelTypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *ModelLibraryOptions) SetTypeOptions(v []ModelModelTypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *ModelLibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.

### GetCollapseSingleItemFolders

`func (o *ModelLibraryOptions) GetCollapseSingleItemFolders() bool`

GetCollapseSingleItemFolders returns the CollapseSingleItemFolders field if non-nil, zero value otherwise.

### GetCollapseSingleItemFoldersOk

`func (o *ModelLibraryOptions) GetCollapseSingleItemFoldersOk() (*bool, bool)`

GetCollapseSingleItemFoldersOk returns a tuple with the CollapseSingleItemFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseSingleItemFolders

`func (o *ModelLibraryOptions) SetCollapseSingleItemFolders(v bool)`

SetCollapseSingleItemFolders sets CollapseSingleItemFolders field to given value.

### HasCollapseSingleItemFolders

`func (o *ModelLibraryOptions) HasCollapseSingleItemFolders() bool`

HasCollapseSingleItemFolders returns a boolean if a field has been set.

### GetEnableAdultMetadata

`func (o *ModelLibraryOptions) GetEnableAdultMetadata() bool`

GetEnableAdultMetadata returns the EnableAdultMetadata field if non-nil, zero value otherwise.

### GetEnableAdultMetadataOk

`func (o *ModelLibraryOptions) GetEnableAdultMetadataOk() (*bool, bool)`

GetEnableAdultMetadataOk returns a tuple with the EnableAdultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdultMetadata

`func (o *ModelLibraryOptions) SetEnableAdultMetadata(v bool)`

SetEnableAdultMetadata sets EnableAdultMetadata field to given value.

### HasEnableAdultMetadata

`func (o *ModelLibraryOptions) HasEnableAdultMetadata() bool`

HasEnableAdultMetadata returns a boolean if a field has been set.

### GetImportCollections

`func (o *ModelLibraryOptions) GetImportCollections() bool`

GetImportCollections returns the ImportCollections field if non-nil, zero value otherwise.

### GetImportCollectionsOk

`func (o *ModelLibraryOptions) GetImportCollectionsOk() (*bool, bool)`

GetImportCollectionsOk returns a tuple with the ImportCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportCollections

`func (o *ModelLibraryOptions) SetImportCollections(v bool)`

SetImportCollections sets ImportCollections field to given value.

### HasImportCollections

`func (o *ModelLibraryOptions) HasImportCollections() bool`

HasImportCollections returns a boolean if a field has been set.

### GetMinCollectionItems

`func (o *ModelLibraryOptions) GetMinCollectionItems() int32`

GetMinCollectionItems returns the MinCollectionItems field if non-nil, zero value otherwise.

### GetMinCollectionItemsOk

`func (o *ModelLibraryOptions) GetMinCollectionItemsOk() (*int32, bool)`

GetMinCollectionItemsOk returns a tuple with the MinCollectionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCollectionItems

`func (o *ModelLibraryOptions) SetMinCollectionItems(v int32)`

SetMinCollectionItems sets MinCollectionItems field to given value.

### HasMinCollectionItems

`func (o *ModelLibraryOptions) HasMinCollectionItems() bool`

HasMinCollectionItems returns a boolean if a field has been set.

### GetMusicFolderStructure

`func (o *ModelLibraryOptions) GetMusicFolderStructure() string`

GetMusicFolderStructure returns the MusicFolderStructure field if non-nil, zero value otherwise.

### GetMusicFolderStructureOk

`func (o *ModelLibraryOptions) GetMusicFolderStructureOk() (*string, bool)`

GetMusicFolderStructureOk returns a tuple with the MusicFolderStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicFolderStructure

`func (o *ModelLibraryOptions) SetMusicFolderStructure(v string)`

SetMusicFolderStructure sets MusicFolderStructure field to given value.

### HasMusicFolderStructure

`func (o *ModelLibraryOptions) HasMusicFolderStructure() bool`

HasMusicFolderStructure returns a boolean if a field has been set.

### GetMinResumePct

`func (o *ModelLibraryOptions) GetMinResumePct() int32`

GetMinResumePct returns the MinResumePct field if non-nil, zero value otherwise.

### GetMinResumePctOk

`func (o *ModelLibraryOptions) GetMinResumePctOk() (*int32, bool)`

GetMinResumePctOk returns a tuple with the MinResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumePct

`func (o *ModelLibraryOptions) SetMinResumePct(v int32)`

SetMinResumePct sets MinResumePct field to given value.

### HasMinResumePct

`func (o *ModelLibraryOptions) HasMinResumePct() bool`

HasMinResumePct returns a boolean if a field has been set.

### GetMaxResumePct

`func (o *ModelLibraryOptions) GetMaxResumePct() int32`

GetMaxResumePct returns the MaxResumePct field if non-nil, zero value otherwise.

### GetMaxResumePctOk

`func (o *ModelLibraryOptions) GetMaxResumePctOk() (*int32, bool)`

GetMaxResumePctOk returns a tuple with the MaxResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResumePct

`func (o *ModelLibraryOptions) SetMaxResumePct(v int32)`

SetMaxResumePct sets MaxResumePct field to given value.

### HasMaxResumePct

`func (o *ModelLibraryOptions) HasMaxResumePct() bool`

HasMaxResumePct returns a boolean if a field has been set.

### GetMinResumeDurationSeconds

`func (o *ModelLibraryOptions) GetMinResumeDurationSeconds() int32`

GetMinResumeDurationSeconds returns the MinResumeDurationSeconds field if non-nil, zero value otherwise.

### GetMinResumeDurationSecondsOk

`func (o *ModelLibraryOptions) GetMinResumeDurationSecondsOk() (*int32, bool)`

GetMinResumeDurationSecondsOk returns a tuple with the MinResumeDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumeDurationSeconds

`func (o *ModelLibraryOptions) SetMinResumeDurationSeconds(v int32)`

SetMinResumeDurationSeconds sets MinResumeDurationSeconds field to given value.

### HasMinResumeDurationSeconds

`func (o *ModelLibraryOptions) HasMinResumeDurationSeconds() bool`

HasMinResumeDurationSeconds returns a boolean if a field has been set.

### GetThumbnailImagesIntervalSeconds

`func (o *ModelLibraryOptions) GetThumbnailImagesIntervalSeconds() int32`

GetThumbnailImagesIntervalSeconds returns the ThumbnailImagesIntervalSeconds field if non-nil, zero value otherwise.

### GetThumbnailImagesIntervalSecondsOk

`func (o *ModelLibraryOptions) GetThumbnailImagesIntervalSecondsOk() (*int32, bool)`

GetThumbnailImagesIntervalSecondsOk returns a tuple with the ThumbnailImagesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImagesIntervalSeconds

`func (o *ModelLibraryOptions) SetThumbnailImagesIntervalSeconds(v int32)`

SetThumbnailImagesIntervalSeconds sets ThumbnailImagesIntervalSeconds field to given value.

### HasThumbnailImagesIntervalSeconds

`func (o *ModelLibraryOptions) HasThumbnailImagesIntervalSeconds() bool`

HasThumbnailImagesIntervalSeconds returns a boolean if a field has been set.

### GetSampleIgnoreSize

`func (o *ModelLibraryOptions) GetSampleIgnoreSize() int32`

GetSampleIgnoreSize returns the SampleIgnoreSize field if non-nil, zero value otherwise.

### GetSampleIgnoreSizeOk

`func (o *ModelLibraryOptions) GetSampleIgnoreSizeOk() (*int32, bool)`

GetSampleIgnoreSizeOk returns a tuple with the SampleIgnoreSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIgnoreSize

`func (o *ModelLibraryOptions) SetSampleIgnoreSize(v int32)`

SetSampleIgnoreSize sets SampleIgnoreSize field to given value.

### HasSampleIgnoreSize

`func (o *ModelLibraryOptions) HasSampleIgnoreSize() bool`

HasSampleIgnoreSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


