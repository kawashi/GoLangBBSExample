# go-bbs-example

![gif](https://i.gyazo.com/02c76a336376f3d5db4c44c92891f595.gif)

GoのAPIサーバを使った掲示板サイトの練習です。  
フロントエンドからGoのAPIを叩き、データの更新と読み取りが出来るようになることを目指しています。

# 構成

- フロントエンド (`/api`) : Vue.js + axios
- バックエンド (`/web`) : Go + Goa + Gorm + MySQL

# 仕様

- ページは1つのみ、投稿画面と投稿一覧が一体になっている
- 画面を開くとAPIの `/user_posts` を `GET` して投稿一覧を取得して表示する
- 投稿ボタンを押すとAPIの `/user_posts` に `POST` されて投稿一覧に追加される
