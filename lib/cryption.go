package kilib


import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "errors"
)


//Adevanced Encryption Standard 
var PwdKey = []byte("HSJ**#CPWNGKXRUO")

//PKCS7 fill mode
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

//Reverse operation of filling, delete filling string
func PKCS7UnPadding(origData []byte) ([]byte, error) {
    length := len(origData)
    if length == 0 {
        return nil, errors.New("Encryption string error!")
    } else {
        unpadding := int(origData[length-1])
        return origData[:(length - unpadding)], nil
    }
}

//Encryption process
func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    origData = PKCS7Padding(origData, blockSize)
    blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    blocMode.CryptBlocks(crypted, origData)
    return crypted, nil
}

//Decryption process
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(cypted))
    blockMode.CryptBlocks(origData, cypted)
    origData, err = PKCS7UnPadding(origData)
    if err != nil {
        return nil, err
    }
    return origData, err
}

//Encryption base64
func EnPwdCode(pwd []byte) (string, error) {
    result, err := AesEcrypt(pwd, PwdKey)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(result), err
}

//Decryption base64
func DePwdCode(pwd string) ([]byte, error) {
    pwdByte, err := base64.StdEncoding.DecodeString(pwd)
    if err != nil {
        return nil, err
    }
    return AesDeCrypt(pwdByte, PwdKey)

}



