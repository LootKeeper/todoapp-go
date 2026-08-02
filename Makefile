
ENV ?= dev

ifeq ($(ENV), prod)
	ENV_FILE := .env.prod
	DOCKER_COMPOSE_FILE := docker-compose.prod.yaml
else 
	ENV_FILE := .env.dev
	DOCKER_COMPOSE_FILE := docker-compose.dev.yaml
endif

include $(ENV_FILE)
export 

export PROJECT_ROOT=$(shell pwd)

env-up:
	@echo "Starting the $(ENV) environment..."
	@echo "Using config $(ENV_FILE)"
	@echo "Using compose $(DOCKER_COMPOSE_FILE)"
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d todoapp-postgres


env-down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down todoapp-postgres

env-cleanup:
	@read -p "Are you want to delete volume? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose -f $(DOCKER_COMPOSE_FILE) down todoapp-postgres && \
		rm -rf $(PROJECT_ROOT)/out/pgdata && \
		echo "Env files deleted"; \
	else \
		echo "Aborted"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Please pass 'seq' value. Example: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose -f $(DOCKER_COMPOSE_FILE) run --rm todoapp-postgres-migrate \
		create -ext sql -dir ./migrations -seq $(seq)
