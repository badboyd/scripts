package main

import (
	"strings"

	"git.chotot.org/corex/mirage/client"
	"git.chotot.org/corex/mirage/proc"
)

var (
	mirageClient *client.ImageClient
)

// why using int64 ???? unimaginable
type imageSetting struct {
	MaxWidth     int64 `mapstructure:"max_width"`
	MaxHeight    int64 `mapstructure:"max_height"`
	PTYMaxWidth  int64 `mapstructure:"pty_max_width"`
	PTYMaxHeight int64 `mapstructure:"pty_max_height"`
	Quality      int64 `mapstructure:"quality"`
	Watermark    bool  `mapstructure:"watermark"`
}

func init() {
	mirageClient = client.NewImageClient(mirageBaseURL, mirageSecret)
}

func getMirageImageOption(settings imageSetting) *proc.Options {
	opt := new(proc.Options)
	opt.Quality = settings.Quality
	opt.OutputType = proc.JPEG

	opt.Watermark = settings.Watermark

	opt.Width = settings.MaxWidth
	opt.Height = settings.MaxHeight

	return opt
}

func getMirageType(t string) int64 {
	switch t {
	case "jpg", "jpeg":
		return proc.JPEG
	case "png":
		return proc.PNG
	case "webp":
		return proc.WEBP
	default:
		return proc.JPEG
	}
}

func getMirageLink(opt *proc.Options) string {
	opt.Id = getOriginMediaID(opt.Id)
	return mirageBaseURL + mirageClient.GenerateQueryByOptions(opt)
}

func getOriginMediaID(id string) string {
	splitedID := strings.Split(id, "/")
	semiOriginID := splitedID[len(splitedID)-1]
	return strings.Replace(semiOriginID, ".jpg", "", -1)
}
