# EdgeDB Golang Docker Sample

『[Go + Docker Composeを使ってEdgeDBを動かしてみた](https://zenn.dev/tatsuya8429/articles/436e99f2e2bcab)』のサンプルコードです。

## 使い方

1. credentialsフォルダに`local_dev.json` を作成する。
2. `docker-compose up -d --build` を実行
3. `docker-compose exec db edgedb -I local_dev migrate`を実行

## ディレクトリ構成

```sh
.
├── Dockerfile
├── credentials // EdgeDBの認証関連のファイルを入れる
│   └── local_dev.json
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── schema　// EdgeDBのスキーマのファイルを入れる
│   ├── default.esdl
│   └── migrations  // マイグレーションファイルを入れる
│       └── 00001.edgeql
└── src
    ├── handlers  // Controller層
    │   ├── todo_handler.go
    │   └── user_handler.go
    ├── infrastructure // Infrastructure層
    │   ├── config.go
    │   └── dbclient.go
    ├── models // Domain層
    │   ├── todo.go
    │   └── user.go
    └── repositories // Repository層
        ├── todo_repository.go
        └── user_repositories.go
```
