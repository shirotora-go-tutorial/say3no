# Go Web プログラミング チュートリアル

[Go Web プログラミング チュートリアル 
](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/index.html)

## 大事な言語特徴

 - GoにはJavaのような例外処理はありません。例外を投げないのです。その代わり、panicとrecoverを使用します。ぜひ覚えておいてください、これは最後の手段として使うことを。つまり、あなたのコードにあってはなりません。もしくはpanicを極力減らしてください。これは非常に強力なツールです。賢く使ってください。では、どのように使うのでしょうか？
 - **配列間の代入は値渡し**.配列が関数と引数となった場合、渡されるのは配列のコピーである。
 - **ポインタを使いたい場合はsliceを使え。**
 - 重要なのは、sliceは参照型だ、ということ。**sliceを何度sliceしようが原初の配列のポインタを参照している。**
 - >mapは他の基本型と異なり、thread-safeではありません。複数のgo-routineを扱う際には必ずmutex lockメカニズムを使用する必要があります。

## 言語特徴っぽいのメモ

 - >そのため大きな構造体を渡す際は、ポインタを使うのが賢い選択というものです。
 - スネークケース使え
 - >注意：一般的にpackageの名前とディレクトリ名は一致させるべきです。
 - `$GOPATH`以下で全てのgoファイルを管理する。
  - >GOPATH内のsrcディレクトリはこれから開発するプログラムにとってメインとなるディレクトリです。全てのソースコードはこのディレクトリに置くことになります。
  - 実際にはもう少し柔軟な対応が出来ると思うけど今のところはこれで納得しておく
 - `go install fuga`で`$GOPATH/pkg/fuga/piyo.a`が作成される。
  - `piyo.a`はアプリケーションパッケージ。たぶんアプリの`.a`
  - これでパスが通っている事になる。以降`import`でパッケージをコールできる。
 - mainを保持するプログラムであれば、`go build`でコンパイル。
  - `go install`によって`go build`で生成されるファイルが`$GOPATH/bin/`に設置される
 - そいつは既にパスが通っていて、簡単にCLIツールが出来るワケだ
 - packageが`main`なら実行アプリ、`main`じゃなければアプリパッケージ
 - init,mainは引数も戻り値も設定できない
 - initはpackageに対してひとつまで。またGoが自動でコールする。
 - `main`packageから全ては始まる。mainが他のパッケージをimportしていたら、コンパイル時にその依存パッケージがimportされる。
 
 
 
 
**リモートパッケージの取得について**

`go get github.com/astaxie/beedb`。リモートパッケージを取得するツール.
やっているのはこういう処理
```go
git clone github.com/astaxie/beedb $GOPATH/src/github.com/astaxie/beedb
cd $GOPATH/src/github.com/astaxie/beedb
go install
```
要するにcloneして`go install`してpkgに追加してる。これによって参照可能にしている。
コールする時はこんな感じ。

```go
import "github.com/astaxie/beedb"
```


### セミコロンはいつつける?

ステートメントの終了のpost_fixとしてC言語同様につかわれている。
Cと違うのは、ソース上には現れないこと。コンパイル時になんかやってくれてるんだろう。
明示的にセミコロンをつける場所はあるんだろうか？



## Goのコマンド
 
### `go build`

 - `go build hoge.go`: コンパイル。
 - `package main`であればカレントディレクトリに実行可能ファイルを生成。
 - そうでない場合はなにもしない。
 - ファイルを指定しない場合はカレントディレクトリの全てのgoファイルをコンパイルする。
 - `_`と`.`から始まるファイルは無視される
 
#### よく使いそうなオプション
 
 - `-o`: 生成ファイルの名前を指定できるのはgccと同じ。
 - `-i`: build後に`go install`してくれる．
 - `-v`: verbose.
 - `-n`: コマンド実行時にどんな処理を行うか教えてくれる。（実行はしない）
 
 
### `go clean`
このコマンドは現在のソースコードパッケージと関連するソースパッケージのなかでコンパイラが生成したファイルを取り除く操作を行います。これらのファイルはすなわち：
 
#### よく使いそうなオプション
 
 - `-i`: go installがインストールするファイル等の、関係するインストールパッケージと実行可能ファイルを取り除きます。
 どゆこと？
 - `-n`: コマンド実行時にどんな処理を行うか教えてくれる。（実行はしない）
 - `-r`: 再帰
 - `-x`: これ-vでよくね？
 
  
### `go fmt`

コードのフォーマットの規約を強いる。劣静的解析ツールみたいなもん。
実際には`gofmt -w -l`がコールされている。
が、実際にはこのコマンドを使うことはそうない。なぜなら*開発ツールには一般的に保存時に自動的に整形を行ってくれるから*。

### `go get`
`go get github.com/hogehoge`みたいなのを。
`git clone github.com/hogehoge $GOPATH/src/.....`
`go install (さっきcloneしたやつ)`

をしてくれる。


### `go install`
`go build`してから、生成物を`pkg`以下の適切なディレクトリに配置してくれる。



### `go test`
ディレクトリ以下の`*_test.go`がオードロードされ、テスト用の実行ファイルが生成/実行されます。

### `go tool`

(U^ω^) アルトくぅ〜ん

yaccってなんやねん。


### `go generate`
(U^ω^) アルトくぅ〜ん

### `godoc`
```go
go get golang.org/x/tools/cmd/godoc
```
でインストール。
`go doc`じゃないよ。`godoc`ね。

`godoc fmt Printf`でfmtのPrintfの確認。
`godoc -src fmt Printf`でコードの確認。


## importについて
### rename

```go
import(
	. "fmt"
	sukinaNamae "os"
)
```

`.`がつくとコール時にpackageNameを省略できる。
`sukinaNamae` ではpythonの`as`みたいに好きな名前をつけられる

```go
Println("hoge") /// fmt.Println("hoge")
sukinaName.Open("hoge.txt")
```


### initだけしたい
 
>  _操作はこのパッケージをインポートするだけでパッケージの中の関数を直接使うわけではなく、このパッケージの中にあるinit関数をコールします。

```go
import(
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
)
````


### recieverについて

>注意深い読者はこのように思うかもしれません。
>PointItBlackの中でSetColorをコールした時、ひょっとして(&bl[i]).SetColor(BLACK)と書かなければならないんじゃないかと。
>SetColorのreceiverは*Boxであり、Boxではありませんから。
>ええ、その通りなんです。この２つの方法はどちらでもかまいません。Goはreceiverがポインタであることを知っています。
>**こいつは自動的に解釈してくれるのです。**

えーつまりだな。レシーバーが値か参照かというのは、コールする際にはほとんど気にする必用がない。
Goのほうでいいかんjいにしてくれる。