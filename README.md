# Gator
Gator is a command-line interface (CLI) application built in Go as a learning exercise. You can use this application to manage RSS feeds. it allows you to register and log in users, add and follow RSS feeds, aggregate and browse posts, and manage your subscriptions using a PostgreSQL database.

## Features
- Register and log in users
- Add new RSS feeds and follow them
- List all available feeds and your followed feeds
- Unfollow feeds
- Aggregate (fetch and store) posts from feeds
- Browse recent posts from followed feeds

## Installation
Install Gator using Go:
```sh
go install github.com/bbrown4/gator
```
Make sure your Go bin directory is in your `PATH`

### Requirements
- Go 1.23+
- PostgreSQL 15+

### Setup
1. Configure the database:
    - Create a PostgreSQL database
    - Run the migrations in sql/schema/
2. Create a config file:
    - Place a `.gatorconfig.json` file in your home directory with your database URL:
```json
{
  "db_url": "postgres://user:password@localhost:5432/db_name?sslmode=disabled",
  "current_user_name": ""
}
```

## Usage
```sh
gator <command> [args...]
```

### Available Commands
- `register <name>`: Register a new user and log in to that user
- `login <name>`: Log in/switch to an existing user
- `users`: List all users
- `addfeed <name> <url>`: Add a new RSS feed and follow it (must be logged in)
- `feeds`: List all feeds
- `follow <feed_url>`: Follow an existing feed (must be logged in)
- `following:` List feeds you are following (must be logged in)
- `unfollow <feed_url>`: Unfollow a feed (must be logged in)
- `browse [limit]`: Browse recent posts from followed feeds (must be logged in)
- `agg <duration>`: Aggregate (fetch) feeds every duration (e.g. `10s`, `1m`, `1h`)

### Example
```sh
gator register johndoe
gator addfeed "HackerNews" "https://hnrss.org/newest"
gator register janedoe
gator feeds
gator follow https://hnrss.org/newest
gator browse 10
```

### License
MIT