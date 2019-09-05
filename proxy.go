package m3u8proxy

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/grafov/m3u8"
)

// Proxy m3u8 文件代理
type Proxy interface {
	// Replace 将播放列表中的相对路径替换成绝对路径
	Replace(*m3u8.MediaPlaylist)
	// ReplaceURL 通过给定的 m3u8 文件链接，替换其中的内容
	ReplaceURL(url string) ([]byte, error)
}

// SimpleProxy
type SimpleProxy struct {
	Prefix string
	Client *http.Client
}

func (p *SimpleProxy) Replace(l *m3u8.MediaPlaylist) {
	for _, seg := range l.Segments {
		if seg != nil {
			seg.URI = p.Prefix + seg.URI
		}
	}
}

func (p *SimpleProxy) ReplaceURL(url string) ([]byte, error) {
	if p.Client == nil {
		p.Client = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("New request error: ", err)
		return nil, err
	}
	request.Header.Set("User-Agent", "AppleCoreMedia/1.0.0.16F203 (iPod touch; U; CPU OS 12_3_1 like Mac OS X; zh_cn)")
	request.Header.Set("Accept", "*/*")
	response, err := p.Client.Do(request)

	if err != nil {
		log.Println("Request real server error: ", err)
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Read from response error: ", err)
		return nil, err
	}
	log.Println("Response body is: ", string(b))
	entity, t, err := m3u8.DecodeFrom(bytes.NewBuffer(b), true)

	if err != nil {
		log.Println("Decode m3u8 content from real server error: ", err)
		return nil, err
	}

	switch t {
	case m3u8.MASTER:
		return b, nil
		break
	case m3u8.MEDIA:
		p.Replace(entity.(*m3u8.MediaPlaylist))
		return entity.Encode().Bytes(), nil
		break
	}

	return nil, errors.New("Never reached")
}
