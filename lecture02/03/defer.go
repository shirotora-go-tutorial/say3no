package main

import (
	"os"
)

// defer。これはjsのdeferでいいのかな？
// 関数の中でdefer文を複数追加することができます。関数が最期まで実行された時、このdefer文が逆順に実行されます。

// TODO:特に、リソースをオープンする操作を行なっているようなとき、エラーの発生に対してロールバックし、必要なリソースをクローズする必要があるかと思います。さもなければとても簡単にリソースのリークといった問題を引き起こすことになります。我々はリソースを開く際は一般的に以下のようにします



/* 重複が見られるコード。本来os.Closeは一度でいいはず。
func ReadWrite() bool {
	os.Open("hoge.txt")

	// なにかを行う

	if failuerX {
		os.Close()
		return false
	}

	if failuerY {
		os.Close()
		return false
	}

	os.Close()
	return true

}

*/





