package service

import (
	"fmt"
	redis "mall-admin-server/common/service"
	"math/rand"
	"strconv"
)

var prefix = "portal:authCode:"

type UmsMemberService struct {
	RedisService redis.RedisService
}

// 生成验证码
func (iService UmsMemberService) GenerateAuthCode(telephone string) error {
	authCode := ""
	for i := 0; i < 6; i++ {
		authCode += strconv.Itoa(rand.Intn(10))
	}
	return iService.RedisService.Set(fmt.Sprintf("%s%s", prefix, telephone), authCode)
}

// 验证码校验
func (iService UmsMemberService) VerifyAuthCode(telephone, authCode string) bool {
	if authCode == "" {
		return false
	}
	realCode := iService.RedisService.Get(fmt.Sprintf("%s%s", prefix, telephone))
	if realCode == "" {
		return false
	}
	if realCode == authCode {
		_ = iService.RedisService.Remove(fmt.Sprintf("%s%s", prefix, telephone))
		return true
	}
	return false
}
