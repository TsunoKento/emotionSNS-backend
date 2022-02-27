emotionsns のバックエンドになります

[フロントエンドはこちら](https://github.com/TsunoKento/emotionSNS-frontend)

## 使用技術・ライブラリ

- Go
- Echo
- GORM
- MySQL
- Docker
- AWS

## データベース構成図
![er drawio](https://user-images.githubusercontent.com/50817827/155904778-b47b4292-c085-4680-9145-2d33866e5b69.png)

## サーバー構成図
![server drawio](https://user-images.githubusercontent.com/50817827/155905040-33c12eaa-1fbb-4e2e-8ea5-94894fe39e65.png)

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
