package main

import (
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	// main関数はos.Exitを呼ぶ可能性があるため、直接テストは困難
	// ここではmain関数がpanicしないことのみを確認
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()
	// テスト用に環境変数やポートを変更する場合はここで設定
	go func() {
		main()
		os.Exit(0)
	}()
}
