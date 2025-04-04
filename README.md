# Gator

A CLI (Command line Interface) tool for fetching and reading content of any RSS feed.

## Features

- Users can add RSS feeds by URL and provide the frequency in which they want to receive new posts
- Multiple individual user profiles are available so each user can have their own feed and update frequency.
- Feeds can be followed and unfollowed again if no longer desired
- Feeds and posts are stored in a local database so they can always be accessed later on.

### Requirements

- Go lang 
- PostGres database server


### Installation

- Clone the git repo 
- Move the config file `.gatorconfig` to your $HOME directory
- Fill in the file details:

`{"db_url":"postgres://<postgres user>:<postgres password>@localhost:5432/gator?sslmode=disable" }`


### Commands

| Name   | Usage   | Description   |
|---|---|---|
| Login   | login <username>   | logs in sets given user to current user   |
| Register   | register <username>   | registers given user to database and sets as current user  |
| Aggregate  | agg <time duration>  | aggregates reeds with a given frequency (time duration in format e.g. 1m , 300s, 2h, ...  |
| Add feed   | addfeed <name> <url>  | registers a feed with given name and url  |
| Follow  | follow <feed url> | follows a feed that has already been registered  |
| Following  | following  | shows all feeds that the current user is following  |
| Unfollow feed  | unfollow <feed url>  | stops following given feed  |
| Browse  | browse <amount of posts>  | shows all fetched posts of followed feeds, limited to given amount, latest one first |


