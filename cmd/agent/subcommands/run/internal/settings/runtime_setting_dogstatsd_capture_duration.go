// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package settings

import (
	"fmt"
	"time"

	pkgsettings "github.com/DataDog/datadog-agent/pkg/config/settings"
)

// DsdCaptureDurationRuntimeSetting wraps operations to change the duration, in seconds, of traffic captures
type DsdCaptureDurationRuntimeSetting struct {
	value  string
	source pkgsettings.Source
}

func NewDsdCaptureDurationRuntimeSetting(value string) *DsdCaptureDurationRuntimeSetting {
	return &DsdCaptureDurationRuntimeSetting{
		value:  value,
		source: pkgsettings.SourceDefault,
	}
}

// Description returns the runtime setting's description
func (l *DsdCaptureDurationRuntimeSetting) Description() string {
	return "Enable/disable dogstatsd traffic captures. Possible values are: start, stop"
}

// Hidden returns whether or not this setting is hidden from the list of runtime settings
func (l *DsdCaptureDurationRuntimeSetting) Hidden() bool {
	return false
}

// Name returns the name of the runtime setting
func (l *DsdCaptureDurationRuntimeSetting) Name() string {
	return l.value
}

// Get returns the current value of the runtime setting
func (l *DsdCaptureDurationRuntimeSetting) Get() (interface{}, error) {
	// TODO
	return 0, nil
}

// Set changes the value of the runtime setting
func (l *DsdCaptureDurationRuntimeSetting) Set(v interface{}, source pkgsettings.Source) error {
	var err error

	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%s.Set: Invalid data type", l.value)
	}

	_, err = time.ParseDuration(s)
	if err != nil {
		return fmt.Errorf("Unsupported type for %s: %v", l.value, err)
	}

	// TODO
	// common.DSD.Capture.SetDuration(d)
	l.source = source
	return nil
}

func (l *DsdCaptureDurationRuntimeSetting) GetSource() pkgsettings.Source {
	return l.source
}
