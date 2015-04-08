# Tube2Cast

Tube2Cast is a web service which converts Youtube Playlists into audio-only podcasts.  Simply add "http://{server}/pl/{youtube playlist id}.rss" to your favorite podcast app, and the podcast will be created on the fly.  New videos can be added to the youtube playlist, and additions will be shown in your podcast app.

## Installation

Create secrets.json file in package root.

This file has form:
{
  "YT_KEY": "your youtube api key"
}


go build main.go
./tube2cast

## Usage

Running this app will expose 2 endpoints:

1. http://localhost/pl/{youtube playlistId}.rss
    - queries youtube API for playlist information
    - formats the youtube playlist as an actual podcast RSS feed
2. http://localhost/asset/{youtube id}.mp3
    - downloads video, converts to mp3, and serves it
