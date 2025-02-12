.PHONY: up
up: #コンテナ生成・実行
	docker compose up -d --build

.PHONY: down
down: #コンテナ生成・実行
	docker compose down