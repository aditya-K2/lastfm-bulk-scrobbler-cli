# Scrobble in Bulk to LastFM

**Note:** The Daily limit of 2800 is forced by the API.
Also take in consideration this will change your scrobble history and will drastically change your stats.

## Installing

#### Downloading the Binary.

Download the binary from the [release](https://github.com/aditya-K2/lastfm-bulk-scrobbler-cli/releases) section.

#### Building

###### Requires root access to install the binary.

```sh
cd /tmp && git clone https://github.com/aditya-K2/lastfm-bulk-scrobbler-cli && cd lastfm-bulk-scrobbler-cli && GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw" go build -v && sudo install -D last-fm-bulk-scrobbler -t "/usr/bin/"
```

## Usage

```sh
$ last-fm-bulk-scrobbler file1.json file2.json file3.json # uses the normal schema for json file.
$ last-fm-bulk-scrobbler -s endsong_1.json endsong_2.json # uses the spotify extended history schema for json file.
$ last-fm-bulk-scrobbler -s endsong_1.json endsong_2.json -t 180000 # uses the spotify extended history schema for json file with scrobble threshold of 180000 milliseconds (3 minutes)
```

## Format for JSON file.

This is the default schema used if `-s` flag is not specified.

```json
[
  {
    "artistName": "JPEGMAFIA",
    "albumName": "I Might Vote 4 Donald Trump",
    "trackName": "I Might Vote 4 Donald Trump",
  },
  {
    "artistName": "JPEGMAFIA",
    "albumName": "I Might Vote 4 Donald Trump",
    "trackName": "I Might Vote 4 Donald Trump",
  },
]
```
## Spotify Extended History

By default you can directly use your spotify extended history to scrobble to LastFM by specifying the `-s` flag. If you want to clean up your spotify data. Please follow [this](https://docs.google.com/document/d/1IhFMol3wZs24uKnh2rbxHpLaxhETcfB8KqzYIkEW_iM/edit#heading=h.vci5eys83lyn) very helpful guide *(Read the Prepping Your Data Section)*.

## Scrobble Threshold

`ms_played` in case of spotify extended history should be greater than `threshold` to be eligible for being scrobbled.
`threshold` is only allowed for spotify schema (i.e when -s flag is used), else it is ignored.
