# Go Discord Bot

This project is a Discord bot built with Go, designed to ticket & bug management on Discord. It leverages the DiscordGo library to interact with the Discord API, providing a range of features from message handling to dynamic interaction creation.

## Disclaimer
- This project is currently under development and is not yet in the maintenance phase. 
- This is my first Go project, so it may contain unexpected errors and memory leaks.

## Features

- **Environment Variable Configuration**: Leverages a `.env` file for easy configuration of bot token, server IDs, and role management.

## Getting Started

To get started with this Discord bot, follow the steps below:

1. **Clone the Repository**

    ```sh
    git clone https://github.com/Aik-10/go-discordbot.git
    cd go-discordbot
    ```

2. **Set Up Environment Variables**

    Copy the .env.example file to a new file named .env and fill in the necessary details such as your Discord bot token and server IDs.
    ```sh
    cp .env.example .env
    ```


3. **Build the Bot**

    Use the provided GitHub Actions workflow or run the following command:
    ```sh
    go build -v ./cmd/discordbot/main.go
    ```

4. **Run the Bot**

    After building, you can start the bot by running:
    ```sh
    ./discordbot
    ```

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests to help improve the bot.