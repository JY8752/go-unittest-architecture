# go-unittest-architecture

[単体テストの考え方/使い方](https://www.amazon.co.jp/%E5%8D%98%E4%BD%93%E3%83%86%E3%82%B9%E3%83%88%E3%81%AE%E8%80%83%E3%81%88%E6%96%B9-%E4%BD%BF%E3%81%84%E6%96%B9-Vladimir-Khorikov/dp/4839981728)を読んで、テストの書き方、アーキテクチャについて再考するためのリポジトリ。モック、クリーンアーキテクチャを使わないような構成で簡単なAPIを書いてみる。

## 今回作るAPI

例の如くガチャのシュミレーションAPIを作成する。

- ガチャを抽選する(ランダム性のテスト)
- ガチャにアイテムを追加する(プロセス外依存を含んだテスト)

## stack

- wire
- [sqlboiler](https://github.com/volatiletech/sqlboiler)
- [golang-migtate](https://github.com/golang-migrate/migrate)

### DB task

```
task db:create
task db:connect
task db:stop
```

### golang-migrate

#### install 

```
brew install golang-migrate
```

#### task

```
task migrate:create -- create_gacha_table
task migrate:up
task migrate:down
```

### sqlbiler

#### install

```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
```
