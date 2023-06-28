# Twitch Chat Bot
This repository contains a Go application that serves as a chat bot for Twitch. It listens to the chat of a specific 
channel and responds to a specific command with a random Chuck Norris fact.

## Installation
Before you start, please ensure you have Go installed on your computer. You can download and install Go from 
[here](https://go.dev/dl/).

### Prerequisites
You should replace the placeholders in the `credentials.json.example` file with your actual Twitch username, OAuth token, and 
channel name. Here's an example of what it might look like:

```json
{
    "username": "yourusername",
    "oauth_token": "youroauthtoken",
    "channel_name": "yourchannelname"
}
```
Store this file as `credentials.json` in that same directory. The bot will use these credentials to connect to the Twitch. 
Remember to never share this file, as it contains your sensitive data. It is ignored by git in the `.gitignore` file.

To build the bot, navigate to root of the project and use the Go build command:

```bash
go build
```
This will create an executable file in the same directory.
## Usage
### Running
To start the bot, simply run the executable that you built in the installation step:

```bash
./dev
```

### How it works
The bot uses the standard Go libraries to connect to the Twitch IRC server over SSL. It then joins the specified 
channel and listens for messages. If a message matches the `!chucknorris` command, the bot fetches a random fact from 
the Chuck Norris API and sends it to the channel.

The bot runs asynchronously, using goroutines to handle incoming messages and API responses. It also 
handles PING/PONG messages from the Twitch server to avoid disconnection. It is designed to comply with the Twitch API 
rate limits.


## License

[MIT](LICENSE)