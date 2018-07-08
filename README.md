# Discat: Simple Discord Bot for your non-programmer friends

Discat is a simple discord bot.
You are able to make a bot with your non-programmer friends.

## Feature

## Set responses from csv.

You can set responses from simple csv.
So, edit your bot with your non-programmer fiends.

## How to Use

1. `go get github.com/aimof/discat`
2. make your csv (from google spreadsheet or something)
3. register discordapp and get your token
4. place config.txt and csv
5. Up docker-compose

### CSV format

```csv
word, resp0, resp1, resp2,...
Hello,Hello\,world!,Hi,...
```

### config.txt format

```config.txt
token
bot_name
bot_nickname
csv_dictionary_path
```

example

```config.txt
Bot xxx
mybot
mybot
bot.csv
```
