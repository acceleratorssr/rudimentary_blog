package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"server/global"
	"time"
)

const (
	red    = 31
	green  = 32
	yellow = 33
	gray   = 37
)

type logFileWriter struct {
	file     *os.File
	fileDate string //判断日期切换目录
	appName  string
}

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = green
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		_, err := fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", global.Config.Logger.Prefix, timestamp, levelColor, entry.Level, funcVal, fileVal, entry.Message)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", global.Config.Logger.Prefix, timestamp, levelColor, entry.Level, entry.Message)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}

	//判断是否需要切换日期
	fileDate := time.Now().Format("2006-01-02")
	if p.fileDate != fileDate {
		//err := p.file.Close()
		//if err != nil {
		//	return 0, err
		//}

		//err = os.MkdirAll(fmt.Sprintf("%s/%s", global.Config.Logger.Director, fileDate), os.ModePerm)
		//if err != nil {
		//	return 0, err
		//}

		filename := fmt.Sprintf("%s/%s.log", global.Config.Logger.Director, fileDate)

		p.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			return 0, err
		}
	}

	n, e := p.file.Write(data)
	return n, e
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()

	fileDate := time.Now().Format("20060102")
	//// 创建目录
	//// 目录内文件所有者、组成员和其他用户都具有读、写和执行权限。
	//err := os.MkdirAll(fmt.Sprintf("%s/%s", global.Config.Logger.Director, fileDate), os.ModePerm)
	//if err != nil {
	//	logrus.Error(err)
	//	return nil
	//}

	// 构建日志路径
	filename := fmt.Sprintf("%s/%s.log", global.Config.Logger.Director, fileDate)
	// 打开/创建文件
	// 只写模式打开文件；在文件末尾追加数据；如果文件不存在，则创建文件
	// 只有文件所有者有读写权限
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	// 输出到文件和终端上
	writes := []io.Writer{
		file,
		os.Stdout,
	}
	mLog.SetOutput(io.MultiWriter(writes...))

	mLog.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
		return nil
	}
	mLog.SetLevel(level)

	mLog.SetReportCaller(global.Config.Logger.ShowLine)
	InitDefaultLogger()
	return mLog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)

	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
		return
	}
	logrus.SetLevel(level)

	logrus.SetLevel(logrus.DebugLevel)
}
