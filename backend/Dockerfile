FROM golang:1.23-alpine

# ワーキングディレクトリの設定
WORKDIR /app

# go.mod と go.sum をコンテナにコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 全てのソースファイルをコンテナにコピー
COPY . .

# ポート公開
EXPOSE 8080
