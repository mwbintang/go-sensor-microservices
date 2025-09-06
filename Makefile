DB_USER=root
DB_PASSWORD=Test123!
DB_NAME=sensor_db
DB_HOST=localhost
DB_PORT=3306

MIGRATIONS_DIR=migrations

migrate:
	@echo "🚀 Running migrations..."
	@for file in $(MIGRATIONS_DIR)/*.sql; do \
		echo "📑 Applying $$file..."; \
		mysql -h $(DB_HOST) -P $(DB_PORT) -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < $$file || exit 1; \
	done

reset:
	@echo "⚠️ Dropping and recreating database $(DB_NAME)"
	@dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) --if-exists $(DB_NAME)
	@createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)
	@$(MAKE) migrate
