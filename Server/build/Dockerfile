# ベースイメージ (開発用)
FROM golang:1.19.2 AS development

ENV TZ=Asia/Tokyo

RUN go install github.com/go-delve/delve/cmd/dlv@v1.21.2 && \ 
		go install github.com/cosmtrek/air@v1.41.0

WORKDIR /app
# Goモジュールとソースコードをコピー
COPY go.mod go.sum ./
RUN go mod download

# ベースイメージ (本番用)
FROM alpine:3.16 AS production

# 必要なランタイム環境をインストール
RUN apk add --no-cache \
	bash \
	libc6-compat

WORKDIR /app

# 開発用ステージからビルド成果物をコピー
COPY --from=development /app/main .

# ポートを公開
EXPOSE 8080

# 本番用エントリポイント
CMD ["./main"]