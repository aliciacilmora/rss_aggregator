# RSS Aggregator with Web Crawler
This project is an RSS aggregator with a web crawler written in Go. It collects RSS feeds from various sources and aggregates them into a single feed. The web crawler extends the functionality by discovering new RSS feeds from websites.

## Features
- Fetch RSS feeds from predefined sources.
- Discover new RSS feeds through web crawling.
- Aggregate and consolidate RSS feeds into a single feed.
- Customizable and extendable.

## Installation
``
    git clone https://github.com/aliciacilmora/rss_aggregator.git
``
## Install dependencies
``
    go mod tidy
``

## Configuration
Create .env file in the main directory:-
``` .env
PORT=8080
DB_URL=postgres://[username]:[password]@localhost:5432/[database_name]
```

## Run the application
``
    go build && ./rss_aggreagator
``
