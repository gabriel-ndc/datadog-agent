// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package gohai

import (
	"net"
	"sync"

	"github.com/DataDog/datadog-agent/pkg/gohai/cpu"
	"github.com/DataDog/datadog-agent/pkg/gohai/filesystem"
	"github.com/DataDog/datadog-agent/pkg/gohai/memory"
	"github.com/DataDog/datadog-agent/pkg/gohai/network"
	"github.com/DataDog/datadog-agent/pkg/gohai/platform"

	"github.com/DataDog/datadog-agent/pkg/config"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

var (
	// we can use this a hint that docker is running in host mode and it's safe to use detect
	docker0Detected = false
)

// GetPayload builds a payload of every metadata collected with gohai except processes metadata.
func GetPayload() *Payload {
	return &Payload{
		Gohai: getGohaiInfo(),
	}
}

func getGohaiInfo() *gohai {
	res := new(gohai)

	cpuPayload, _, err := cpu.CollectInfo().AsJSON()
	if err == nil {
		res.CPU = cpuPayload
	} else {
		log.Errorf("Failed to retrieve cpu metadata: %s", err)
	}

	var fileSystemPayload interface{}
	fileSystemInfo, err := filesystem.CollectInfo()
	if err == nil {
		fileSystemPayload, _, err = fileSystemInfo.AsJSON()
	}
	if err == nil {
		res.FileSystem = fileSystemPayload
	} else {
		log.Errorf("Failed to retrieve filesystem metadata: %s", err)
	}

	memoryPayload, _, err := memory.CollectInfo().AsJSON()
	if err == nil {
		res.Memory = memoryPayload
	} else {
		log.Errorf("Failed to retrieve memory metadata: %s", err)
	}

	if !config.IsContainerized() || detectDocker0() {
		var networkPayload interface{}
		networkInfo, err := network.CollectInfo()
		if err == nil {
			networkPayload, _, err = networkInfo.AsJSON()
		}
		if err == nil {
			res.Network = networkPayload
		} else {
			log.Errorf("Failed to retrieve network metadata: %s", err)
		}
	}

	platformPayload, _, err := platform.CollectInfo().AsJSON()
	if err == nil {
		res.Platform = platformPayload
	} else {
		log.Errorf("Failed to retrieve platform metadata: %s", err)
	}

	return res
}

var docker0Detector sync.Once

func detectDocker0() bool {
	docker0Detector.Do(func() {
		iface, _ := net.InterfaceByName("docker0")
		docker0Detected = iface != nil
	})

	return docker0Detected
}
