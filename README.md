# Go Web プログラミング チュートリアル

[Go Web プログラミング チュートリアル 
](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/index.html)

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

続きは*プログラムの全体構成*から
