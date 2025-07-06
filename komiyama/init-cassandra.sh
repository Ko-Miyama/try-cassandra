#!/bin/bash
# Cassandraが起動するまで待機
until cqlsh cassandra -e "describe keyspaces"; do
  echo "Cassandraの起動待ち..."
  sleep 5
done

# キースペースとテーブル作成
cqlsh cassandra -e "CREATE KEYSPACE IF NOT EXISTS testkeyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};"
cqlsh cassandra -e "CREATE TABLE IF NOT EXISTS testkeyspace.items (id UUID PRIMARY KEY, value TEXT);"
