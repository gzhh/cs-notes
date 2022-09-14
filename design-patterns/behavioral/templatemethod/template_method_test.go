package templatemethod

import "testing"

func TestDownloader(t *testing.T) {
	var httpDownloader Downloader = NewHTTPDownloader()
	httpDownloader.Download("https://example.com/abc.zip")

	var ftpDownloader Downloader = NewFTPDownloader()
	ftpDownloader.Download("ftp://example.com/abc.zip")
}
