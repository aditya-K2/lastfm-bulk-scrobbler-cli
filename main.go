package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/shkh/lastfm-go/lastfm"
)

var (
	secret    = os.Getenv("LASTFM_SECRET")
	key       = os.Getenv("LASTFM_KEY")
	username  = os.Getenv("LASTFM_USERNAME")
	password  = os.Getenv("LASTFM_PASSWORD")
	schema    = "normal"
	threshold = 60 * 4 * 1000
)

type scrobble struct {
	ArtistName   string `normal:"artistName"   spotify:"master_metadata_album_artist_name"`
	AlbumName    string `normal:"albumName"    spotify:"master_metadata_album_album_name"`
	TrackName    string `normal:"trackName"    spotify:"master_metadata_track_name"`
	Milliseconds int    `normal:"milliseconds" spotify:"ms_played"`
}

type scrobbleList []scrobble

func lessThan(l, r int) bool {
	return l < r
}

func _scrobble(api *lastfm.Api, file string) error {
	var f []byte
	var err error
	if f, err = os.ReadFile(file); err != nil {
		return err
	}
	sl := &scrobbleList{}
	json := jsoniter.Config{TagKey: schema}.Froze()
	if err := json.Unmarshal(f, sl); err != nil {
		return err
	}

	fmt.Printf("Total of %d Songs To be scrobbled.\n", len(*sl))
	v := *sl
	batchNo := 0
	totalScrobbles := 0
	totalIgnored := 0

	for k := 0; k <= len(v); k++ {
		acceptedScrobbles := 0
		ignored := 0
		a := lastfm.P{}
		batchNo++
		artists := []string{}
		albums := []string{}
		tracks := []string{}
		timestamps := []string{}
		count := 0
		for count < 50 {
			if k == len(*sl) {
				break
			}
			if lessThan(v[k].Milliseconds, threshold) || v[k].ArtistName == "" || v[k].TrackName == "" {
				k++
				continue
			}
			artists = append(artists, v[k].ArtistName)
			albums = append(albums, v[k].AlbumName)
			tracks = append(tracks, v[k].TrackName)
			timestamps = append(timestamps, time.Now().String())
			k++
			count++
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
	api := lastfm.New(key, secret)
	api.Login(username, password)
	fmt.Printf("Logged in as %s.\n", username)

	if len(os.Args) == 1 {
		fmt.Println("No File Provided")
		os.Exit(-1)
	}

	files := []string{}
	for k := range os.Args {
		if k != 0 {
			if os.Args[k] == "-s" {
				schema = "spotify"
			} else if os.Args[k] == "-t" {
				k++
				_t, err := strconv.ParseInt(os.Args[k], 10, 64)
				if err != nil {
					fmt.Printf("There was an error parsing the default scrobble threshold. You Provided %s", os.Args[k])
					panic(err)
				}
				threshold = int(_t)
			} else {
				files = append(files, os.Args[k])
			}
		}
	}

	for _, file := range files {
		fmt.Printf("Scrobbling From File: %s.\n", file)
		if err := _scrobble(api, file); err != nil {
			panic(err)
		}
	}
}
