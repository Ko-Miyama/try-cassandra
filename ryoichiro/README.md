# Spring Boot + Cassandra + React サンプル

## 概要
- `server/` : Spring Boot (Java, Gradle) REST API（Cassandra連携）
- `front/`  : Vite + React + TypeScript フロントエンド
- `docker-compose.yml` : cassandra, server, front の3サービス構成

## 起動方法

1. `docker-compose up --build`
2. ブラウザで `http://localhost:3000` を開く

## 機能
- フロントで「ユーザ名・メールアドレス」入力フォーム
- サーバAPI経由でCassandraに登録

## 備考
- Spring Boot API: `http://localhost:8080/api/users` (POST)
- Cassandra: `localhost:9042` (docker内)
