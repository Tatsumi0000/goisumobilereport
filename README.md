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

### function
- NumberOfNewDownloadsByCountry
  - 新規DLした国コードと機種、数のmapポインタと、合計新規DL数を返す
- NumberOfReDownloadsByCountry
  - 再DLした国コードと機種、数のmapポインタと、合計再DL数を返す
- NumberOfNewDownloads
  - 新規DLした機種と回数のmapポインタと、合計新規DL数を返す
- NumberOfReDownloads
  - 新規DLした機種と回数のmapポインタと、合計新規DL数を返す

### Blog
- [App Store Connect APIから新規DL数を取得](https://tech.pepabo.com/2021/03/29/app-store-connect/)

### refs
- [iTunes Connectの売上とトレンドを自動取得する](https://qiita.com/yosan/items/b820b7b59d33259a7e90#%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E8%A6%8B%E6%96%B9)
- [New Product Type Identifier - 3F? in App Sales Report](https://developer.apple.com/forums/thread/24203)
- [vickxxx/appstore](https://github.com/vickxxx/appstore)
