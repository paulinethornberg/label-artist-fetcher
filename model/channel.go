package model

import "strings"

type ChannelID string

var (
	ChannelIDP1 ChannelID = "132"
	ChannelIDP2 ChannelID = "163"
	ChannelIDP3 ChannelID = "164"
)

type Channel string

var (
	ChannelP1 Channel = "p1"
	ChannelP2 Channel = "p2"
	ChannelP3 Channel = "p3"
)

// TODO LOOK OVER THIS STRUCTURE AND DO IT NICER :)
func IsValidChannel(channel string) bool {
	switch strings.ToLower(channel) {
	case string(ChannelP1), string(ChannelIDP2), string(ChannelP3):
		return true
	default:
		return false
	}
}

func ConvertFromStringToID(channel string) ChannelID {
	switch strings.ToLower(channel) {
	case "p1":
		return ChannelIDP1
	case "p2":
		return ChannelIDP2
	case "p3":
		return ChannelIDP3
	default:
		return ""
	}
}
