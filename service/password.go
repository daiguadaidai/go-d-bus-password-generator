package service

import (
	"github.com/daiguadaidai/go-d-bus/common"
	"github.com/ngaut/log"
	"fmt"
)


/* 加密
Params:
    _value: 需要加密的串
 */
func Encrypt(_value string) {
	encodePassowrd, err := common.Encrypt(_value)
	if err != nil {
		log.Errorf("加密失败. 需要加密的字符串: %v. %v", _value, err)
	}

	fmt.Println("加密前:", _value)
	fmt.Println("加密后:", encodePassowrd)
}

/* 解密
Params:
    _value: 已经加密的串
 */
func Decrypt(_value string) {
	password, err := common.Decrypt(_value)
	if err != nil {
		log.Errorf("解密失败. 需要解密的字符串: %v. %v", _value, err)
	}

	fmt.Println("解密前:", _value)
	fmt.Println("解密后:", password)
}
