# ModelAccessSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**ModelModelDynamicDayOfWeek**](ModelDynamicDayOfWeek.md) |  | [optional] 
**StartHour** | Pointer to **float64** |  | [optional] 
**EndHour** | Pointer to **float64** |  | [optional] 

## Methods

### NewModelAccessSchedule

`func NewModelAccessSchedule() *ModelAccessSchedule`

NewModelAccessSchedule instantiates a new ModelAccessSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAccessScheduleWithDefaults

`func NewModelAccessScheduleWithDefaults() *ModelAccessSchedule`

NewModelAccessScheduleWithDefaults instantiates a new ModelAccessSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *ModelAccessSchedule) GetDayOfWeek() ModelModelDynamicDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *ModelAccessSchedule) GetDayOfWeekOk() (*ModelModelDynamicDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *ModelAccessSchedule) SetDayOfWeek(v ModelModelDynamicDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *ModelAccessSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetStartHour

`func (o *ModelAccessSchedule) GetStartHour() float64`

GetStartHour returns the StartHour field if non-nil, zero value otherwise.

### GetStartHourOk

`func (o *ModelAccessSchedule) GetStartHourOk() (*float64, bool)`

GetStartHourOk returns a tuple with the StartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHour

`func (o *ModelAccessSchedule) SetStartHour(v float64)`

SetStartHour sets StartHour field to given value.

### HasStartHour

`func (o *ModelAccessSchedule) HasStartHour() bool`

HasStartHour returns a boolean if a field has been set.

### GetEndHour

`func (o *ModelAccessSchedule) GetEndHour() float64`

GetEndHour returns the EndHour field if non-nil, zero value otherwise.

### GetEndHourOk

`func (o *ModelAccessSchedule) GetEndHourOk() (*float64, bool)`

GetEndHourOk returns a tuple with the EndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHour

`func (o *ModelAccessSchedule) SetEndHour(v float64)`

SetEndHour sets EndHour field to given value.

### HasEndHour

`func (o *ModelAccessSchedule) HasEndHour() bool`

HasEndHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


