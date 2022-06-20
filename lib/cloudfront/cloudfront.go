package cloudfront

import (
	"fmt"
	"net/url"
	"strings"
)

// GeneratePathFromS3Url is
func GeneratePathFromS3Url(path, env string) string {
	urls := strings.Split(path, ",")
	for i := range urls {
		if urls[i] == "" {
			if i == 0 {
				return ""
			}
			continue
		}
		urls[i] = url.PathEscape(urls[i])
		urls[i] = fmt.Sprint(
			"https://",
			env,
			".cloudfront.net/",
			urls[i],
		)
	}
	return strings.Join(urls, ",")
}

// GenerateCloudFrontURL is
func GenerateCloudFrontURL(s3URL, env string) string {
	urls := strings.Split(s3URL, ",")
	for i := range urls {
		urlFormat := fmt.Sprint(
			"https://",
			env,
			".cloudfront.net/",
		)
		urls[i] = urls[i][len(urlFormat):]
		urls[i], _ = url.PathUnescape(urls[i])
	}
	return strings.Join(urls, ",")
}
