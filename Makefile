
build_golang:
	@rm -r -f bin
	@mkdir bin
	@mkdir bin/linux
	@mkdir bin/windows
	

	@GOOS=windows go build -ldflags "-s -w" -o win_server.exe .
	@GOOS=linux go build -ldflags "-s -w" -o linux_server.bin .
	@cp win_server.exe ./bin/windows
	@cp linux_server.bin ./bin/linux

	@cd client && GOOS=linux go build -ldflags "-s -w" -o bot.bin .
	@cd client && GOOS=windows go build -ldflags "-H windowsgui -s -w" -o bot.exe .

	@cp client/bot.bin ./bin/linux
	@cp client/bot.exe ./bin/windows

	@rm client/bot.bin
	@rm client/bot.exe

	@rm ./linux_server.bin
	@rm ./win_server.exe

	@echo Success

test:
	python3 builder.py
	./bin/linux/bot.bin

mod: 
	rm go.mod
	rm go.sum
	go mod init homo
	go mod tidy
git: 
	@read -p "Commit: " com; \
	git add .; \
	git add -u; \
    git commit -m "$$com"; \
	git push; 