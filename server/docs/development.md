# development

## 使用ライブラリ

- [Echo](https://echo.labstack.com/): Webフレームワーク
- [sqlx](https://github.com/jmoiron/sqlx): SQLラッパー
- [melody](https://github.com/olahol/melody): WebSocketフレームワーク
- [math/big](https://pkg.go.dev/math/big): 式を誤差なく計算するためのライブラリ

## ディレクトリ構成

package by feature の考え方を採用している。
機能ごとにディレクトリを分け、関連するコードをまとめて同じディレクトリに格納する。
機能のディレクトリの下は、技術的関心に応じて分割している。

```txt
.
├── Dockerfile
├── README.md
├── Taskfile.yml
├── compose.yaml
├── docs
├── go.mod
├── internal
│   ├── cards カードに関する処理
│   ├── core
│   │   └── coredb DBに関する共通処理
│   ├── expressions 式に関する処理
│   ├── games ゲームに関する処理
│   │   ├── handler.go ゲームに関するHTTPハンドラ (echoを使用、ファイルは適宜分ける)
│   │   └── internal
│   │       ├── events イベント通知に関する interface 定義
│   │       │    └── ws WebSocket によるイベント処理の実装 (melodyを使用)
│   │       ├── domain ドメインモデル (関連するデータを表す構造体を定義)
│   │       └── repository データの永続化に関する interface 定義
│   │           └── db DB による永続化の実装 (sqlx を使用)
│   ├── items アイテムに関する処理
│   ├── sql スキーマのSQL定義
│   └── users ユーザーに関する処理
└── main.go
```

games ディレクトリを例示したが、他の機能に関するディレクトリも同様の構成になっている。

## WebSocketの動作確認方法

### Postman を使う

Postmanを開いて、左上のハンバーガーバーから「新規」を選択すると、いろいろなのを選べるので、そこからWebSocketを選択する。

### CLI を使う

[websocat](https://github.com/vi/websocat) などを使うことで、ターミナルから動作確認ができる。

```bash
websocat ws://localhost:8080/games/ws
```
