package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"time"
)

func main() {
	fmt.Println("クリップボード監視を開始します。Ctrl+Cで終了します。")

	var lastClipboardContent string

	for {
		// クリップボードの内容を取得
		content, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("クリップボードの読み取りに失敗しました:", err)
			time.Sleep(1 * time.Second) // エラー時に少し待つ
			continue
		}

		// 前回と内容が異なれば表示
		if content != lastClipboardContent {
			lastClipboardContent = content
			fmt.Println("新しいクリップボード内容:", content)
			fmt.Println("文字数:", len(content))
			fmt.Println()
		}

		// 短い間隔でチェック
		time.Sleep(500 * time.Millisecond)
	}
}
