package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
)




/*
 *密码加密
 */
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

/*
 *密码校验
 */
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}



/**
 * 获取远程的Json并校验
 * param: string url
 * return: string
 * return: error
 */
func GetJson(url string) (jsonString string, err error) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	jsonString = string(data)

	if !gjson.Valid(jsonString) {
		err = errors.New(url + "\n" + "Json Invalid ！！！")
	}

	return jsonString, err
}

func CreateRandomNumber() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
