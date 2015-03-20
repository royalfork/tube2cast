Tube2Cast is a utility which converts youtube playlists to podcasts.

It:
  1. Exposes an endpoint: http://localhost/pl/{youtube playlistId}.rss
      - queries youtube API for playlist information
      - formats the youtube playlist as an actual podcast RSS feed
  2. Exposes an endpoint: http://localhost/asset/{youtube id}.mp3
      - downloads video, converts to mp3, and serves it

This lets anyone make ad-hoc audio podcasts from a playlist of youtube videos, that can be streamed, or saved.
