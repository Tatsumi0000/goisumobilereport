# sugoimobilereport

- App Store Connect APIからSales Report(tsvファイル)をダウンロードし、ダウンロードしたtsvファイルを簡単にパースするためのGo言語用のライブラリです。

```
go get -u github.com/Tatsumi0000/sugoimobilereport
```

## 使い方

```
touch main.go
```
- 下の例は、新規DLと再DLしたアプリの機種とDL数を表示するサンプルです。

```go
package main

import (
	"fmt"

	"github.com/Tatsumi0000/sugoimobilereport"
)

func main() {
	fmt.Println("Hello, World!")
	issUserID := ""
	keyID := ""
	vendorID := ""
	p8 := ""
	p8Bytes := []byte(p8)

	app, err := sugoimobilereport.NewAppStoreConnectaAPIJwt(issUserID, keyID, p8Bytes)
	if err != nil {
		panic(err)
	}
	// https://help.apple.com/app-store-connect/#/dev8a5831138
	date := "2020-12-31"
	filepath := "./examples.tsv"
	err = app.StoreConnectaAPIRquest("SALES", "SUMMARY", "DAILY", "1_0", vendorID, date, filepath)
	if err != nil {
		panic(err)
	}

	contents, err := sugoimobilereport.ParseTsvFile(filepath)
	if err != nil {
		panic(err)
	}

	for _, content := range contents {
		if content.SKU == "YOUR_SKU_NAME" {
			if content.ProductTypeIdentifier == sugoimobilereport.FreeOrPaidiPhoneAndiPod || content.ProductTypeIdentifier == sugoimobilereport.FreeOrPaidAppUniversal || content.ProductTypeIdentifier == sugoimobilereport.FreeOrPaidAppiPad {
				fmt.Printf("New Install: %v %v times", content.Device, content.Units)
			} else if content.ProductTypeIdentifier == sugoimobilereport.RedownloadOfUniversalApp || content.ProductTypeIdentifier == sugoimobilereport.RedownloadOfiPadOnlyApp || content.ProductTypeIdentifier == sugoimobilereport.RedownloadOfiPhoneOnlyOriOSAndtvOSApp {
				fmt.Printf("Re Install: %v %v times", content.Device, content.Units)
			}
		}
	}
}

```
