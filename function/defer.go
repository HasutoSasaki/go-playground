package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a

	// exiting: 30
	// second: 20
	// first: 10
}

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error){
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { // defer される関数の定義
		if err === nil {
			err = fx.Commit() // エラーがなければコミット
		}
		if err != nil {
			tx.Rollback() // コミットした結果エラーがあればロールバック
		}
	}() // 無名関数を実行
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// txを使ってさらにデータベースに書き込むコードをここで追加する
	return nil
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func ()  {
		file.Close()
	}, nil
}

func main() {
	if len(os.Args) < 2 { // ファイル名が指定されているか
		log.Fatal("ファイルが指定されていません")
	}
	f, err := os.Open(os.Args[1]) // ファイルをオープン
	if err != nil {
		log.Fatal(err) // オープンに問題あり。エラーを出力して終了
	}
	defer f.Close() // 後始末のコード

	data := make([]byte, 2048) // バイトのスライスを生成
	for {                      // 無限ループ
		count, err := f.Read(data)    // 読み込んだバイト数とエラーを返す
		os.Stdout.Write(data[:count]) // 「標準出力」に出力
		if err != nil {
			if err != io.EOF { // ファイルの終わりでないならば
				log.Fatal(err)
			}
			break // forループを抜ける（ファイルの終わり）
		}
	}

	// 簡易版catの使用例
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
}
