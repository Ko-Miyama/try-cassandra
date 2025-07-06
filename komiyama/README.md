# cassandraにデータを入れる簡易アプリの仕様

1. ```docker compose up -d```を行い、go-appコンテナとcassandraコンテナ起動
2. localhost:8080にブラウザでアクセスし、データベースの中身を見つつデータ操作できる
3. ```docker compose down```を行い、go-appコンテナとcassandraコンテナを修了(&削除)
