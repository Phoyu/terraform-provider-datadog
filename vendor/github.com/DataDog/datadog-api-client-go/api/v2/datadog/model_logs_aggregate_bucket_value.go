/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// LogsAggregateBucketValue - A bucket value, can be either a timeseries or a single value
type LogsAggregateBucketValue struct {
	LogsAggregateBucketValueTimeseries *LogsAggregateBucketValueTimeseries
	float64                            *float64
	string                             *string
}

// LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue is a convenience function that returns LogsAggregateBucketValueTimeseries wrapped in LogsAggregateBucketValue
func LogsAggregateBucketValueTimeseriesAsLogsAggregateBucketValue(v *LogsAggregateBucketValueTimeseries) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{LogsAggregateBucketValueTimeseries: v}
}

// float64AsLogsAggregateBucketValue is a convenience function that returns float64 wrapped in LogsAggregateBucketValue
func float64AsLogsAggregateBucketValue(v *float64) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{float64: v}
}

// stringAsLogsAggregateBucketValue is a convenience function that returns string wrapped in LogsAggregateBucketValue
func stringAsLogsAggregateBucketValue(v *string) LogsAggregateBucketValue {
	return LogsAggregateBucketValue{string: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsAggregateBucketValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsAggregateBucketValueTimeseries
	err = json.Unmarshal(data, &dst.LogsAggregateBucketValueTimeseries)
	if err == nil {
		jsonLogsAggregateBucketValueTimeseries, _ := json.Marshal(dst.LogsAggregateBucketValueTimeseries)
		if string(jsonLogsAggregateBucketValueTimeseries) == "{}" { // empty struct
			dst.LogsAggregateBucketValueTimeseries = nil
		} else {
			match++
		}
	} else {
		dst.LogsAggregateBucketValueTimeseries = nil
	}

	// try to unmarshal data into float64
	err = json.Unmarshal(data, &dst.float64)
	if err == nil {
		jsonfloat64, _ := json.Marshal(dst.float64)
		if string(jsonfloat64) == "{}" { // empty struct
			dst.float64 = nil
		} else {
			match++
		}
	} else {
		dst.float64 = nil
	}

	// try to unmarshal data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			match++
		}
	} else {
		dst.string = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LogsAggregateBucketValueTimeseries = nil
		dst.float64 = nil
		dst.string = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(LogsAggregateBucketValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(LogsAggregateBucketValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsAggregateBucketValue) MarshalJSON() ([]byte, error) {
	if src.LogsAggregateBucketValueTimeseries != nil {
		return json.Marshal(&src.LogsAggregateBucketValueTimeseries)
	}

	if src.float64 != nil {
		return json.Marshal(&src.float64)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsAggregateBucketValue) GetActualInstance() interface{} {
	if obj.LogsAggregateBucketValueTimeseries != nil {
		return obj.LogsAggregateBucketValueTimeseries
	}

	if obj.float64 != nil {
		return obj.float64
	}

	if obj.string != nil {
		return obj.string
	}

	// all schemas are nil
	return nil
}

type NullableLogsAggregateBucketValue struct {
	value *LogsAggregateBucketValue
	isSet bool
}

func (v NullableLogsAggregateBucketValue) Get() *LogsAggregateBucketValue {
	return v.value
}

func (v *NullableLogsAggregateBucketValue) Set(val *LogsAggregateBucketValue) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsAggregateBucketValue) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsAggregateBucketValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsAggregateBucketValue(val *LogsAggregateBucketValue) *NullableLogsAggregateBucketValue {
	return &NullableLogsAggregateBucketValue{value: val, isSet: true}
}

func (v NullableLogsAggregateBucketValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsAggregateBucketValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}