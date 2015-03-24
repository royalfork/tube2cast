// Package main provides ...
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type VideoListResponse struct {
	// Etag: Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// EventId: Serialized EventId of the request which produced this
	// response.
	EventId string `json:"eventId,omitempty"`

	// Items: A list of videos that match the request criteria.
	Items []*Video `json:"items,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "youtube#videoListResponse".
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

type Video struct {
	// AgeGating: Age restriction details related to a video.
	//AgeGating *VideoAgeGating `json:"ageGating,omitempty"`

	// ContentDetails: The contentDetails object contains information about
	// the video content, including the length of the video and its aspect
	// ratio.
	ContentDetails *VideoContentDetails `json:"contentDetails,omitempty"`

	// ConversionPings: The conversionPings object encapsulates information
	// about url pings that need to be respected by the App in different
	// video contexts.
	//ConversionPings *VideoConversionPings `json:"conversionPings,omitempty"`

	// Etag: Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// FileDetails: The fileDetails object encapsulates information about
	// the video file that was uploaded to YouTube, including the file's
	// resolution, duration, audio and video codecs, stream bitrates, and
	// more. This data can only be retrieved by the video owner.
	//FileDetails *VideoFileDetails `json:"fileDetails,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the video.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "youtube#video".
	Kind string `json:"kind,omitempty"`

	// LiveStreamingDetails: The liveStreamingDetails object contains
	// metadata about a live video broadcast. The object will only be
	// present in a video resource if the video is an upcoming, live, or
	// completed live broadcast.
	//LiveStreamingDetails *VideoLiveStreamingDetails `json:"liveStreamingDetails,omitempty"`

	// Localizations: List with all localizations.
	//Localizations map[string]VideoLocalization `json:"localizations,omitempty"`

	// MonetizationDetails: The monetizationDetails object encapsulates
	// information about the monetization status of the video.
	//MonetizationDetails *VideoMonetizationDetails `json:"monetizationDetails,omitempty"`

	// Player: The player object contains information that you would use to
	// play the video in an embedded player.
	//Player *VideoPlayer `json:"player,omitempty"`

	// ProcessingDetails: The processingProgress object encapsulates
	// information about YouTube's progress in processing the uploaded video
	// file. The properties in the object identify the current processing
	// status and an estimate of the time remaining until YouTube finishes
	// processing the video. This part also indicates whether different
	// types of data or content, such as file details or thumbnail images,
	// are available for the video.
	//
	// The processingProgress object is
	// designed to be polled so that the video uploaded can track the
	// progress that YouTube has made in processing the uploaded video file.
	// This data can only be retrieved by the video owner.
	//ProcessingDetails *VideoProcessingDetails `json:"processingDetails,omitempty"`

	// ProjectDetails: The projectDetails object contains information about
	// the project specific video metadata.
	//ProjectDetails *VideoProjectDetails `json:"projectDetails,omitempty"`

	// RecordingDetails: The recordingDetails object encapsulates
	// information about the location, date and address where the video was
	// recorded.
	//RecordingDetails *VideoRecordingDetails `json:"recordingDetails,omitempty"`

	// Snippet: The snippet object contains basic details about the video,
	// such as its title, description, and category.
	//Snippet *VideoSnippet `json:"snippet,omitempty"`

	// Statistics: The statistics object contains statistics about the
	// video.
	//Statistics *VideoStatistics `json:"statistics,omitempty"`

	// Status: The status object contains information about the video's
	// uploading, processing, and privacy statuses.
	//Status *VideoStatus `json:"status,omitempty"`

	// Suggestions: The suggestions object encapsulates suggestions that
	// identify opportunities to improve the video quality or the metadata
	// for the uploaded video. This data can only be retrieved by the video
	// owner.
	//Suggestions *VideoSuggestions `json:"suggestions,omitempty"`

	// TopicDetails: The topicDetails object encapsulates information about
	// Freebase topics associated with the video.
	//TopicDetails *VideoTopicDetails `json:"topicDetails,omitempty"`
}

type VideoContentDetails struct {
	// Caption: The value of captions indicates whether the video has
	// captions or not.
	Caption string `json:"caption,omitempty"`

	// ContentRating: Specifies the ratings that the video received under
	// various rating schemes.
	//ContentRating *ContentRating `json:"contentRating,omitempty"`

	// CountryRestriction: The countryRestriction object contains
	// information about the countries where a video is (or is not)
	// viewable.
	//CountryRestriction *AccessPolicy `json:"countryRestriction,omitempty"`

	// Definition: The value of definition indicates whether the video is
	// available in high definition or only in standard definition.
	Definition string `json:"definition,omitempty"`

	// Dimension: The value of dimension indicates whether the video is
	// available in 3D or in 2D.
	Dimension string `json:"dimension,omitempty"`

	// Duration: The length of the video. The tag value is an ISO 8601
	// duration in the format PT#M#S, in which the letters PT indicate that
	// the value specifies a period of time, and the letters M and S refer
	// to length in minutes and seconds, respectively. The # characters
	// preceding the M and S letters are both integers that specify the
	// number of minutes (or seconds) of the video. For example, a value of
	// PT15M51S indicates that the video is 15 minutes and 51 seconds long.
	Duration string `json:"duration,omitempty"`

	// LicensedContent: The value of is_license_content indicates whether
	// the video is licensed content.
	LicensedContent bool `json:"licensedContent,omitempty"`

	// RegionRestriction: The regionRestriction object contains information
	// about the countries where a video is (or is not) viewable. The object
	// will contain either the contentDetails.regionRestriction.allowed
	// property or the contentDetails.regionRestriction.blocked property.
	//RegionRestriction *VideoContentDetailsRegionRestriction `json:"regionRestriction,omitempty"`
}

func (vcd VideoContentDetails) DurationToItunesFormat() string {
	// convert PT99H59M59S -> 99:59:59
	r, _ := regexp.Compile("^PT(([0-9]{1,2})H)?(([0-9]{1,2})M)?(([0-9]{1,2})S)?")
	matches := r.FindStringSubmatch(vcd.Duration)
	hr, _ := strconv.Atoi(matches[2])
	min, _ := strconv.Atoi(matches[4])
	sec, _ := strconv.Atoi(matches[6])
	return fmt.Sprintf("%02d:%02d:%02d", hr, min, sec)
}
