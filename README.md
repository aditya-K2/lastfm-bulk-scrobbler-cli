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
$ lastfm-bulk-scrobbler file1.json file2.json file3.json
```

## Format for JSON file.

Following Schema for JSON file is expected. `time` is not required as it will be overridden with the current time.


```json
[
  {
    "artistName": "JPEGMAFIA",
    "albumName": "I Might Vote 4 Donald Trump",
    "trackName": "I Might Vote 4 Donald Trump",
    "time": "2022-02-19T18:28:00Z"
  },
  {
    "artistName": "JPEGMAFIA",
    "albumName": "I Might Vote 4 Donald Trump",
    "trackName": "I Might Vote 4 Donald Trump",
    "time": "2022-02-19T18:28:00Z"
  },
]
```
## Spotify Extended History

If you have your `spotify` extended history. You can convert it into the compatible format [here](https://lilnasy.github.io/scribblyscrobbly/). Details of how to use this website are available [here.](https://docs.google.com/document/d/1IhFMol3wZs24uKnh2rbxHpLaxhETcfB8KqzYIkEW_iM/edit#heading=h.vci5eys83lyn) *(Read the Prepping Your Data Section)*
