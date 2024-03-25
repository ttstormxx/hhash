package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func calculateMD5String(input string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateMD5File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件: %s", err)
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("无法读取文件: %s", err)
	}

	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateSHA1String(input string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateSHA1File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件: %s", err)
	}
	defer file.Close()

	hasher := sha1.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("无法读取文件: %s", err)
	}

	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateSHA256String(input string) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateSHA256File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件: %s", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("无法读取文件: %s", err)
	}

	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateSHA512String(input string) (string, error) {
	hasher := sha512.New()
	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func calculateSHA512File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件: %s", err)
	}
	defer file.Close()

	hasher := sha512.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("无法读取文件: %s", err)
	}

	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func callHashAlgorithemString(algorithem string, input string) (string, error) {
	switch algorithem {
	case "md5":
		return calculateMD5String(input)
	case "sha1":
		return calculateSHA1String(input)
	case "sha256":
		return calculateSHA256String(input)
	case "sha512":
		return calculateSHA512String(input)
	default:
		return calculateMD5String(input)
	}
}

func callHashAlgorithemFile(algorithem string, filepath string) (string, error) {
	switch algorithem {
	case "md5":
		return calculateMD5File(filepath)
	case "sha1":
		return calculateSHA1File(filepath)
	case "sha256":
		return calculateSHA256File(filepath)
	case "sha512":
		return calculateSHA512File(filepath)
	default:
		return calculateMD5File(filepath)
	}
}

func main() {
	var filePath string     //目标文件路径
	var inputString string  //目标字符串
	var hashStream []string //哈希顺序
	var hashData string     //哈希计算结果
	helpInfo := "Usage:\nhhash string|file md5|sha1|sha256|sha512... [file]\nhhash 12345 md5\nhhash 12345 md5 sha1\nhhash hash.zip sha256 file\nhhash hash.zip sha256 md5 file\n"
	if len(os.Args) == 1 {
		fmt.Print(helpInfo)
		os.Exit(1)
	}
	args := os.Args[1:]
	if len(args) < 2 {
		// 最短形式 hhash 123456 md5
		// 2个参数
		fmt.Print(helpInfo)
		os.Exit(1)
	}
	hashStream = args[1:]
	if len(hashStream) == 1 && hashStream[0] == "file" {
		fmt.Print("请输入合法的哈希名称: md5|sha1|sha256|sha512\n")
		os.Exit(1)
	}
	for _, v := range hashStream {
		if v != "md5" && v != "sha1" && v != "sha256" && v != "sha512" && v != "file" {
			fmt.Print("请输入合法的哈希名称: md5|sha1|sha256|sha512\n")
			os.Exit(1)
		}
	}
	if strings.ToLower(hashStream[len(hashStream)-1]) == "file" {
		// 计算文件哈希
		filePath = args[0]

		hashStream = hashStream[:len(hashStream)-1] //去除file参数
		firstAlgorithemFile := hashStream[0]
		hashData, err := callHashAlgorithemFile(firstAlgorithemFile, filePath)
		if err != nil {
			fmt.Printf("文件不存在: %s\n", filePath)
			os.Exit(1)
		}
		if len(hashStream) > 1 {

			for _, v := range hashStream[1:] {
				hashData, _ = callHashAlgorithemString(v, hashData)
			}
		}
		fmt.Printf("【文件】 %s 的 %s 哈希值为: %s\n", filePath, hashStream, hashData)

	} else {
		// 计算字符串哈希
		inputString = args[0]
		hashData = inputString
		for _, algorithem := range hashStream {
			hashData, _ = callHashAlgorithemString(algorithem, hashData)
		}
		fmt.Printf("【字符串】 \"%s\" 的 %s 哈希值为: %s\n", inputString, hashStream, hashData)
	}

}
