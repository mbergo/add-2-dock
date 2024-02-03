package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/user"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: setupApp <app-name-lowercase>")
        os.Exit(1)
    }

    appName := strings.ToLower(os.Args[1])
    usr, _ := user.Current()
    homeDir := usr.HomeDir
    appPath := fmt.Sprintf("%s/%s-linux-x64", homeDir, appName)
    displayName := strings.Title(appName)

    localAppsDir := fmt.Sprintf("%s/.local/share/applications", homeDir)
    desktopFilePath := fmt.Sprintf("%s/%s.desktop", localAppsDir, displayName)

    // Check if application directory exists
    if _, err := os.Stat(appPath); os.IsNotExist(err) {
        fmt.Printf("The application directory %s does not exist. Please check the path and try again.\n", appPath)
        os.Exit(1)
    }

    // Create applications directory if it does not exist
    os.MkdirAll(localAppsDir, os.ModePerm)

    // Create .desktop file
    fileContent := fmt.Sprintf(`[Desktop Entry]
Type=Application
Name=%s
Comment=%s Desktop Application
Exec=%s/%s %%U
Icon=%s/resources/app/icon.png
Terminal=false
Categories=Office;`, displayName, displayName, appPath, appName, appPath)

    file, err := os.Create(desktopFilePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    writer.WriteString(fileContent)
    writer.Flush()

    // Updating desktop database
    exec.Command("update-desktop-database", localAppsDir).Run()

    fmt.Printf("%s has been successfully added to your applications.\n", displayName)
}

