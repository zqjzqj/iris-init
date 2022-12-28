package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm/logger"
	"iris-init/logs"
	"iris-init/orm"
	"iris-init/sErr"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	EnvDev  = "dev"
	EnvPro  = "pro"
	EnvTest = "test"
)

var cfg *Cfg

func init() {
	cfg = &Cfg{}
}

type Cfg struct {
	env           string
	runDaemon     bool
	runBackground bool
	pprofPort     int
	logs          LogsCfg
	web           Web
}

type LogsCfg struct {
	IsPrint     bool
	LogFilePath string
	LogFile     *os.File
}

func PprofPort() int {
	return cfg.pprofPort
}

func RunDaemon() bool {
	return cfg.runDaemon
}

func RunBackground() bool {
	return cfg.runBackground
}

func EnvIsDev() bool {
	return cfg.env == EnvDev
}

func EnvIsPro() bool {
	return cfg.env == EnvPro
}

func EnvIsTest() bool {
	return cfg.env == EnvTest
}

func GetEnv() string {
	return cfg.env
}

func GetWebCfg() Web {
	return cfg.web
}

func GetLogsCfg() LogsCfg {
	return cfg.logs
}

func LoadConfigJson(p string) error {
	logs.PrintlnInfo("reload config.....")
	defer logs.PrintlnSuccess("reload config success!")
	paths, fileName := filepath.Split(p)
	viper.SetConfigName(fileName)
	viper.AddConfigPath(paths)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg.env = viper.GetString("env")
	cfg.runDaemon = viper.GetBool("run_daemon")
	cfg.runBackground = viper.GetBool("run_background")
	cfg.pprofPort = viper.GetInt("pprof_port")
	//载入db配置
	if err = loadDbCfg(); err != nil {
		return err
	}
	//载入web配置
	if err = loadWebCfg(); err != nil {
		return err
	}
	//载入日志配置
	if err = loadLogsCfg(); err != nil {
		return err
	}
	return nil
}

func loadDbCfg() error {
	dbConfigs := viper.GetStringMap("db")
	isSetDef := false
	for key, v := range dbConfigs {
		b := make([]byte, 0)
		b, err := json.Marshal(v)
		if err != nil {
			return sErr.NewByError(err)
		}
		dbConf := gjson.ParseBytes(b)
		maxIdleCounts, _ := strconv.Atoi(dbConf.Get("max_idle_counts").String())
		charset := dbConf.Get("charset").String()
		if charset == "" {
			charset = "utf8"
		}
		loggerLevel := logger.Error
		if EnvIsDev() {
			loggerLevel = logger.Info
		}
		db, err := orm.NewDatabaseMysql(
			dbConf.Get("host").String(),
			dbConf.Get("port").String(),
			dbConf.Get("database").String(),
			charset,
			dbConf.Get("username").String(),
			dbConf.Get("password").String(),
			maxIdleCounts,
			int(dbConf.Get("max_open_counts").Int()),
			loggerLevel,
		)
		if err != nil {
			return sErr.NewByError(err)
		}
		if dbConf.Get("default").Bool() == true && !isSetDef {
			isSetDef = true
			orm.SetMysql("default", db, true)
		} else {
			orm.SetMysql(key, db, false)
		}
	}
	if orm.GetMysqlDef() == nil {
		iDB, ok := orm.GetMysql("c_customer")
		if !ok {
			return sErr.New("not default database")
		}
		orm.SetMysql("default", iDB, true)
	}
	return nil
}

func loadWebCfg() error {
	w := Web{}
	w.port = uint64(viper.GetInt64("web.port"))
	if w.port == 0 {
		w.port = 80
	}
	cfg.web = w
	return nil
}

func loadLogsCfg() error {
	var err error
	logCfg := LogsCfg{}
	logCfg.IsPrint = viper.GetBool("log.isPrint")
	logCfg.LogFilePath = viper.GetString("log.logFilePath")
	logs.IsPrintLog = logCfg.IsPrint
	if logCfg.LogFilePath != "" {
		logCfg.LogFile, err = os.OpenFile(logCfg.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		os.Stdout = logCfg.LogFile
		os.Stderr = logCfg.LogFile
		go func() {
			t := time.NewTicker(36 * time.Hour)
			defer t.Stop()
			log.Println("创建自动清除日志文件内容")
			for {
				select {
				case <-t.C:
					err = logCfg.LogFile.Truncate(0)
					if err != nil {
						log.Println("清空文件失败")
					}
					_, _ = logCfg.LogFile.Seek(0, 0)
				}
			}
		}()
	}
	cfg.logs = logCfg
	return nil
}
