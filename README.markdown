m3u8-proxy
===

这个项目的初衷是因为腾讯云直播 HLS 拉流的播放列表中的路径都是相对路径，这在原生
的 IOS HTML5 页面中会有播放的问题，所以用这个项目代理一下。

## Usage

```
修改腾讯云直播 m3u8 拉流中的播放列表数据。主要是需要将相对路径转换成绝对路径

Usage:
  m3u8-proxy [command]

Available Commands:
  help        Help about any command
  serve       启动服务

Flags:
      --config string   config file (default is $HOME/.m3u8-proxy.yaml)
  -h, --help            help for m3u8-proxy
  -t, --toggle          Help message for toggle

Use "m3u8-proxy [command] --help" for more information about a command.
```
