# glog
新日志包，采用全局logger变量，直接调用包方法即可


## 使用方式
`go get github.com/mao888/go-log`

以 gia 目前日志格式作为示范,目前 gia 日志的格式为:
```json
{
  "level":"info",
  "timestamp":"2021-11-16T14:27:43.178+0800",
  "label":"ss-dispatcher@v0.1.1/dispatcher.go:78",
  "message":"Method[POST] Path[gia_company_go/fbaccount/list] [FBAccountController] 新增fb账户信息",
  "service":"gia-company-go"
}
```
其中 zap 默认的 time、caller、msg 字段被替换为 timestamp、label、message ，
且添加了全局唯一的字段 service，使用 glog 的方式有两种，分别如下:
### 一、返回全局 logger 变量方式
```golang
import glog "github.com/mao888/go-log"

// 初始化全局logger
levelType := InfoLevel
globalFields := Fields{
    "service": "service_name",
}
Init(
	//打开控制台日志，默认关闭
    WithConsoleStdout(),
	//默认 level 为 info
	WithLevel(levelType),
	//设置关闭自动压缩文件，默认打开
	WithOffCompress(),
	//日志文件位置，默认 ./log.log
	WithFileLocation("test.log"),
	// 设置日志保存天数，默认30
	WithLogMaxAge(30),
	//设置最大文件大小（MB），默认256
	WithLogMaxSize(250),
	//设置全局自定义字段
	WithCustomizedGlobalField(globalFields),
	//设置覆盖默认字段
	WithCoverDefaultKey(CoverDefaultKey{
	TimeKey:    "timestamp",
	CallerKey:  "label",
	MessageKey: "message"}),
)

// 日志打印
glog.C(ctx).Debug("test debug")
glog.C(ctx).Infof("test: %s","info")

// 也支持打印时新加字段，但仅影响本次调用，不会影响全局字段，仅支持打印 info 日志
glog.C(ctx).InfoWithField(map[string]interface{}{
	"temp_field":"glog is good "
}, "msg1","msg2")

```

### 二、直接调用包方法
```go
//初始化流程如上，此处不再重复

// 日志打印
glog.Debug(ctx,"test debug")
glog.Infof(ctx,"test: %s","info")

// 也支持打印时新加字段，但仅影响本次调用，不会影响全局字段，仅支持打印 info 日志
glog.InfoWithField(ctx,map[string]interface{}{
"temp_field":"glog is good "
}, "msg1","msg2")

```

## 性能

基本相比原日志包提升 10% 左右

goos: darwin

goarch: amd64

cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz

打印两个全局变量加一个msg变量:

| package | Time | Mem Per op| MemAlloc |
| :---- | :---- | :---- |:---- |
| ftlog | ~10000 ns/op   | 1962 B/op  | 17 allocs/op |
| glog  | ~9000 ns/op    | 1858 B/op  | 12 allocs/op |



## 版本

| 版本     | 修订说明                                             | 提交人 | 状态 | 起止时间           |
|--------|--------------------------------------------------|----| ----- |----------------|
| v1.0.0 | 第一版 logger                                       | 胡超 | 开发完成 | ~ 至 2022-11-09 |
| v1.0.1 | 优化日志性能，全局字段并发安全化                                 | 胡超 | 开发完成 | ~ 至 2022-11-12 |
| v1.0.2 | 未初始化时调用静默失败处理 <br> 日志结构体克隆优化 <br> 添加控制台输出 option | 胡超 | 开发完成 | 2022-11-17     |
| v1.0.3 | 默认字段覆盖                                           | 胡超 | 开发完成 | 2021-12-17     |
| v1.0.4 | 修复 bug                                           | 胡超 | 开发完成 |                |
| v1.0.5 | 修复 panic，panicf 方法                               | 胡超 | 开发完成 | 2023-01-17     |
| v1.0.6 | 增加打印 "s_time" 字段方法 Time                          | 胡超 | 开发完成 | 2023-01-19     |





