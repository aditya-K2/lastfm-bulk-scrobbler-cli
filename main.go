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
	secret = os.Getenv("LASTFM_SECRET")
	key    = os.Getenv("LASTFM_KEY")
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
	for _k := 0; _k <= len(v); _k++ {
		acceptedScrobbles := 0
		a := lastfm.P{}
		batchNo++
		for _i := 0; _i < 50; _i++ {
			if _k == len(*sl) {
				break
			}
			v := *(sl)
			a[fmt.Sprintf("artist[%d]", _i)] = v[_k].ArtistName
			a[fmt.Sprintf("album[%d]", _i)] = v[_k].AlbumName
			a[fmt.Sprintf("track[%d]", _i)] = v[_k].TrackName
			a[fmt.Sprintf("timestamp[%d]", _i)] = time.Now().String()
			a[fmt.Sprintf("chosenByUser[%d]", _i)] = 1
			_k++
		}
		if result, err := api.Track.Scrobble(a); err != nil {
			fmt.Println(result)
			return err
		} else {
			if val, _err := strconv.Atoi(result.Accepted); _err != nil {
				return _err
			} else {
				acceptedScrobbles += val
			}
		}
		fmt.Printf("\r Batch %d (contains %d entries). Out of which %d were accepted.", batchNo, len(a)/5, acceptedScrobbles)
	}
	return nil
}

func main() {
	api := lastfm.New(key, secret)
	api.Login(os.Getenv("LASTFM_USERNAME"), os.Getenv("LASTFM_PASSWORD"))

	for k := range os.Args {
		if k != 0 {
			fmt.Printf("Scrobbling From File: %s.\n", os.Args[k])
			if err := _scrobble(api, os.Args[k]); err != nil {
				panic(err)
			}
		}
	}
}
