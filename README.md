# Mukhachitram

Mukhachitram (ముఖచిత్రం in [Telugu](https://en.wikipedia.org/wiki/Telugu_language)) is a little script that fetches album arts for music folders and saves them as an image file inside them.

### Why?
I wrote this purely for learning Golang.

### Pre-requisities
You need a LastFM API key. If you don't have one, create it from here: https://www.last.fm/api/account/create

### Usage
1. Clone this repo and build the executable by running `go build` inside the root directory.
*or*
Download the executable file from the releases.
2. Place the executable file inside the directory that contains all your music folders.
3. Run `./mukhachitram -apikey "yourapikeyhere"` (without the actual quotes)

### To-do
- [ ] Optimize using Go routines
- [ ] Improve error messages and status indicators
- [ ] Refactor
- [ ] Support for choosing a service other than LastFM

### How to contribute?
Please open an issue if you have a feature request or if you can come across any issue. Thank you.

## License
Code in this repository is licensed under [MIT](https://mit-license.org).
