# server

## 開発に必要なもの

- Go 1.24
- Docker
- [golangci-lint](https://golangci-lint.run/welcome/install/#local-installation) (静的解析)

### VSCode の場合

golangci-lint v2に対応するために、以下の手順が必要です。

- [Go 拡張](https://marketplace.visualstudio.com/items?itemName=golang.Go) の **プレリリースバージョン** をインストールする。
- [{リポジトリルート}/.vscode/settings.template.json](../.vscode/settings.template.json) を `{リポジトリルート}/.vscode/settings.json` としてコピーする。

## ローカルで動かすだけの場合に必要なもの

- Docker

## コマンド

### 開発サーバーの起動

`.env` ファイルを作成して、以下の内容を記述する。

```dotenv
NS_MARIADB_DATABASE=h25s_15
NS_MARIADB_HOSTNAME=db
NS_MARIADB_PASSWORD=password
NS_MARIADB_PORT=3306
NS_MARIADB_USER=root
```

```bash
task up
```

ホットリロードの設定がされているので、コードを変更すると自動でコンテナのビルドが実行されて更新される。
ちょっと待つ必要がある。

### 開発サーバーの停止

```bash
task down
```

### 静的解析

```bash
task lint
```
