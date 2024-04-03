package dbutil

import (
	"log"
	"testing"
)

//CREATE TABLE `user_info` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`name` varchar(255) NOT NULL,
//`pwd` varchar(255) NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=latin1;

func TestSql(t *testing.T) {
	type User struct {
		ID       int    `gorm:"column:id"`
		Name     string `gorm:"column:name"`
		Password string `gorm:"column:pwd"`
	}

	user := User{Name: "dida", Password: "123456"}

	result := GetMysqlClient().Table("user_info").Create(&user) // 通过数据的指针来创建
	log.Println("result:", result)
	// 获取插入后的自增主键
	log.Printf("获取插入后的自增主键:%d", user.ID)
	log.Printf("返回 error:%s", result.Error)
	log.Printf("返回插入记录的条数:%d", result.RowsAffected)
	//log.Printf("", result.Statement)

	users := []User{
		{Name: "dida", Password: "789"},
		{Name: "dida", Password: "777"},
	}
	result = GetMysqlClient().Table("user_info").Create(&users)

	log.Printf("批量插入获取插入后的自增主键:%d", user.ID)
	log.Printf("批量插入返回 error:%s", result.Error)
	log.Printf("批量插入返回插入记录的条数:%d", result.RowsAffected)

	user = User{}
	GetMysqlClient().Table("user_info").Where("id = ?", 1).Limit(1).First(&user)
	log.Println("查询结果:", user)

	result = GetMysqlClient().Table("user_info").Where("id = ?", 1).Updates(User{Name: "keke-dida", Password: "keke"})
	log.Println("更新：", result)
	log.Printf("更新受影响的行数：%d", result.RowsAffected)
	log.Printf("更新返回 error:%s", result.Error)

	result = GetMysqlClient().Table("user_info").Where("id = ?", 1).Delete(&User{})
	log.Println("删除：", result)
	log.Printf("删除受影响的行数：%d", result.RowsAffected)
	log.Printf("删除返回 error:%s", result.Error)

}
