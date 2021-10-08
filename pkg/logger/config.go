// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package logger

import (
	"os"
)

const (
	defaultLogFilePath  = "/var/log/go-west/access.log"
	defaultLogLevel     = "Debug"
	defaultLogFilePath2 = "/var/log/go-west/error.log"
	defaultLogLevel2    = "ERROR"
	envLogLevel         = "AWS_VPC_K8S_CNI_LOGLEVEL"
	envLogFilePath      = "AWS_VPC_K8S_CNI_LOG_FILE"
)

// Configuration stores the config for the logger
type Configuration struct {
	LogLevel     string
	LogLevel2    string
	LogLocation  string
	LogLocation2 string
}

// LoadLogConfig returns the log configuration
func LoadLogConfig() *Configuration {
	return &Configuration{
		LogLevel:     GetLogLevel(),
		LogLevel2:    GetLogLevel2(),
		LogLocation:  GetLogLocation(),
		LogLocation2: GetLogLocation2(),
	}
}

// GetLogLocation returns the log file path
func GetLogLocation() string {
	logFilePath := os.Getenv(envLogFilePath)
	if logFilePath == "" {
		logFilePath = defaultLogFilePath
	}
	return logFilePath
}

// GetLogLevel returns the log level
func GetLogLevel() string {
	logLevel := os.Getenv(envLogLevel)
	switch logLevel {
	case "":
		logLevel = defaultLogLevel
		return logLevel
	default:
		return logLevel
	}
}

// GetLogLocation2 GetLogLocation returns the log file path
func GetLogLocation2() string {
	logFilePath := os.Getenv(envLogFilePath)
	if logFilePath == "" {
		logFilePath = defaultLogFilePath2
	}
	return logFilePath
}

// GetLogLevel2 GetLogLevel returns the log level
func GetLogLevel2() string {
	logLevel := os.Getenv(envLogLevel)
	switch logLevel {
	case "":
		logLevel = defaultLogLevel2
		return logLevel
	default:
		return logLevel
	}
}
