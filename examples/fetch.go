package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bmorton/go-livebarn"
)

// This example will download segments of video between 12:00pm and 2:30pm on
// March 17th, 2018 at Oakland Ice Center (NHL surface).
//
// In total, there will be 5 segments of 30 minutes:
//  - Part 1.mp4 - 12:00pm-12:30pm
//  - Part 2.mp4 - 12:30pm-1:00pm
//  - Part 3.mp4 - 1:00pm-1:30pm
//  - Part 4.mp4 - 1:30pm-2:00pm
//  - Part 5.mp4 - 2:00pm-2:30pm
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
