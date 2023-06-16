package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	// "strings"
)

type Download struct {
	URL      string
	Hash     string
	CacheDir string
}

func (d *Download) Download() (string, error) {
	u, err := url.Parse(d.URL)
	if err != nil {
		return "", err
	}

	var filename string
	switch u.Scheme {
	case "http", "https", "ftp":
		filename, err = d.downloadHTTP(u)
	case "file":
		filename, err = d.downloadFile(u)
	default:
		return "", fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
	if err != nil {
		return "", err
	}

	if d.Hash != "" {
		if err = d.verifyHash(filename); err != nil {
			return "", err
		}
	}

	return filename, nil
}

func (d *Download) downloadHTTP(u *url.URL) (string, error) {
	filename := filepath.Base(u.Path)
	cacheFilename := filepath.Join(d.CacheDir, filename)

	// check if file is already cached
	if _, err := os.Stat(cacheFilename); err == nil {
		fmt.Printf("Using cached file: %s\n", cacheFilename)
		return cacheFilename, nil
	}

	resp, err := http.Get(d.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	f, err := os.Create(cacheFilename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	size, err := io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Printf("Downloaded %s (%d bytes)\n", filename, size)
	return d.extractFile(cacheFilename)
}

func (d *Download) downloadFile(u *url.URL) (string, error) {
	filename := u.Path
	cacheFilename := filepath.Join(d.CacheDir, filename)

	// check if file is already cached
	if _, err := os.Stat(cacheFilename); err == nil {
		fmt.Printf("Using cached file: %s\n", cacheFilename)
		return cacheFilename, nil
	}

	fi, err := os.Stat(filename)
	if err != nil {
		return "", err
	}

	if fi.IsDir() {
		return "", fmt.Errorf("%s is a directory", filename)
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	cf, err := os.Create(cacheFilename)
	if err != nil {
		return "", err
	}
	defer cf.Close()

	size, err := io.Copy(cf, f)
	if err != nil {
		return "", err
	}

	fmt.Printf("Downloaded %s (%d bytes)\n", filename, size)
	return d.extractFile(cacheFilename)
}

func (d *Download) verifyHash(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return err
	}

	if hex.EncodeToString(h.Sum(nil)) != d.Hash {
		return fmt.Errorf("hash mismatch: expected %s, got %s", d.Hash, hex.EncodeToString(h.Sum(nil)))
	}

	return nil
}

func (d *Download) extractFile(filename string) (string, error) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".zip":
		return d.extractZip(filename)
	case ".tar.gz", ".tgz":
		return d.extractTarGz(filename)
	default:
		return filename, nil
	}
}

func (d *Download) extractZip(filename string) (string, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return "", err
	}
	defer r.Close()

	var dest string
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		defer rc.Close()

		if dest == "" {
			dest = filepath.Join(d.CacheDir, f.Name)
			if f.FileInfo().IsDir() {
				if err = os.MkdirAll(dest, f.Mode()); err != nil {
					return "", err
				}
				continue
			}
			dest = filepath.Dir(dest)
		}

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(path, f.Mode()); err != nil {
				return "", err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return "", err
		}

		w, err := os.Create(path)
		if err != nil {
			return "", err
		}
		defer w.Close()

		if _, err = io.Copy(w, rc); err != nil {
			return "", err
		}
	}

	return dest, nil
}

func (d *Download) extractTarGz(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	var dest string
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if dest == "" {
			dest = filepath.Join(d.CacheDir, hdr.Name)
			if hdr.FileInfo().IsDir() {
				if err = os.MkdirAll(dest, hdr.FileInfo().Mode()); err != nil {
					return "", err
				}
				continue
			}
			dest = filepath.Dir(dest)
		}

		path := filepath.Join(dest, hdr.Name)
		if hdr.FileInfo().IsDir() {
			if err = os.MkdirAll(path, hdr.FileInfo().Mode()); err != nil {
				return "", err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return "", err
		}

		w, err := os.Create(path)
		if err != nil {
			return "", err
		}
		defer w.Close()

		if _, err = io.Copy(w, tr); err != nil {
			return "", err
		}
	}

	return dest, nil
}

func main() {
	downloads := []Download{
		Download{
			URL:      "https://example.com/file.zip",
			Hash:     "c7c6a3c6e7b7f276a9f4e6e0e8a8e0b9",
			CacheDir: "/tmp/cache",
		},
		Download{
			URL:      "ftp://example.com/file.tgz",
			Hash:     "9f5b9c0a7cc6c5e9c8d8c8c5e9c5e9c8",
			CacheDir: "/tmp/cache",
		},
		Download{
			URL:      "file:///path/to/file.txt",
			Hash:     "d41d8cd98f00b204e9800998ecf8427e",
			CacheDir: "/tmp/cache",
		},
	}

	for _, d := range downloads {
		if _, err := d.Download(); err != nil {
			fmt.Println(err)
		}
	}
}