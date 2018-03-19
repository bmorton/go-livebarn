# go-livebarn

`go-livebarn` is an attempt at creating a client library for interacting with
[LiveBarn](https://livebarn.com) through inspecting its browser-based
interactions with the server.

This library is a work-in-progress.

## Example

The example provided in `examples/fetch.go` demonstrates how to fetch URLs to
download video segments from an ice arena at a certain time.

```go
func main() {
	client := livebarn.New("my-token-here", "my-uuid-here")
	client.DebugMode = true

	oaklandIceNHL := &livebarn.Surface{UUID: "dff0fc40649943109e4ddab3118f3da2"}

	resp, _ := client.GetMedia(
		oaklandIceNHL,
		&livebarn.DateRange{
			Start: time.Date(2018, 03, 17, 12, 00, 00, 00, time.Local),
			End:   time.Date(2018, 03, 17, 14, 00, 00, 00, time.Local),
		},
	)

	for i, part := range resp.Result {
		filename := fmt.Sprintf("Part %d.mp4", i+1)
		log.Printf("Downloading %s...\n", filename)
		url, _ := client.GetMediaDownload(part.URL)

		err := livebarn.DownloadFile(filename, url.Result.URL)
		if err != nil {
			panic(err)
		}
	}
}
```

## `ffmpeg` Bonus

This little snippet will concatenate multiple video files into one:

```
$ ffmpeg -i "concat:input1.mp4|input2.mp4|input3.mp4" -c copy output.mp4
```
