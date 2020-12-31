package sugoimobilereport

import (
	"compress/gzip"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AppStoreConnectAPIJwt 認証に必要な構造体
type AppStoreConnectAPIJwt struct {
	IssUserID  string // IssUserID
	KeyID      string // KeyID
	PrivateKey string // PrivateKey
	JwtToken   string // JwtToken
}

func NewAppStoreConnectaAPIJwt(issUserID, keyID string, p8 []byte) (*AppStoreConnectAPIJwt, error) {
	appStoreConnectAPIJwt := AppStoreConnectAPIJwt{}
	appStoreConnectAPIJwt.IssUserID = issUserID
	appStoreConnectAPIJwt.KeyID = keyID
	pemDecode, _ := pem.Decode(p8)
	key, err := x509.ParsePKCS8PrivateKey(pemDecode.Bytes)
	if err != nil {
		return &appStoreConnectAPIJwt, err
	}
	header := map[string]interface{}{
		"alg": "ES256",
		"kid": appStoreConnectAPIJwt.KeyID,
		"typ": "JWT",
	}
	payload := jwt.StandardClaims{
		Issuer:    appStoreConnectAPIJwt.IssUserID,
		ExpiresAt: time.Now().Unix() + 20, // 20 seconds
		Audience:  "appstoreconnect-v1",
	}
	token := jwt.Token{
		Header: header,
		Claims: payload,
		Method: jwt.SigningMethodES256,
	}
	jwtToken, _ := token.SignedString(key)
	appStoreConnectAPIJwt.JwtToken = jwtToken
	return &appStoreConnectAPIJwt, nil
}

func (appStoreConnectAPIJwt *AppStoreConnectAPIJwt) StoreConnectaAPIRquest(reportType, reportSubType, frequency, version, vendorNumber, reportDate, filePath string) error {
	const url = "https://api.appstoreconnect.apple.com/v1/salesReports"
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/a-gzip")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", appStoreConnectAPIJwt.JwtToken))
	params := request.URL.Query()
	params.Add("filter[reportType]", reportType)
	params.Add("filter[reportSubType]", reportSubType)
	params.Add("filter[frequency]", frequency)
	params.Add("filter[version]", version)
	params.Add("filter[vendorNumber]", vendorNumber)
	params.Add("filter[reportDate]", reportDate)
	request.URL.RawQuery = params.Encode()
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		fmt.Println("400 ERROR")
		msg := resp.Body
		return fmt.Errorf("err: %s", msg)
	} else if resp.StatusCode == 403 {
		fmt.Println("403 ERROR")
		msg := resp.Body
		return fmt.Errorf("err: %s", msg)
	}
	convertGzipToTsv(resp.Body, filePath)
	return nil
}

// convertGzipToTsv gzipファイルをtsvファイルに変換
func convertGzipToTsv(gzipFile io.ReadCloser, filePath string) {
	// 出力先のファイルパスを設定
	outPutFile, err := os.Create(filePath)
	// gzipの中身を読み込む
	ungzip, err := gzip.NewReader(gzipFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// tsvファイルを出力
	io.Copy(outPutFile, ungzip)
}
