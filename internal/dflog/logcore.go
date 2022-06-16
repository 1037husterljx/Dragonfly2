/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	"strings"

	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	CoreLogFileName       = "core.log"
	GrpcLogFileName       = "grpc.log"
	GCLogFileName         = "gc.log"
	StorageGCLogFileName  = "storage-gc.log"
	JobLogFileName        = "job.log"
	StatSeedLogFileName   = "stat/seed.log"
	DownloaderLogFileName = "downloader.log"
	KeepAliveLogFileName  = "keepalive.log"
	SqlLogFileName        = "sql.log"
)

const (
	defaultRotateMaxSize    = 200
	defaultRotateMaxBackups = 10
	defaultRotateMaxAge     = 7
)

const (
	encodeTimeFormat = "2006-01-02 15:04:05.000"
)

var DefaultLogConfig = LogConfig{
	MaxSize:    defaultRotateMaxSize,
	MaxBackups: defaultRotateMaxBackups,
	MaxAge:     defaultRotateMaxAge,
	Compress:   false,
	Structural: true,
}

var coreLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
var customCoreLevel atomic.Bool
var grpcLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
var customGrpcLevel atomic.Bool

type LogConfig struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets rotated.
	// It defaults to 40 megabytes.
	MaxSize int `yaml:"maxSize" mapstructure:"maxSize"`
	// MaxBackups is the maximum number of old log files to retain.
	// The default value is 1.
	MaxBackups int `yaml:"maxBackups" mapstructure:"maxBackups"`
	// MaxAge is the maximum number of days log files to retain.
	// The default value is 7
	MaxAge int `yaml:"maxAge" mapstructure:"maxAge"`
	// Compress is the option of compress log data
	Compress bool `yaml:"compress" mapstructure:"compress"`
	// Structural is the option whether log using json-formatter or console-formatter
	Structural bool `yaml:"structural" mapstructure:"structural"`
}

type LogConfigs struct {
	Core       LogConfig `yaml:"core" mapstructure:"core"`
	Grpc       LogConfig `yaml:"grpc" mapstructure:"grpc"`
	GC         LogConfig `yaml:"gC" mapstructure:"gc"`
	StorageGc  LogConfig `yaml:"storageGc" mapstructure:"storageGc"`
	Job        LogConfig `yaml:"job" mapstructure:"job"`
	Sql        LogConfig `yaml:"sql" mapstructure:"sql"`
	StatSeed   LogConfig `yaml:"statSeed" mapstructure:"statSeed"`
	Downloader LogConfig `yaml:"downloader" mapstructure:"downloader"`
	KeepAlive  LogConfig `yaml:"keepAlive" mapstructure:"keepAlive"`
}

func CreateLogger(filePath string, stats bool, verbose bool, cfg LogConfig) (*zap.Logger, zap.AtomicLevel, error) {
	rotateConfig := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  true,
		Compress:   cfg.Compress,
	}
	syncer := zapcore.AddSync(rotateConfig)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(encodeTimeFormat)
	var level = zap.NewAtomicLevel()
	if verbose {
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	if strings.HasSuffix(filePath, GrpcLogFileName) && customGrpcLevel.Load() {
		level = grpcLevel
	} else if strings.HasSuffix(filePath, CoreLogFileName) && customCoreLevel.Load() {
		level = coreLevel
	}

	var encoder zapcore.Encoder
	if cfg.Structural {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		encoder,
		syncer,
		level,
	)

	var opts []zap.Option
	if !stats {
		opts = append(opts, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel), zap.AddCallerSkip(1))
	}

	return zap.New(core, opts...), level, nil
}

func NewDefaultConfigs() *LogConfigs {
	return &LogConfigs{
		Core:       DefaultLogConfig,
		Grpc:       DefaultLogConfig,
		GC:         DefaultLogConfig,
		StorageGc:  DefaultLogConfig,
		Job:        DefaultLogConfig,
		Sql:        DefaultLogConfig,
		StatSeed:   DefaultLogConfig,
		Downloader: DefaultLogConfig,
		KeepAlive:  DefaultLogConfig,
	}
}

type Option func(o *LogConfig) error

func WithMaxSize(maxSize int) Option {
	return func(o *LogConfig) error {
		o.MaxSize = maxSize
		return nil
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(o *LogConfig) error {
		o.MaxBackups = maxBackups
		return nil
	}
}

func WithMaxAge(maxAge int) Option {
	return func(o *LogConfig) error {
		o.MaxAge = maxAge
		return nil
	}
}

func WithCompress(compress bool) Option {
	return func(o *LogConfig) error {
		o.Compress = compress
		return nil
	}
}

func WithStructural(structural bool) Option {
	return func(o *LogConfig) error {
		o.Structural = structural
		return nil
	}
}

func SetCoreLevel(level zapcore.Level) {
	customCoreLevel.Store(true)
	coreLevel.SetLevel(level)
}

func SetGrpcLevel(level zapcore.Level) {
	customGrpcLevel.Store(true)
	grpcLevel.SetLevel(level)
}
