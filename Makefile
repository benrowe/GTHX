
dev:
	cd web && yarn run build

serve:
	docker compose up -d

ssh:
	docker compose exec air bash