package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/Cheng1622/news_go_server/pkg/clog"
)

// RSA加密
func RSAEncrypt(data, publicBytes []byte) []byte {
	var res []byte
	// 解析公钥
	block, _ := pem.Decode(publicBytes)

	if block == nil {
		clog.Log.DPanicln("无法加密, 公钥可能不正确")
		return res
	}

	// 使用X509将解码之后的数据 解析出来
	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		clog.Log.DPanicln("无法加密, 公钥可能不正确:", err)
		return res
	}
	// 使用公钥加密数据
	pubKey := keyInit.(*rsa.PublicKey)
	res, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		clog.Log.DPanicln("无法加密, 公钥可能不正确:", err)
		return res
	}
	// 将数据加密为base64格式
	return []byte(EncodeStr2Base64(string(res)))
}

// 对数据进行解密操作
func RSADecrypt(base64Data, privateBytes []byte) []byte {
	var res []byte
	// 将base64数据解析
	data := []byte(DecodeStrFromBase64(string(base64Data)))
	// 解析私钥
	block, _ := pem.Decode(privateBytes)
	if block == nil {
		clog.Log.DPanicln("无法解密, 私钥可能不正确,解析私钥失败")
		return res
	}
	// 还原数据
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		clog.Log.DPanicln("无法解密, 私钥可能不正确,解析PKCS失败:", err)
		return res
	}
	// 类型断言
	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		clog.Log.DPanicln("无法解密, 得到意外的密钥类型")
		return res
	}
	res, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if err != nil {
		clog.Log.DPanicln("无法解密, 私钥可能不正确,解密PKCS1v15失败:", err)
		return res
	}
	return res
}

// 加密base64字符串
func EncodeStr2Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解密base64字符串
func DecodeStrFromBase64(str string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(decodeBytes)
}
