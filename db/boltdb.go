package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var dbfd *bolt.DB //boltdb实例

const (
	bucketname = "base" //桶
)

//自动调用init初始化
func Dbinit() {
	dbfd, err := bolt.Open("my.db", 0600, nil) //打开数据库文件 初始化
	if err != nil {
		fmt.Println("open boltdb error")
		return
	}
	err = dbfd.Update(func(tx *bolt.Tx) error {

		//判断要创建的表是否存在
		b := tx.Bucket([]byte(bucketname))
		if b == nil {
			//创建叫"MyBucket"的表
			_, err := tx.CreateBucket([]byte(bucketname))
			if err != nil {
				//也可以在这里对表做插入操作
				log.Fatal(err)
			}
		}
		//一定要返回nil
		return nil
	})
	//更新数据库失败
	if err != nil {
		log.Fatal(err)
	}

}
func DbClose() {
	dbfd.Close()
}

//func (b *Bucket) Get(key []byte) []byte
func GetDb(key string) (string, error) {
	if key == "" {
		return "", errors.New("key if empty")
	}
	var res string
	err := dbfd.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname)) //打开视图
		v := b.Get([]byte(key))            // Get查询 key为answer的value
		if v == nil {
			return errors.New("key exsit")
		}
		res = string(v) //转换为string
		return nil
	})
	if err != nil {
		return "", err
	}
	return res, nil
}

func PutDb(key, value string) error {
	if key == "" || value == "" {
		return errors.New("parm is empty")
	}
	err := dbfd.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketname)) //打开视图
		// put 数据库插入
		//func (b *Bucket) Put(key []byte, value []byte) error
		if err := b.Put([]byte(key), []byte(value)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("put key value error")
	}
	return nil
}
