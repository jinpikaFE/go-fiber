package setting

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	SecretId  string
	SecretKey string
	CosUrl    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logging.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadTenCentCos()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("dev")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logging.Fatal("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logging.Fatal("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadTenCentCos() {
	sec, err := Cfg.GetSection("tencent_cos")
	if err != nil {
		logging.Fatal("Fail to get section 'tencent_cos': %v", err)
	}

	SecretId = sec.Key("SECRET_ID").MustString("")
	SecretKey = sec.Key("SECRET_KEY").MustString("")
	CosUrl = sec.Key("COS_URL").MustString("")
}
