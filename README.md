emotionsns のバックエンドになります

[フロントエンドはこちら](https://github.com/TsunoKento/emotionSNS-frontend)

## 使用技術・ライブラリ

- Go
- Echo
- GORM
- MySQL
- Docker
- AWS

## 導入方法

1. プロジェクトをクローンする

```zsh
git clone https://github.com/TsunoKento/emotionSNS-backend.git
```

2. イメージをビルドする

```zsh
docker-compose build
```

3. コンテナを起動する

```zsh
docker-compose up -d
```

4. サーバーを起動する

```zsh
docker-compose exec go run .
```
