# Setup Nativefier App Script

A Go script to automate the creation of a `.desktop` file for applications packaged with Nativefier, making them easily accessible from the GNOME applications menu and optionally adding them to GNOME favorites.

## Requirements

- Go (1.15 or later)
- A Nativefier application in `/home/<your-username>/<app-name-lowercase>-linux-x64`

## Usage

1. Ensure you have Go installed on your system. You can install Go by following the instructions on the [official Go website](https://golang.org/doc/install).

2. Save the Go script as `setupApp.go`.

3. Make the script executable:

    ```bash
    chmod +x setupApp.go
    ```

4. Run the script with the name of your Nativefier application as the argument. Replace `<app-name-lowercase>` with the actual name of your application:

    ```bash
    go run setupApp.go <app-name-lowercase>
    ```

    For example, if your Nativefier application is for Evernote, you would run:

    ```bash
    go run setupApp.go evernote
    ```

5. The script will create a `.desktop` file in `~/.local/share/applications`, making the application accessible from the GNOME applications menu.

## Note

- This script assumes that the executable inside the application directory has the same name as the application itself and that it resides in `/home/<your-username>/<app-name-lowercase>-linux-x64`.
- Make sure to adjust the application path in the script if your Nativefier apps are located in a different directory.

