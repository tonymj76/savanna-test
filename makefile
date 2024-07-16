run:
	docker-compose up --build -d
log:
	docker logs -f savanna_api
gen:
	@go generate ./ent

update: ## run update after you run create_migration
	@docker-compose down
	@go mod tidy -e -v
	@go generate ./ent
	$(MAKE) run

create_migration: ## Generates a sql migration file based on the defined Ent models locally. Usage is `make create_migration name=Foobar`
	@go run -mod=mod entgo.io/ent/cmd/ent new $(name)
	$(MAKE) update
check:


.PHONY:run down gen update create_migration
