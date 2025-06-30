
dev:
	cd web && yarn run build

dev-watch:
	cd web && yarn run hot-reload

serve:
	docker compose up

ssh:
	docker compose exec air bash