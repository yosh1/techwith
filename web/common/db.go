
/*
RDS_HOSTNAME – DB インスタンスのホスト名。
RDS_PORT – DB インスタンスが接続を許可するポート。デフォルト値は DB エンジンによって異なります。
RDS_DB_NAME – データベース名、ebdb。
RDS_USERNAME – お客様のデータベース用に設定したユーザー名。
RDS_PASSWORD – お客様のデータベース用に設定したパスワード。
*/

package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	conn, err := gorm.Open("mysql", "user:user@tcp(db:3306)/kirakira?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
		panic(err)
	}
	DB = conn
  return DB
}

func GetDB() *gorm.DB {
	return DB
}
