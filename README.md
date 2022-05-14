# bmf-tech-client
# Dockerhub
[bmfsan/bmf-tech-client](https://hub.docker.com/r/bmfsan/bmf-tech-client)

# ローカル開発環境
[gobel-api](https://github.com/bmf-san/gobel-api)の起動が必要。

## envファイルの準備
```
cp .env_example .env
```

## hostsの編集
```
127.0.0.1 bmf-tech-client.local
```

## コンテナのビルド
```
make docker-compose-build
```

## コンテナの起動
```
make docker-compose-up
```

or

```
make docker-compose-up-d
```

Then go to `bmf-tech-client.local:81`
