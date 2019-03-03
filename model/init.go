package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Local *gorm.DB
}

var DB *Database

func InitQSC(db *Database) error {
	_, err := GetAssociationByName("求是潮")
	if err != nil {
		// Create QSC association
		association := AssociationModel{
			Name: "求是潮",
			DepartmentList: "人力资源部门&技术研发中心&产品运营部门&设计与视觉中心&推广与策划中心&水朝夕工作室&新闻资讯中心&视频团队(主持人)&视频团队(拍摄与制作)&摄影部",
		}

		if err := association.Create(); err != nil {
			panic(err)
		}
	}
	return nil
}

func (db *Database) Init() {
	DB = &Database{
		Local: GetLocalDB(),
	}
	if err := DB.Local.AutoMigrate(&UserModel{}, &InstanceModel{}, &FormModel{}, &FreshmanModel{}, &IntentModel{}, &InterviewModel{}, &AssociationModel{}, &LogModel{}, &SmsModel{}).Error; err != nil {
		log.Debug(err.Error())
		panic(err)
	}
	InitQSC(db)
}
func (db *Database) Close() {
	DB.Local.Close()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database: %s connection failed.", name)
	}
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}

func InitLocalDB() *gorm.DB {
	res := openDB(viper.GetString("local_db.username"),
		viper.GetString("local_db.password"),
			viper.GetString("local_db.addr"),
				viper.GetString("local_db.name"))
	return res
}
func GetLocalDB() *gorm.DB {
	return InitLocalDB()
}
