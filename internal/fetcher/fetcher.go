package fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/mandriota/what-anime-tui/internal/config"
)

type Fetcher struct {
	payload *bytes.Buffer
	
	cfg config.FetcherConfig
}

func New(cfg config.FetcherConfig) Fetcher {
	return Fetcher{
		payload: bytes.NewBuffer(nil),
		cfg: cfg,
	}
}

func decode(dst *Response, resp *http.Response) error {
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return fmt.Errorf("error while decoding response: %v", err)
	}

	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf("error while closing response: %v", err)
	}

	if dst.Error != "" {
		return fmt.Errorf("restAPI error: %s", dst.Error)
	}

	return nil
}

func writeImagePayload(payload io.Writer, r io.Reader, fpath string) (contentType string, err error) {
	mwriter := multipart.NewWriter(payload)
	fImg, err := mwriter.CreateFormFile("image", filepath.Base(fpath))
	if err != nil {
		return "", fmt.Errorf("error while creating form file: %v", err)
	}
	if _, err := io.Copy(fImg, r); err != nil {
		return "", fmt.Errorf("error while copying: %v", err)
	}
	if err := mwriter.Close(); err != nil {
		return "", fmt.Errorf("error while closing multipart.Writer: %v", err)
	}
	return mwriter.FormDataContentType(), nil
}

func (f Fetcher) FetchByFile(dst *Response, path string) error {
	fs, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error while openning file: %v", err)
	}
	defer fs.Close()

	f.payload.Reset()
	contentType, err := writeImagePayload(f.payload, fs, path)
	if err != nil {
		return fmt.Errorf("error while writing payload: %v", err)
	}

	resp, err := http.Post(f.cfg.ApiUrlByFile, contentType, f.payload)
	if err != nil {
		log.Fatalln("error while sending post request:", err)
	}

	return decode(dst, resp)
}

func (f Fetcher) FetchByURL(dst *Response, path string) error {
	resp, err := http.Get(strings.Replace(f.cfg.ApiUrlByUrl, "{{ .Path }}", url.QueryEscape(path), 1))
	if err != nil {
		return fmt.Errorf("error while fetching: %v", err)
	}

	return decode(dst, resp)
}
