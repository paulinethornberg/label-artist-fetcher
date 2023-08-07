package model

import "strings"

type ChannelID string

// set of channel IDs
var (
	P1ID ChannelID = "132"
	P2ID ChannelID = "163"
	P3ID ChannelID = "164"
)

// set of channels we want to support
var (
	ChannelP1 = "p1"
	ChannelP2 = "p2"
	ChannelP3 = "p3"
)

func IsValidChannel(channel string) bool {
	switch strings.ToLower(channel) {
	case ChannelP1, ChannelP2, ChannelP3:
		return true
	default:
		return false
	}
}

func ConvertFromStringToID(channel string) ChannelID {
	switch strings.ToLower(channel) {
	case ChannelP1:
		return P1ID
	case ChannelP2:
		return P2ID
	case ChannelP3:
		return P3ID
	default:
		return ""
	}
}
