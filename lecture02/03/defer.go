package main

import (
	"os"
	"fmt"
)

// defer。これはjsのdeferでいいのかな？
// 関数の中でdefer文を複数追加することができます。関数が最期まで実行された時、このdefer文が逆順に実行されます。

// TODO:特に、リソースをオープンする操作を行なっているようなとき、エラーの発生に対してロールバックし、必要なリソースをクローズする必要があるかと思います。さもなければとても簡単にリソースのリークといった問題を引き起こすことになります。我々はリソースを開く際は一般的に以下のようにします

func main() {
	ReadWrite()

	// deferが複数ある場合はLIFO
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}

// この場合、deferはreturnの後に実行される。
func ReadWrite() bool {
	var file, err = os.Open("hoge.txt")
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}
