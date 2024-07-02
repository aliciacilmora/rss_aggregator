# RSS Aggregator with Web Crawler
This project is an RSS aggregator with a web crawler written in Go. It collects RSS feeds from various sources and aggregates them into a single feed. The web crawler extends the functionality by discovering new RSS feeds from websites.

## Features
- Fetch RSS feeds from predefined sources.
- Discover new RSS feeds through web crawling.
- Aggregate and consolidate RSS feeds into a single feed.
- Customizable and extendable.

## Installation
```Bash
git clone https://github.com/aliciacilmora/rss_aggregator.git
```
## Install dependencies
```Bash
go mod tidy
```

## Configuration
Create .env file in the main directory:-
``` .env
PORT=8080
DB_URL=postgres://[username]:[password]@localhost:5432/[database_name]
```

## Run the application
```Bash
go build && ./rss_aggreagator
```

## Usage
You can use Thunder Client by downloading the Thunder Client extension on VS code or curl directly from terminal.

It runs of http://localhost:8080/v1/

### API Endpoints
- User Management:
    * GET /v1/users: Get user information.
    * POST /v1/users: Create a new user.
- Error Handling:
    * GET /v1/err: Handle errors.
- Feeds:
    * GET /v1/feeds: Get list of feeds.
    * POST /v1/feeds: Add new feeds from XML or similar.
- Feed Follows:
    * GET /v1/feed_follows: See followed feeds.
    * POST /v1/feed_follows: Follow a new feed.
    * DELETE /v1/feed_follows/{feedFollowID}: Unfollow a feed.
- Posts:
    * GET /v1/posts: See posts from followed feeds.
