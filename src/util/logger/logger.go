package logger

import (
	"os"
	"gin-use/configs"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logConf *LogConfig
var logConfOnce = new(sync.Once)
var logger *zap.SugaredLogger
var loggerOnce = new(sync.Once)

// writer type
const (
	WriterTypStdou string = "stdout"
	WriterTypFile  string = "file"
)

// encoder type
const (
	EncoderTypConsole string = "console"
	EncoderTypJson    string = "json"
)

type LogConfig struct {
	WriterTypes []string `json:"writer_types"` // stdout, file
	Lvl         string   `json:"lvl"`          // debug, info, warn, error, dpanic, panic, fatal
	Encoding    string   `json:"encoding"`     // console, json
	Filename    string   `json:"filename"`
	MaxSizeMB   int      `json:"max_size_mb"`
	MaxAgeDay   int      `json:"max_age_day"`
	MaxBackup   int      `json:"max_backup"`
	IsLocalTime bool     `json:"is_local_time"`
	IsCompress  bool     `json:"is_compress"`
}

func Log() *zap.SugaredLogger {
	loggerOnce.Do(func() {
		// 打印机
		var writer zapcore.WriteSyncer
		writers := []zapcore.WriteSyncer{}
		for _, typ := range LogConf().WriterTypes {
			switch typ {
			case WriterTypStdou:
				writers = append(writers, os.Stdout)
			case WriterTypFile:
				writers = append(writers, zapcore.AddSync(&lumberjack.Logger{
					Filename:   LogConf().Filename,
					MaxSize:    LogConf().MaxSizeMB,
					MaxAge:     LogConf().MaxAgeDay,
					MaxBackups: LogConf().MaxBackup,
					LocalTime:  LogConf().IsLocalTime,
					Compress:   LogConf().IsCompress,
				}))
			}
		}
		writer = zapcore.NewMultiWriteSyncer(writers...)

		// 编码器
		var encoder zapcore.Encoder
		encCfg := zap.NewProductionEncoderConfig()
		encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encCfg.EncodeCaller = zapcore.ShortCallerEncoder
		switch LogConf().Encoding {
		case EncoderTypConsole:
			encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // 开启颜色
			encoder = zapcore.NewConsoleEncoder(encCfg)
		case EncoderTypJson:
			encoder = zapcore.NewJSONEncoder(encCfg)
		default:
			encoder = zapcore.NewJSONEncoder(encCfg)
		}

		// 最小等级
		var lvl zapcore.Level
		if err := lvl.Set(LogConf().Lvl); err != nil {
			lvl = zapcore.DebugLevel
		}

		core := zapcore.NewCore(encoder, writer, lvl)
		logger = zap.New(core, zap.AddCaller()).Sugar()
	})
	return logger
}

func LogConf() *LogConfig {
	logConfOnce.Do(func() {
		logConf = &LogConfig{
			WriterTypes: []string{"stdout", "file"},
			Lvl:         "debug",
			Encoding:    "json",
			Filename:    configs.ProjectLogFile(), // 文件路径
			MaxSizeMB:   128,                      // 单个文件最大尺寸，默认单位 M
			MaxAgeDay:   10,					   // 最大时间，默认单位 day
			MaxBackup:   300,					   // 最多保留 300 个备份
			IsCompress:  true,					   // 是否压缩 disabled by default
			IsLocalTime: true,					   // 使用本地时间
		}
	})
	return logConf
}
