# goisumobilereport

- App Store Connect APIからSales Report(tsvファイル)をダウンロードし、ダウンロードしたtsvファイルを簡単にパースするためのGo言語用のライブラリです。

```
go get github.com/Tatsumi0000/goisumobilereport
```

```go
package main

import (
	"fmt"
	"io/ioutil"

	goisu "github.com/Tatsumi0000/goisumobilereport/goisumobilereport"
)

func main() {
	const (
		// Issuer ID
		issUserID = "ISSUER_ID"
		// 生成したキーID
		keyID = "KEY_ID"
		// App Store Connectで生成したファイル
		p8Filepath = "./XXXXXX.p8"
		// App Store Connect APIで取得したtsvファイルを保存するパス
		filepath = "./salesReport.tsv"
	)
	// p8の中身を読み込む
	p8, _ := ioutil.ReadFile(p8Filepath)
	// JWTの準備
	app, _ := goisu.NewAppStoreConnectAPIJwt(issUserID, keyID, p8)

	// VENDOR_IDにチームのベンダーIDを入れて下さい。
	// 2021年03月23日のレポートを取得
	app.StoreConnectAPIRequest("SALES", "SUMMARY", "DAILY", "1_0", "VENDOR_ID", "2021-03-23", filepath)

	// tsvファイルをパースしてSlicesのポインタを返す
	contents, _ := goisu.ParseTsvFile(filepath)

	// SKUの箇所に取得したいアプリのSKUを入れて下さい。
	// 新規DLした機種と回数のmapポインタと、合計新規DL数を返す
	newInstallCounts, newInstallSumCounts := goisu.NumberOfNewDownloads(contents, "SKU")

	for key, value := range *newInstallCounts {
		fmt.Printf("%v: %v回\n", key, value)
	}
	fmt.Printf("合計新規インストール数: %v回\n", newInstallSumCounts)
}
```

### Blog
- [App Store Connect APIから新規DL数を取得](https://tech.pepabo.com/2021/03/29/app-store-connect/)
