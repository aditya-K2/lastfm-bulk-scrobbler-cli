package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

var (
	secret   = os.Getenv("LASTFM_SECRET")
	key      = os.Getenv("LASTFM_KEY")
	username = os.Getenv("LASTFM_USERNAME")
	password = os.Getenv("LASTFM_PASSWORD")
)

type scrobble struct {
	ArtistName string    `json:"artistName"`
	AlbumName  string    `json:"albumName"`
	TrackName  string    `json:"trackName"`
	Time       time.Time `json:"time"`
}

type scrobbleList []scrobble

func _scrobble(api *lastfm.Api, file string) error {
	var f []byte
	var err error
	if f, err = os.ReadFile(file); err != nil {
		return err
	}
	sl := &scrobbleList{}
	json.Unmarshal(f, sl)

	fmt.Printf("%d Songs To be scrobbled.\n", len(*sl))
	v := *sl
	batchNo := 0
	totalScrobbles := 0
	totalIgnored := 0
	for _k := 0; _k <= len(v); _k++ {
		acceptedScrobbles := 0
		ignored := 0
		a := lastfm.P{}
		batchNo++
		artists := []string{}
		albums := []string{}
		tracks := []string{}
		timestamps := []string{}
		for _i := 0; _i < 50; _i++ {
			if _k == len(*sl) {
				break
			}
			v := *(sl)
			artists = append(artists, v[_k].ArtistName)
			albums = append(albums, v[_k].AlbumName)
			tracks = append(tracks, v[_k].TrackName)
			timestamps = append(timestamps, time.Now().String())
			_k++
		}
		a["artist"] = artists
		a["album"] = albums
		a["track"] = tracks
		a["timestamp"] = timestamps
		if result, err := api.Track.Scrobble(a); err != nil {
			fmt.Println(result)
			return err
		} else {
			if val, _err := strconv.Atoi(result.Accepted); _err != nil {
				return _err
			} else {
				acceptedScrobbles += val
				totalScrobbles += acceptedScrobbles
			}
			if val, _err := strconv.Atoi(result.Ignored); _err != nil {
				return _err
			} else {
				ignored += val
				totalIgnored += ignored
			}
		}
		fmt.Printf("\r Batch %d (contains %d entries). Out of which %d were accepted and %d were ignored. Total Accepted: %d Total Ignored: %d", batchNo, len(artists), acceptedScrobbles, ignored, totalScrobbles, totalIgnored)
	}
	return nil
}

func main() {
	fmt.Println(key, secret, username, password)
	api := lastfm.New(key, secret)
	api.Login(username, password)

	if len(os.Args) == 1 {
		fmt.Println("No File Provided")
		os.Exit(-1)
	}

	for k := range os.Args {
		if k != 0 {
			fmt.Printf("Scrobbling From File: %s.\n", os.Args[k])
			if err := _scrobble(api, os.Args[k]); err != nil {
				panic(err)
			}
		}
	}
}
