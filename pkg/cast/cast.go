// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast provides easy and safe casting in Go.

// Forked from github.com/spf13/cast
// Some parts of this file has been modified to make it functional in this package

// Package cast provides the typeCasting of an object
package cast

import (
	ca "github.com/spf13/cast"
)

type configValue struct {
	value interface{}
	err   error
}

// NewValue creates an object for an interface X.
func NewValue(val interface{}, err error) Value {
	confVal := new(configValue)
	confVal.value = val
	confVal.err = err
	return confVal
}

// Value is an interface to typecast an Object.
type Value interface {
	ToInt64() (int64, error)
	ToInt32() (int32, error)
	ToInt16() (int16, error)
	ToInt8() (int8, error)
	ToInt() (int, error)
	ToUint() (uint, error)
	ToUint64() (uint64, error)
	ToUint32() (uint32, error)
	ToUint16() (uint16, error)
	ToUint8() (uint8, error)
	ToString() (string, error)
	ToStringMapStringSlice() (map[string][]string, error)
	ToStringMapBool() (map[string]bool, error)
	ToStringMap() (map[string]interface{}, error)
	ToSlice() ([]interface{}, error)
	ToBoolSlice() ([]bool, error)
	ToStringSlice() ([]string, error)
	ToIntSlice() ([]int, error)
	ToBool() (bool, error)
	ToFloat64() (float64, error)
}

func (val *configValue) ToInt64() (int64, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToInt64E(val.value)
}

func (val *configValue) ToInt32() (int32, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToInt32E(val.value)
}

func (val *configValue) ToInt16() (int16, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToInt16E(val.value)
}

func (val *configValue) ToInt8() (int8, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToInt8E(val.value)
}

func (val *configValue) ToInt() (int, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToIntE(val.value)
}

func (val *configValue) ToUint() (uint, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToUintE(val.value)
}

func (val *configValue) ToUint64() (uint64, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToUint64E(val.value)
}

func (val *configValue) ToUint32() (uint32, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToUint32E(val.value)
}

func (val *configValue) ToUint16() (uint16, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToUint16E(val.value)
}

func (val *configValue) ToUint8() (uint8, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToUint8E(val.value)
}

func (val *configValue) ToString() (string, error) {
	if val.err != nil {
		return "", val.err
	}

	return ca.ToStringE(val.value)
}

func (val *configValue) ToStringMapStringSlice() (map[string][]string, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToStringMapStringSliceE(val.value)
}

func (val *configValue) ToStringMapBool() (map[string]bool, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToStringMapBoolE(val.value)
}

func (val *configValue) ToStringMap() (map[string]interface{}, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToStringMapE(val.value)
}

func (val *configValue) ToSlice() ([]interface{}, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToSliceE(val.value)
}

func (val *configValue) ToBoolSlice() ([]bool, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToBoolSliceE(val.value)
}

func (val *configValue) ToStringSlice() ([]string, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToStringSliceE(val.value)
}

func (val *configValue) ToIntSlice() ([]int, error) {
	if val.err != nil {
		return nil, val.err
	}

	return ca.ToIntSliceE(val.value)
}

func (val *configValue) ToBool() (bool, error) {
	if val.err != nil {
		return false, val.err
	}

	return ca.ToBoolE(val.value)
}

func (val *configValue) ToFloat64() (float64, error) {
	if val.err != nil {
		return 0, val.err
	}

	return ca.ToFloat64E(val.value)
}
