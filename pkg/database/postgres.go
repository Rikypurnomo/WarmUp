package database

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Rikypurnomo/warmup/config"
	log "github.com/Rikypurnomo/warmup/pkg/logger"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// Dbcon ..
	Dbcon *gorm.DB

	// Errdb ..
	Errdb error
)

func Connect() {
	log.Debug("Connetic to postgresql")
	if err := DbOpen(); err != nil {
		log.Error("Failed to connect  db", err)
	}

	Dbcon = GetDbCon()
}

func CloseConnect() {
	log.Debug("Close db connection")
	sqlDB, err := Dbcon.DB()
	if err != nil {
		log.Error("Failed to close db", err)
	}
	sqlDB.Close()
}

func DbOpen() error {
	Dbcon, Errdb = gorm.Open(postgres.Open(config.PGToAddr()), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevePG()),
	})

	if Errdb != nil {
		log.Error("open db Err ", Errdb.Error())
		return Errdb
	}

	Errdb = Dbcon.Use(otelgorm.NewPlugin(otelgorm.WithDBName(config.GetDbName())))
	if Errdb != nil {
		log.Error("otelgorm err ", Errdb.Error())
	}

	db, err := Dbcon.DB()
	if err != nil {
		log.Error("Db Not Connect test Ping :", err.Error())
	}

	if errping := db.Ping(); errping != nil {
		log.Error("Db Not Connect test Ping :", errping.Error())
		return errping
	}
	return nil
}

// GetDbCon ..
func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	if Dbcon == nil {
		if errping := DbOpen(); errping != nil {
			log.Error("try to connect again but error:", errping.Error())
		}
	}
	db, err := Dbcon.DB()
	if err != nil {
		log.Errorf(fmt.Sprintf("Db Not Connect test Ping : %s", err.Error()))
	}
	if errping := db.Ping(); errping != nil {
		log.Errorf(fmt.Sprintf("Db Not Connect test Ping : %s", errping.Error()))
		//errping = nil
		if errping = DbOpen(); errping != nil {
			log.Errorf(fmt.Sprintf("try to connect again but error : %s", errping.Error()))
		}
	}

	return Dbcon
}

// Migrate ..
func Migrate(data ...interface{}) {
	if config.IsEnabledPG() && config.IsDev() {
		fmt.Println("Migrate DB")
		GetDbCon().AutoMigrate(data...)
		for _, v := range data {
			fmt.Printf("Migrate Table %s at %s\n", getType(v), time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
