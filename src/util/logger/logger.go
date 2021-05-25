package log

import (
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)


const (
    // logFormat
    LOGFORMAT_JSON    = "json"
    LOGFORMAT_CONSOLE = "console"

    // EncoderConfig
    TIME_KEY       = "time"
    LEVLE_KEY      = "level"
    NAME_KEY       = "logger"
    CALLER_KEY     = "caller"
    MESSAGE_KEY    = "msg"
    STACKTRACE_KEY = "stacktrace"

    // 日志归档配置项
    // 每个日志文件保存的最大尺寸 单位：M
    MAX_SIZE = 1
    // 文件最多保存多少天
    MAX_BACKUPS = 5
    // 日志文件最多保存多少个备份
    MAX_AGE = 7
)


// 设置日志级别、输出格式和日志文件的路径
func SetLogs(logLevel zapcore.Level, logFormat, fileName string) {

    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        TIME_KEY,
        LevelKey:       LEVLE_KEY,
        NameKey:        NAME_KEY,
        CallerKey:      CALLER_KEY,
        MessageKey:     MESSAGE_KEY,
        StacktraceKey:  STACKTRACE_KEY,
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器(相对路径+行号)
        EncodeName:     zapcore.FullNameEncoder,
    }

    // 设置日志输出格式
    var encoder zapcore.Encoder
    switch logFormat {
    case LOGFORMAT_JSON:
        encoder = zapcore.NewJSONEncoder(encoderConfig)
    default:
        encoder = zapcore.NewConsoleEncoder(encoderConfig)
    }

    // 添加日志切割归档功能
    hook := lumberjack.Logger{
        Filename:   fileName,    // 日志文件路径
        MaxSize:    MAX_SIZE,    // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: MAX_BACKUPS, // 日志文件最多保存多少个备份
        MaxAge:     MAX_AGE,     // 文件最多保存多少天
        Compress:   true,        // 是否压缩
    }

    core := zapcore.NewCore(
        encoder, // 编码器配置
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)), // 打印到控制台和文件
        zap.NewAtomicLevelAt(logLevel), // 日志级别
    )

    // 开启文件及行号
    caller := zap.AddCaller()
    // 开启开发模式，堆栈跟踪
    development := zap.Development()
    // 构造日志
    logger := zap.New(core, caller, development)
   
    // 将自定义的logger替换为全局的logger
    zap.ReplaceGlobals(logger)
}