# Go Web プログラミング チュートリアル

[Go Web プログラミング チュートリアル 
](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/index.html)

## 大事な言語特徴

- **配列間の代入は値渡し**.配列が関数と引数となった場合、渡されるのは配列のコピーである。
- **ポインタを使いたい場合はsliceを使え。**
- 重要なのは、sliceは参照型だ、ということ。**sliceを何度sliceしようが原初の配列のポインタを参照している。**

## わからん

![hoge](https://github.com/shirotora-go-tutorial/say3no/wakaran.png)



## 言語特徴っぽいのメモ

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




