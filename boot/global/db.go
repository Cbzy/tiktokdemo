package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DBEngine *gorm.DB
	DB       *gorm.DB
)

func GormMysql() *gorm.DB {

	m := DyDatabase
	if m.DBName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN: m.Dsn(),
	}
	//fmt.Println(mysqlConfig)
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		log.Println("数据库链接错误", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(int(m.MaxIdleConns))
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetMaxOpenConns(int(m.MaxOpenConns))
		return db
	}
}

func (m *Database) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.DBName + "?" + "&parseTime=True&loc=Local"
}
