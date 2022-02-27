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

2. env.exampleを.envに変更して、db/dataの`.gitkeep`を削除する
```zsh
cd emotionSNS-backend && mv .env.example .env && rm db/data/.gitkeep
```

3. イメージをビルドする

```zsh
docker-compose build
```

4. コンテナを起動する

```zsh
docker-compose up -d
```

5. サーバーを起動する

```zsh
docker-compose exec go go run .
```
