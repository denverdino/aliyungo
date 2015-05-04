package oss

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func (b *Bucket) Path(path string) string {
	return "/" + b.Name + "/" + path
}

func (b *Bucket) objectOp(method, path string, headers http.Header, data io.Reader) (httpResp *http.Response, err error) {
	httpResp, err = b.Client.Invoke(method, b.Path(path), data, headers)
	return
}

// PutObject creates object/updates with byte array
func (b *Bucket) Put(path string, data []byte, contentType string) error {
	headers := make(http.Header)
	body := bytes.NewBuffer(data)
	if contentType == "" {
		contentType = DefaultContentType
	}

	headers.Set("Content-Length", strconv.Itoa(len(data)))
	headers.Set("Content-Type", contentType)

	_, err := b.objectOp("PUT", path, headers, body)
	return err
}

// PutCopy puts a copy of an object given by the key path into bucket b using b.Path as the target key
func (b *Bucket) PutCopy(path string, source string, headers http.Header) (*CopyObjectResult, error) {
	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("x-oss-copy-source", url.QueryEscape(source))
	resp, err := b.objectOp("PUT", path, headers, nil)
	if err == nil {
		result := &CopyObjectResult{}
		b.Client.decodeResponse(resp, result)
		if err == nil {
			return result, nil
		}
	}
	return nil, err
}

// PutReader inserts an object into the bucket by consuming data
// from r until EOF.
func (b *Bucket) PutReader(path string, body io.Reader, length int64, contType string, headers http.Header) error {
	if headers == nil {
		headers = make(http.Header)
	}
	headers.Set("Content-Length", strconv.FormatInt(length, 10))
	headers.Set("Content-Type", contType)
	_, err := b.objectOp("PUT", path, headers, body)
	return err
}

// PutFile creates/updates object with file
func (b *Bucket) PutFile(path string, file *os.File) error {
	var contentType string
	if dotPos := strings.LastIndex(file.Name(), "."); dotPos == -1 {
		contentType = DefaultContentType
	} else {
		if mimeType := mime.TypeByExtension(file.Name()[dotPos:]); mimeType == "" {
			contentType = DefaultContentType
		} else {
			contentType = mimeType
		}
	}
	stats, err := file.Stat()
	if err != nil {
		log.Panicf("Unable to read file %s stats.", file.Name())
		return nil
	}

	return b.PutReader(path, file, stats.Size(), contentType, nil)
}

// GetObject retrieves object content
func (b *Bucket) Get(path string) (data []byte, err error) {

	body, err := b.GetReader(path)

	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(body)
	return
}

// GetObjectReader retrieves an object content as Reader
func (b *Bucket) GetReader(path string) (body io.ReadCloser, err error) {
	resp, err := b.GetResponse(path)
	if resp != nil {
		return resp.Body, err
	}
	return nil, err
}

// GetResponse retrieves an object from an OSS bucket,
func (b *Bucket) GetResponse(path string) (resp *http.Response, err error) {
	return b.GetResponseWithHeaders(path, make(http.Header))
}

// GetReaderWithHeaders retrieves an object from an OSS bucket
func (b *Bucket) GetResponseWithHeaders(path string, headers http.Header) (resp *http.Response, err error) {
	return b.objectOp("GET", path, headers, nil)
}

// DeleteObject deletes object
func (b *Bucket) Delete(path string) error {
	_, err := b.objectOp("DELETE", path, nil, nil)
	return err
}

// GetObjectMetadata gets object metadata with HEAD request
func (b *Bucket) Head(path string, headers http.Header) (httpResp *http.Response, err error) {
	return b.objectOp("HEAD", path, headers, nil)
}

// Exists checks whether or not an object exists on an OSS bucket using a HEAD request.
func (b *Bucket) Exists(path string) (exists bool, err error) {

	resp, err := b.Head(path, nil)

	if err != nil {
		// We can treat a 403 or 404 as non existance
		if e, ok := err.(*Error); ok && (e.StatusCode == 403 || e.StatusCode == 404) {
			return false, nil
		}
		return false, err
	}

	if resp.StatusCode/100 == 2 {
		exists = true
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
	return exists, err

}

// DelMulti removes up to 1000 objects from the S3 bucket.
func (b *Bucket) DeleteMultiple(objects Delete) error {
	doc, err := xml.Marshal(objects)
	if err != nil {
		return err
	}

	buf := makeXmlBuffer(doc)
	digest := md5.New()
	size, err := digest.Write(buf.Bytes())
	if err != nil {
		return err
	}

	headers := make(http.Header)

	headers.Set("Content-Length", strconv.Itoa(size))
	headers.Set("Content-MD5", base64.StdEncoding.EncodeToString(digest.Sum(nil)))
	headers.Set("Content-Type", "text/xml")

	_, err = b.objectOp("POST", "/?delete", headers, buf)

	return err

}

func makeXmlBuffer(doc []byte) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString(xml.Header)
	buf.Write(doc)
	return buf
}
