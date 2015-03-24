// Package main provides ...
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PlaylistListResponse struct {
	// Etag: Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// EventId: Serialized EventId of the request which produced this
	// response.
	EventId string `json:"eventId,omitempty"`

	// Items: A list of playlists that match the request criteria.
	Items []*Playlist `json:"items,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "youtube#playlistListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PrevPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PrevPageToken string `json:"prevPageToken,omitempty"`

	TokenPagination *TokenPagination `json:"tokenPagination,omitempty"`

	// VisitorId: The visitorId identifies the visitor.
	VisitorId string `json:"visitorId,omitempty"`
}

type Playlist struct {
	// ContentDetails: The contentDetails object contains information like
	// video count.
	//ContentDetails *PlaylistContentDetails `json:"contentDetails,omitempty"`

	// Etag: Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the playlist.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "youtube#playlist".
	Kind string `json:"kind,omitempty"`

	// Localizations: Localizations for different languages
	Localizations map[string]PlaylistLocalization `json:"localizations,omitempty"`

	// Player: The player object contains information that you would use to
	// play the playlist in an embedded player.
	//Player *PlaylistPlayer `json:"player,omitempty"`

	// Snippet: The snippet object contains basic details about the
	// playlist, such as its title and description.
	Snippet *PlaylistSnippet `json:"snippet,omitempty"`

	// Status: The status object contains status information for the
	// playlist.
	//Status *PlaylistStatus `json:"status,omitempty"`

	// PlaylistItems must be populated with populatePlaylistItems method
	PlaylistItems []*PlaylistItem
}

type PlaylistLocalization struct {
	// Description: The localized strings for playlist's description.
	Description string `json:"description,omitempty"`

	// Title: The localized strings for playlist's title.
	Title string `json:"title,omitempty"`
}

type PlaylistSnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// that published the playlist.
	ChannelId string `json:"channelId,omitempty"`

	// ChannelTitle: The channel title of the channel that the video belongs
	// to.
	ChannelTitle string `json:"channelTitle,omitempty"`

	// DefaultLanguage: The language of the playlist's default title and
	// description.
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// Description: The playlist's description.
	Description string `json:"description,omitempty"`

	// Localized: Localized title and description, read-only.
	Localized *PlaylistLocalization `json:"localized,omitempty"`

	// PublishedAt: The date and time that the playlist was created. The
	// value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Tags: Keyword tags associated with the playlist.
	Tags []string `json:"tags,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the playlist.
	// For each object in the map, the key is the name of the thumbnail
	// image, and the value is an object that contains other information
	// about the thumbnail.
	Thumbnails *ThumbnailDetails `json:"thumbnails,omitempty"`

	// Title: The playlist's title.
	Title string `json:"title,omitempty"`
}

func (pl *Playlist) GetLink() string {
	return "https://www.youtube.com/playlist?list=" + pl.Id
}

// makes youtube api request for items in playlist
// returns PlaylistItemListResponse
func (pl *Playlist) PopulatePlaylistItems() error {

	// craft URL for getting playlist items
	url := baseURL +
		"playlistItems" +
		"?key=" + Config.YT_KEY +
		"&playlistId=" + pl.Id +
		"&part=snippet" +
		"&maxResults=2"

	// make request, get body
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal resp json to plalistitemlistresponse object
	plResp := &PlaylistItemListResponse{}
	err = json.Unmarshal(body, plResp)
	if err != nil {
		panic(err)
	}

	// copy playlistitems onto playlist object
	pl.PlaylistItems = plResp.Items
	return nil
}

func (pl *Playlist) GetItemsDetails() {
	// make contentDetails API call, and save the duration of videos

	// create comma delimited string of video id's
	var videoIds string
	ids := make(map[string]int) // this maps a video id with its index in the playlist
	for idx, item := range pl.PlaylistItems {
		videoIds = videoIds + item.Snippet.ResourceId.VideoId + ","
		ids[item.Snippet.ResourceId.VideoId] = idx
	}
	videoIds = videoIds[:len(videoIds)-1]

	// make call to youtube to get the content details
	// https://www.googleapis.com/youtube/v3/videos?part=contentDetails&id=52Gg9CqhbP8%2CBQ9YtJC-Kd8%2CaiYzrCjS02k&key={YOUR_API_KEY}
	url := baseURL +
		"videos" +
		"?key=" + Config.YT_KEY +
		"&id=" + url.QueryEscape(videoIds) +
		"&part=contentDetails"

	// make request, get body
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal resp json
	vlResp := &VideoListResponse{}
	err = json.Unmarshal(body, vlResp)
	if err != nil {
		panic(err)
	}

	// set details for the playlist item
	for _, video := range vlResp.Items {
		idx := ids[video.Id]
		pl.PlaylistItems[idx].Details = video.ContentDetails
	}
}

func NewPlaylist(id string) *Playlist {
	url := baseURL +
		"playlists" +
		"?key=" + Config.YT_KEY +
		"&id=" + id +
		"&part=snippet"

	// make request, get body
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal resp json to plalistitemlistresponse object
	plResp := &PlaylistListResponse{}
	err = json.Unmarshal(body, plResp)
	if err != nil {
		panic(err)
	}
	return plResp.Items[0]
}
