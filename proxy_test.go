package m3u8proxy

import (
	"strings"
	"testing"

	"github.com/grafov/m3u8"
)

func TestProxy_Replace(t *testing.T) {
	proxy := &SimpleProxy{
		Prefix: "https://live.pamore.net/live/",
	}

	src := `#EXTM3U
#EXT-X-VERSION:3
#EXT-X-ALLOW-CACHE:NO
#EXT-X-MEDIA-SEQUENCE:1567658160
#EXT-X-TARGETDURATION:5
#EXTINF:4.055,
stream:lessions:134-1567658160.ts
#EXTINF:4.054,
stream:lessions:134-1567658161.ts
#EXTINF:4.056,
stream:lessions:134-1567658162.ts
`

	p, listType, err := m3u8.DecodeFrom(strings.NewReader(src), true)
	if err != nil {
		t.Error(err)
	}

	switch listType {
	case m3u8.MASTER:
		t.Error("Type of parsed result should be MEDIA, but got MASTER")
		break
	case m3u8.MEDIA:
		media := p.(*m3u8.MediaPlaylist)
		proxy.Replace(media)

		if media.Segments[0].URI != "https://live.pamore.net/live/stream:lessions:134-1567658160.ts" {
			t.Error("Replace uri error")
		}
		break
	}
}

func TestProxy_ReplaceURL(t *testing.T) {
	proxy := &SimpleProxy{
		Prefix: "https://live.pamore.net/live/",
	}
	url := "https://live.pamore.net/live/stream:lessions:134.m3u8?txSecret=472749dcc57cb2a090e1daf23db3d17a&txTime=5d7093b4"
	b, err := proxy.ReplaceURL(url)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}
