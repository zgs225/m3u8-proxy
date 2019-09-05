/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	m3u8 "github.com/zgs225/m3u8-proxy"
	"log"
	"net/http"
	"time"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	Long:  `启动 m3u8 代理服务`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetDefault("timeout", 5)
		viper.SetDefault("addr", ":8080")

		realHost := viper.GetString("real_host")
		prefix := viper.GetString("prefix")

		if len(realHost) == 0 {
			log.Panic("缺少配置 real_host")
		}

		if len(prefix) == 0 {
			log.Panic("缺少配置 prefix")
		}

		s := &m3u8.HTTPServer{
			RealHost: realHost,
			Proxy: &m3u8.SimpleProxy{
				Prefix: prefix,
				Client: &http.Client{
					Timeout: time.Duration(viper.GetInt64("timeout")) * time.Second,
				},
			},
		}

		errc := make(chan error)
		go func() {
			errc <- http.ListenAndServe(viper.GetString("addr"), s)
		}()
		log.SetPrefix("[M3U8-Proxy]")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("HTTP service listening on ", viper.GetString("addr"))
		log.Panic(<-errc)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
