# \UserActivityAPIAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserUsageStatsByBreakdowntypeBreakdownreport**](UserActivityAPIAPI.md#GetUserUsageStatsByBreakdowntypeBreakdownreport) | **Get** /user_usage_stats/{BreakdownType}/BreakdownReport | Gets a breakdown of a usage metric
[**GetUserUsageStatsByUseridByDateGetitems**](UserActivityAPIAPI.md#GetUserUsageStatsByUseridByDateGetitems) | **Get** /user_usage_stats/{UserID}/{Date}/GetItems | Gets activity for {USER} for {Date} formatted as yyyy-MM-dd
[**GetUserUsageStatsGetItemPath**](UserActivityAPIAPI.md#GetUserUsageStatsGetItemPath) | **Get** /user_usage_stats/get_item_path | Get a list of items for type and filtered
[**GetUserUsageStatsGetItemStats**](UserActivityAPIAPI.md#GetUserUsageStatsGetItemStats) | **Get** /user_usage_stats/get_item_stats | Get a list of items for type and filtered
[**GetUserUsageStatsGetItems**](UserActivityAPIAPI.md#GetUserUsageStatsGetItems) | **Get** /user_usage_stats/get_items | Get a list of items for type and filtered
[**GetUserUsageStatsHourlyreport**](UserActivityAPIAPI.md#GetUserUsageStatsHourlyreport) | **Get** /user_usage_stats/HourlyReport | Gets a report of the available activity per hour
[**GetUserUsageStatsLoadBackup**](UserActivityAPIAPI.md#GetUserUsageStatsLoadBackup) | **Get** /user_usage_stats/load_backup | Loads a backup from a file
[**GetUserUsageStatsMoviesreport**](UserActivityAPIAPI.md#GetUserUsageStatsMoviesreport) | **Get** /user_usage_stats/MoviesReport | Gets Movies counts
[**GetUserUsageStatsPlayactivity**](UserActivityAPIAPI.md#GetUserUsageStatsPlayactivity) | **Get** /user_usage_stats/PlayActivity | Gets play activity for number of days
[**GetUserUsageStatsSaveBackup**](UserActivityAPIAPI.md#GetUserUsageStatsSaveBackup) | **Get** /user_usage_stats/save_backup | Saves a backup of the playback report data to the backup path
[**GetUserUsageStatsSessionList**](UserActivityAPIAPI.md#GetUserUsageStatsSessionList) | **Get** /user_usage_stats/session_list | Gets Session Info
[**GetUserUsageStatsTvshowsreport**](UserActivityAPIAPI.md#GetUserUsageStatsTvshowsreport) | **Get** /user_usage_stats/TvShowsReport | Gets TV Shows counts
[**GetUserUsageStatsTypeFilterList**](UserActivityAPIAPI.md#GetUserUsageStatsTypeFilterList) | **Get** /user_usage_stats/type_filter_list | Gets types filter list items
[**GetUserUsageStatsUserActivity**](UserActivityAPIAPI.md#GetUserUsageStatsUserActivity) | **Get** /user_usage_stats/user_activity | Gets a report of the available activity per hour
[**GetUserUsageStatsUserList**](UserActivityAPIAPI.md#GetUserUsageStatsUserList) | **Get** /user_usage_stats/user_list | Get users
[**GetUserUsageStatsUserManageByActionById**](UserActivityAPIAPI.md#GetUserUsageStatsUserManageByActionById) | **Get** /user_usage_stats/user_manage/{Action}/{Id} | Get users
[**GetUserUsageStatsUserplaylist**](UserActivityAPIAPI.md#GetUserUsageStatsUserplaylist) | **Get** /user_usage_stats/UserPlaylist | Gets a report of all played items for a user in a date period
[**PostUserUsageStatsImportBackup**](UserActivityAPIAPI.md#PostUserUsageStatsImportBackup) | **Post** /user_usage_stats/import_backup | Post a backup for importing
[**PostUserUsageStatsSubmitCustomQuery**](UserActivityAPIAPI.md#PostUserUsageStatsSubmitCustomQuery) | **Post** /user_usage_stats/submit_custom_query | Submit an SQL query



## GetUserUsageStatsByBreakdowntypeBreakdownreport

> map[string]interface{} GetUserUsageStatsByBreakdowntypeBreakdownreport(ctx, breakdownType).UserId(userId).Days(days).EndDate(endDate).Execute()

Gets a breakdown of a usage metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	breakdownType := "breakdownType_example" // string | Breakdown type
	userId := "userId_example" // string | User Id (optional)
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsByBreakdowntypeBreakdownreport(context.Background(), breakdownType).UserId(userId).Days(days).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsByBreakdowntypeBreakdownreport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsByBreakdowntypeBreakdownreport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsByBreakdowntypeBreakdownreport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**breakdownType** | **string** | Breakdown type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsByBreakdowntypeBreakdownreportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsByUseridByDateGetitems

> map[string]interface{} GetUserUsageStatsByUseridByDateGetitems(ctx, userID, date).Filter(filter).Execute()

Gets activity for {USER} for {Date} formatted as yyyy-MM-dd



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	userID := "userID_example" // string | User Id
	date := "date_example" // string | UTC DateTime, Format yyyy-MM-dd
	filter := "filter_example" // string | Comma separated list of media types to filter (movies,series) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsByUseridByDateGetitems(context.Background(), userID, date).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsByUseridByDateGetitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsByUseridByDateGetitems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsByUseridByDateGetitems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User Id | 
**date** | **string** | UTC DateTime, Format yyyy-MM-dd | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsByUseridByDateGetitemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | Comma separated list of media types to filter (movies,series) | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsGetItemPath

> map[string]interface{} GetUserUsageStatsGetItemPath(ctx).Id(id).Execute()

Get a list of items for type and filtered



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	id := int32(56) // int32 | item id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsGetItemPath(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsGetItemPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsGetItemPath`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsGetItemPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsGetItemPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | item id | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsGetItemStats

> map[string]interface{} GetUserUsageStatsGetItemStats(ctx).Id(id).Execute()

Get a list of items for type and filtered



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	id := int32(56) // int32 | item id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsGetItemStats(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsGetItemStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsGetItemStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsGetItemStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsGetItemStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | item id | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsGetItems

> map[string]interface{} GetUserUsageStatsGetItems(ctx).Filter(filter).ItemType(itemType).Parent(parent).Execute()

Get a list of items for type and filtered



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	filter := "filter_example" // string | filter string (optional)
	itemType := "itemType_example" // string | type of items to return (optional)
	parent := int32(56) // int32 | parentid (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsGetItems(context.Background()).Filter(filter).ItemType(itemType).Parent(parent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsGetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsGetItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsGetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter string | 
 **itemType** | **string** | type of items to return | 
 **parent** | **int32** | parentid | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsHourlyreport

> map[string]interface{} GetUserUsageStatsHourlyreport(ctx).UserId(userId).Days(days).EndDate(endDate).Filter(filter).Execute()

Gets a report of the available activity per hour



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	userId := "userId_example" // string | User Id (optional)
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)
	filter := "filter_example" // string | Comma separated list of media types to filter (movies,series) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsHourlyreport(context.Background()).UserId(userId).Days(days).EndDate(endDate).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsHourlyreport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsHourlyreport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsHourlyreport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsHourlyreportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 
 **filter** | **string** | Comma separated list of media types to filter (movies,series) | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsLoadBackup

> map[string]interface{} GetUserUsageStatsLoadBackup(ctx).Backupfile(backupfile).Execute()

Loads a backup from a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	backupfile := "backupfile_example" // string | File name of file to load

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsLoadBackup(context.Background()).Backupfile(backupfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsLoadBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsLoadBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsLoadBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsLoadBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupfile** | **string** | File name of file to load | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsMoviesreport

> map[string]interface{} GetUserUsageStatsMoviesreport(ctx).UserId(userId).Days(days).EndDate(endDate).Execute()

Gets Movies counts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	userId := "userId_example" // string | User Id (optional)
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsMoviesreport(context.Background()).UserId(userId).Days(days).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsMoviesreport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsMoviesreport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsMoviesreport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsMoviesreportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsPlayactivity

> map[string]interface{} GetUserUsageStatsPlayactivity(ctx).Days(days).EndDate(endDate).Filter(filter).DataType(dataType).Execute()

Gets play activity for number of days



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)
	filter := "filter_example" // string | Comma separated list of media types to filter (movies,series) (optional)
	dataType := "dataType_example" // string | Data type to return (count,time) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsPlayactivity(context.Background()).Days(days).EndDate(endDate).Filter(filter).DataType(dataType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsPlayactivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsPlayactivity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsPlayactivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsPlayactivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 
 **filter** | **string** | Comma separated list of media types to filter (movies,series) | 
 **dataType** | **string** | Data type to return (count,time) | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsSaveBackup

> map[string]interface{} GetUserUsageStatsSaveBackup(ctx).Execute()

Saves a backup of the playback report data to the backup path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsSaveBackup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsSaveBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsSaveBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsSaveBackup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsSaveBackupRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsSessionList

> map[string]interface{} GetUserUsageStatsSessionList(ctx).Execute()

Gets Session Info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsSessionList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsSessionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsSessionList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsSessionList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsSessionListRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsTvshowsreport

> map[string]interface{} GetUserUsageStatsTvshowsreport(ctx).UserId(userId).Days(days).EndDate(endDate).Execute()

Gets TV Shows counts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	userId := "userId_example" // string | User Id (optional)
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsTvshowsreport(context.Background()).UserId(userId).Days(days).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsTvshowsreport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsTvshowsreport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsTvshowsreport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsTvshowsreportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsTypeFilterList

> map[string]interface{} GetUserUsageStatsTypeFilterList(ctx).Execute()

Gets types filter list items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsTypeFilterList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsTypeFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsTypeFilterList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsTypeFilterList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsTypeFilterListRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsUserActivity

> map[string]interface{} GetUserUsageStatsUserActivity(ctx).Days(days).EndDate(endDate).Execute()

Gets a report of the available activity per hour



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserActivity(context.Background()).Days(days).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsUserActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsUserActivity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsUserActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsUserActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsUserList

> map[string]interface{} GetUserUsageStatsUserList(ctx).Execute()

Get users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsUserList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsUserList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsUserListRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsUserManageByActionById

> map[string]interface{} GetUserUsageStatsUserManageByActionById(ctx, action, id).Execute()

Get users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	action := "action_example" // string | action to perform
	id := "id_example" // string | user Id to perform the action on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserManageByActionById(context.Background(), action, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsUserManageByActionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsUserManageByActionById`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsUserManageByActionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string** | action to perform | 
**id** | **string** | user Id to perform the action on | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsUserManageByActionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsageStatsUserplaylist

> map[string]interface{} GetUserUsageStatsUserplaylist(ctx).UserId(userId).AggregateData(aggregateData).FilterName(filterName).Days(days).EndDate(endDate).Filter(filter).Execute()

Gets a report of all played items for a user in a date period



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	userId := "userId_example" // string | User Id
	aggregateData := true // bool | Aggregate the data to total duration per user per item
	filterName := "filterName_example" // string | Name Filter (optional)
	days := int32(56) // int32 | Number of Days (optional)
	endDate := "endDate_example" // string | End date of the report in yyyy-MM-dd format (optional)
	filter := "filter_example" // string | Comma separated list of media types to filter (movies,series) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserplaylist(context.Background()).UserId(userId).AggregateData(aggregateData).FilterName(filterName).Days(days).EndDate(endDate).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.GetUserUsageStatsUserplaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserUsageStatsUserplaylist`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.GetUserUsageStatsUserplaylist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsageStatsUserplaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **aggregateData** | **bool** | Aggregate the data to total duration per user per item | 
 **filterName** | **string** | Name Filter | 
 **days** | **int32** | Number of Days | 
 **endDate** | **string** | End date of the report in yyyy-MM-dd format | 
 **filter** | **string** | Comma separated list of media types to filter (movies,series) | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserUsageStatsImportBackup

> PostUserUsageStatsImportBackup(ctx).Body(body).Execute()

Post a backup for importing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserActivityAPIAPI.PostUserUsageStatsImportBackup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.PostUserUsageStatsImportBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserUsageStatsImportBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserUsageStatsSubmitCustomQuery

> map[string]interface{} PostUserUsageStatsSubmitCustomQuery(ctx).ModelPlaybackReportingApiCustomQuery(modelPlaybackReportingApiCustomQuery).Execute()

Submit an SQL query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	modelPlaybackReportingApiCustomQuery := *openapiclient.NewModelPlaybackReportingApiCustomQuery() // ModelPlaybackReportingApiCustomQuery | CustomQuery

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserActivityAPIAPI.PostUserUsageStatsSubmitCustomQuery(context.Background()).ModelPlaybackReportingApiCustomQuery(modelPlaybackReportingApiCustomQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserActivityAPIAPI.PostUserUsageStatsSubmitCustomQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUserUsageStatsSubmitCustomQuery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserActivityAPIAPI.PostUserUsageStatsSubmitCustomQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserUsageStatsSubmitCustomQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelPlaybackReportingApiCustomQuery** | [**ModelPlaybackReportingApiCustomQuery**](ModelPlaybackReportingApiCustomQuery.md) | CustomQuery | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

