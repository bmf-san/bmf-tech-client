# bmf-tech-client
[![GitHub release](https://img.shields.io/github/release/bmf-san/bmf-tech-client.svg)](https://github.com/bmf-san/bmf-tech-client/releases)
[![CircleCI](https://circleci.com/gh/bmf-san/bmf-tech-client/tree/master.svg?style=svg)](https://circleci.com/gh/bmf-san/bmf-tech-client/tree/master)
[![GitHub license](https://img.shields.io/github/license/bmf-san/bmf-tech-client)](https://github.com/bmf-san/bmf-tech-client/blob/master/LICENSE)
[![Sourcegraph](https://sourcegraph.com/github.com/bmf-san/bmf-tech-client/-/badge.svg)](https://sourcegraph.com/github.com/bmf-san/bmf-tech-client?badge)

[bmf-tech.com](https://bmf-tech.com/)で運用されているアプリケーションです。

cf. [bmf-techを支える技術](https://bmf-tech.com/posts/bmf-tech%E3%82%92%E6%94%AF%E3%81%88%E3%82%8B%E6%8A%80%E8%A1%93)

# Dockerhub
アプリケーションはコンテナイメージとしてDockerhubにプッシュしています。
[bmfsan/bmf-tech-client](https://hub.docker.com/r/bmfsan/bmf-tech-client)

# ローカル開発環境
1. gobel-apiのセットアップ
APIサーバーとして[gobel-api](https://github.com/bmf-san/gobel-api)の起動が必要です。

2. envファイルの作成
```
cp .env_example .env
```

3. hostsの編集
```
127.0.0.1 bmf-tech-client.local
```

4. 証明書の作成
```
make create-certs
```

5. コンテナのビルド
```
make docker-compose-build
```

6. コンテナの起動
```
make docker-compose-up
または
make docker-compose-up-d
```

`bmf-tech-client.local`にアクセス。

# コントリビューション
IssueやPull Requestを受け付けています。

- [CODE_OF_CONDUCT](https://github.com/bmf-san/bmf-tech-client/blob/master/.github/CODE_OF_CONDUCT.md)
- [CONTRIBUTING](https://github.com/bmf-san/bmf-tech-client/blob/master/.github/CONTRIBUTING.md)

# ライセンス
MITライセンスに基づいています。

[LICENSE](https://github.com/bmf-san/bmf-tech-client/blob/master/LICENSE)

## 作者
[bmf-san](https://github.com/bmf-san)

- Email
  - bmf.infomation@gmail.com
- Blog
  - [bmf-tech.com](http://bmf-tech.com)
- Twitter
  - [bmf-san](https://twitter.com/bmf-san)
