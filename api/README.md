# APIサーバ

goa + gorm 構成のAPIサーバです。

# 参考サイト

## Golang

- [A Tour of Go](https://tour.golang.org/welcome/1)
- [他言語プログラマがgolangの基本を押さえる為のまとめ - Qiita](https://qiita.com/gcfuji/items/e2a3d7ce7ab8868e37f7)
- [Golangで環境変数を扱う](https://orebibou.com/2017/04/golang%E3%81%A7%E7%92%B0%E5%A2%83%E5%A4%89%E6%95%B0%E3%82%92%E6%89%B1%E3%81%86/)

## Goa

- [goadesign/goa - GitHub](https://github.com/goadesign/goa)
- [goa をはじめる](https://goa.design/ja/learn/guide/)
- [GolangのgoaでAPIをデザインしよう（基本編）](http://tikasan.hatenablog.com/entry/2017/05/05/212501)

## Gorm

- [【GORM】Go言語でORM触ってみた - Qiita](https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81)
- [Go lang入門 GinとGORMでWebApp - Qiita](https://qiita.com/Anharu/items/ce644c521a4d52fafb7e)
- [GolangのgoaでAPIをデザインしよう（Model編）](http://tikasan.hatenablog.com/entry/2017/05/15/172535)
- [gorma プラグイン](https://goa.design/ja/extend/gorma/)
- [ReSTARTR/goa-gorm-sample - GitHub](https://github.com/ReSTARTR/goa-gorm-sample)
- [クラスター社のGo製WebAPI開発で主に使ってるライブラリについて - Qiita](https://qiita.com/kyokomi/items/dcd8384a0a042d72d22d)

gormaの導入中にエラーが出ましたが、下記の記事で解決しました。

```sh
$ goagen gen -d BBS-Example/api/bbs-example-server/design --pkg-path=github.com/goadesign/gorma

../../../../github.com/goadesign/gorma/relationalmodel.go:9:2: cannot find package "bitbucket.org/pkg/inflect" in any of:
	/usr/local/Cellar/go/1.10.2/libexec/src/bitbucket.org/pkg/inflect (from $GOROOT)
	/Users/allen/.go/src/bitbucket.org/pkg/inflect (from $GOPATH)
```

- [Code generated following the example in the README is invalid - GitHub](https://github.com/goadesign/goa/issues/231)

```sh
$ cd $GOPATH/src/github.com/goadesign/gorma
$ make depend

go: missing Mercurial command. See https://golang.org/s/gogetcmd
```

- [goのフレームワーク ginの"gin-scaffold"を入れようとしたらエラーが出た。](http://tabilike.hatenablog.com/entry/2017/11/01/113023)
