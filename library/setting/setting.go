package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

// 注意: logging 包依赖 setting，不能在这里导入 logging，否则会循环依赖
// 这里使用标准 log 包

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string

	WanIP      string
	ServerName string
	XmLogPath  string
	CFGAreaUrl string

	AreaUrlCN string
	AreaUrlAS string
	AreaUrlEU string
	AreaUrlNA string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	TcpIp        string
	TcpPort      int
	UdpIp        string
	UdpPort      int
	HttpIp       string
	HttpPort     int
	HttpsPort    int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var MysqlSetting = &Mysql{}

type Mysql struct {
	Host         string
	Port         int
	Username     string
	Password     string
	Database     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

// Wechat 微信小程序配置
type Wechat struct {
	AppID       string
	AppSecret   string
	TokenExpire int    // token过期时间(小时)
	JwtSecret   string // JWT签名密钥
}

var WechatSetting = &Wechat{}

var cfg *ini.File

func Setup(env string) {
	var err error
	cfg, err = ini.Load("conf/" + env + ".ini")
	if err != nil {
		log.Printf("{\"err\":\"%s\",\"desc\":%s}", err, "setting.Setup, fail to parse 'conf/prod.ini'")
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("mysql", MysqlSetting)
	mapTo("wechat", WechatSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Printf("{\"err\":\"%s\",\"section\":\"%s\"}", err, section)
	}
}
