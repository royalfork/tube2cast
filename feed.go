// Package main provides ...
package main

import (
	"encoding/xml"
)

type Feed struct {
	XMLName     xml.Name
	Version     string      `xml:"version,attr"`
	FeedChannel interface{} `xml:"channel"`
}

type FeedChannel struct {
	Title   string      `xml:"title"`
	Link    string      `xml:"link"`
	PubDate string      `xml:"pubDate"` // form Tue, 03 Feb 2015 04:02:11 +0000
	Ttl     int         `xml:"ttl"`
	Image   FeedImage   `xml:"image"`
	Item    []*FeedItem `xml:"item"`
}

type FeedChannelItunes struct {
	FeedChannel
	ItunesImage FeedImageItunes `xml:"itunes:image"`
}

type FeedImageItunes struct {
	Href string `xml:"href,attr"`
}

type FeedImage struct {
	Url   string `xml:"url"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type FeedItem struct {
	Title       string          `xml:"title"`
	PubDate     string          `xml:"pubDate"` // form Tue, 03 Feb 2015 04:02:11 +0000
	Link        string          `xml:"link"`
	Description string          `xml:"description"`
	Duration    string          `xml:"itunes:duration"` // form hh:mm:ss
	Content     FeedContent     `xml:"media:content"`
	Image       FeedImageItunes `xml:"itunes:image,omitempty"`
}

type FeedContent struct {
	Url  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

func NewFeedItem(plItem *PlaylistItem) *FeedItem {
	item := FeedItem{
		Title:       plItem.Snippet.Title,
		PubDate:     plItem.Snippet.PublishedAt,
		Link:        plItem.GetLink(),
		Description: plItem.Snippet.Description,
		Duration:    plItem.Details.Duration,
	}
	content := FeedContent{
		Url:  "http://LOCALHOST",
		Type: "audio/mpeg",
	}
	item.Content = content
	image := FeedImageItunes{
		Href: plItem.Snippet.Thumbnails.Default.Url,
	}
	item.Image = image
	return &item
}

func NewFeed(pl Playlist) *Feed {
	// create channel
	channel := FeedChannel{
		Title: pl.Snippet.Title,
		Image: FeedImage{
			Url:   pl.Snippet.Thumbnails.Default.Url,
			Title: pl.Snippet.Title,
			Link:  pl.GetLink(),
		},
		Link:    pl.GetLink(),
		PubDate: pl.Snippet.PublishedAt,
		Ttl:     60,
	}

	// create items from pl items
	for _, plItem := range pl.PlaylistItems {
		channel.Item = append(channel.Item, NewFeedItem(plItem))
	}

	// add itunes things
	ii := FeedImageItunes{channel.Image.Url}
	ch := FeedChannelItunes{channel, ii}

	// create RSS wrapper for channel
	rss := xml.Name{"", "rss"}
	feed := Feed{rss, "2.0", ch}

	return &feed
}
