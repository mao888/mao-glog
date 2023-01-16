package glog

import (
	"context"
	"testing"
)

func init() {
	_ = Init(
		//WithConsoleStdout(),
		//WithConsoleLevel(DebugLevel),
		//WithConsoleFormatJson(),
		WithLevel(DebugLevel),
		//WithOffCompress(),
		WithFileLocation("test.log"),
		//WithLogMaxAge(30),
		//WithLogMaxSize(250),
		WithCustomizedGlobalField(map[string]interface{}{"s_code": "51400"}),
		//WithCoverDefaultKey(CoverDefaultKey{
		//	TimeKey:       "timestamp",
		//	CallerKey:     "label",
		//	MessageKey:    "message",
		//	StacktraceKey: "stack"},
		//),
	)
}

func TestInit(t *testing.T) {
	//vctx := context.Background()
	vctx := context.WithValue(context.Background(), TrackKey, "12345")
	//C(vctx).Debug("Package Debug")
	//C(vctx).Debugf("Package Debugf: %s", "Debugf")
	//C(vctx).Info("Package Info")
	//C(vctx).Infof("Package Infof: %s", "Infof")
	//C(vctx).Warn("Package Warn")
	//C(vctx).Warnf("Package Warnf: %s", "Warnf")
	//C(vctx).Error("Package Error")
	//C(vctx).Errorf("Package Errorf: %s", "Errorf")
	//C(vctx).InfoWithField(map[string]interface{}{
	//	"testField": 12345,
	//}, "yeyey")
	//isDebug := IsDebug()
	//C(vctx).Info("Package is debug ", isDebug)
	//ChangeFileStdoutLevel(InfoLevel)
	//isDebug = IsDebug()
	//C(vctx).Info("after ChangeFileStdoutLevel is debug ", isDebug)
	Time(0.03).Errorf("test with time field: %.2f", 0.02)
	Time(0.03).Infof("test with time field: %.2f", 0.03)
	C(vctx).Warnf("test with time field: %.2f", 0.04)
	C(vctx).Time(0.05).Warnf("test with time field: %.2f", 0.05)
}

func TestFunctionalCall(t *testing.T) {
	//vctx := context.Background()
	vctx := context.WithValue(context.Background(), TrackKey, "12345")
	Debug(vctx, "Package Debug")
	Debugf(vctx, "Package Debugf: %s", "Debugf")
	Info(vctx, "Package Info")
	Infof(vctx, "Package Infof: %s", "Infof")
	Warn(vctx, "Package Warn")
	Warnf(vctx, "Package Warnf: %s", "Warnf")
	Error(vctx, "Package Error")
	Errorf(vctx, "Package Errorf: %s", "Errorf")
	InfoWithField(vctx, map[string]interface{}{
		"testField": 12345,
	}, "yeyey")
	isDebug := IsDebug()
	Info(vctx, "Package is debug ", isDebug)
	ChangeFileStdoutLevel(InfoLevel)
	isDebug = IsDebug()
	Info(vctx, "after ChangeFileStdoutLevel is debug ", isDebug)
	Panic(vctx, "test panic ")
}

func BenchmarkInit(b *testing.B) {
	//ctx := context.Background()
	vctx := context.WithValue(context.Background(), TrackKey, "1234")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//log.Printf("test Debugf:%v", "error")
		//Debug(vctx, "Debug")
		//Debugf(vctx, "test Debugf: %v", "error")
		//Info(vctx, "Info")
		Infof(vctx, "test Infof: %v", "error")
		//Warn(vctx, "Warn")
		//C(vctx).Warnf("test Warnf: %v", "error")
		//Error(vctx, "Error")
		//C(vctx).Errorf("test Infof: %v", "Infof")
		//Panic(vctx, "Panic")
		//Panicf(vctx, "test Panicf: %v", "error")
		//InfoWithField(vctx, nil, "yeyey")
	}
}
