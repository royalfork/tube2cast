// Package main provides ...
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// api url for playlists and playlist items
const baseURL string = "https://www.googleapis.com/youtube/v3/"

type PlaylistItemListResponse struct {
	// Etag: Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// EventId: Serialized EventId of the request which produced this
	// response.
	EventId string `json:"eventId,omitempty"`

	// Items: A list of playlist items that match the request criteria.
	Items []*PlaylistItem `json:"items,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "youtube#playlistItemListResponse".
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

type PlaylistItem struct {
	// Etag: Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the playlist item.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "youtube#playlistItem".
	Kind string `json:"kind,omitempty"`

	// Snippet: The snippet object contains basic details about the playlist
	// item, such as its title and position in the playlist.
	Snippet *PlaylistItemSnippet `json:"snippet,omitempty"`
}

type PlaylistItemSnippet struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the user
	// that added the item to the playlist.
	ChannelId string `json:"channelId,omitempty"`

	// ChannelTitle: Channel title for the channel that the playlist item
	// belongs to.
	ChannelTitle string `json:"channelTitle,omitempty"`

	// Description: The item's description.
	Description string `json:"description,omitempty"`

	// PlaylistId: The ID that YouTube uses to uniquely identify the
	// playlist that the playlist item is in.
	PlaylistId string `json:"playlistId,omitempty"`

	// Position: The order in which the item appears in the playlist. The
	// value uses a zero-based index, so the first item has a position of 0,
	// the second item has a position of 1, and so forth.
	Position int64 `json:"position,omitempty"`

	// PublishedAt: The date and time that the item was added to the
	// playlist. The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ)
	// format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// ResourceId: The id object contains information that can be used to
	// uniquely identify the resource that is included in the playlist as
	// the playlist item.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// Thumbnails: A map of thumbnail images associated with the playlist
	// item. For each object in the map, the key is the name of the
	// thumbnail image, and the value is an object that contains other
	// information about the thumbnail.
	Thumbnails *ThumbnailDetails `json:"thumbnails,omitempty"`

	// Title: The item's title.
	Title string `json:"title,omitempty"`
}

type ResourceId struct {
	// ChannelId: The ID that YouTube uses to uniquely identify the referred
	// resource, if that resource is a channel. This property is only
	// present if the resourceId.kind value is youtube#channel.
	ChannelId string `json:"channelId,omitempty"`

	// Kind: The type of the API resource.
	Kind string `json:"kind,omitempty"`

	// PlaylistId: The ID that YouTube uses to uniquely identify the
	// referred resource, if that resource is a playlist. This property is
	// only present if the resourceId.kind value is youtube#playlist.
	PlaylistId string `json:"playlistId,omitempty"`

	// VideoId: The ID that YouTube uses to uniquely identify the referred
	// resource, if that resource is a video. This property is only present
	// if the resourceId.kind value is youtube#video.
	VideoId string `json:"videoId,omitempty"`
}

type ThumbnailDetails struct {
	// Default: The default image for this resource.
	Default *Thumbnail `json:"default,omitempty"`

	// High: The high quality image for this resource.
	High *Thumbnail `json:"high,omitempty"`

	// Maxres: The maximum resolution quality image for this resource.
	Maxres *Thumbnail `json:"maxres,omitempty"`

	// Medium: The medium quality image for this resource.
	Medium *Thumbnail `json:"medium,omitempty"`

	// Standard: The standard quality image for this resource.
	Standard *Thumbnail `json:"standard,omitempty"`
}

type Thumbnail struct {
	// Height: (Optional) Height of the thumbnail image.
	Height int64 `json:"height,omitempty"`

	// Url: The thumbnail image's URL.
	Url string `json:"url,omitempty"`

	// Width: (Optional) Width of the thumbnail image.
	Width int64 `json:"width,omitempty"`
}

type PageInfo struct {
	// ResultsPerPage: The number of results included in the API response.
	ResultsPerPage int64 `json:"resultsPerPage,omitempty"`

	// TotalResults: The total number of results in the result set.
	TotalResults int64 `json:"totalResults,omitempty"`
}

type TokenPagination struct {
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
		"&maxResults=5"

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
