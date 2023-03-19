build:
	@cd client && go build -ldflags="-s -w" . && GOOS=windows GOOS=windows go build -ldflags "-H windowsgui -s -w" .
	@go build .
	@echo Success

git: 
	@read -p "Commit: " com; \
	git add .; \
	git add -u; \
    git commit -m "$$com"; \
	git push; 