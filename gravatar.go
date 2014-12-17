package gravatar

import (
	"fmt"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

type Gravatar struct {
	Default string
	Size int
	Rating string
	SslFlag bool
}

func New(default_image string, size int, rating string, ssl_flag bool) *Gravatar {
	g := Gravatar{}
	g.Default = default_image
	g.Size = size
	g.Rating = rating
	g.SslFlag = ssl_flag
	return &g
}

func (g *Gravatar) GetImageUrl(email_address string) string {
	scheme := "http"
	if g.SslFlag {
		scheme = "https"
	}
	email_address_normalized := strings.ToLower(email_address)
	hasher := md5.New()
	hasher.Write([]byte(email_address_normalized))
	email_hash := hex.EncodeToString(hasher.Sum(nil))
	default_variable := url.QueryEscape(g.Default)
	rating_variable := url.QueryEscape(g.Rating)
	size_variable := g.Size
	gravatar_image_url := fmt.Sprintf("%s://www.gravatar.com/avatar/%s?d=%s&s=%d&r=%s", scheme, email_hash, default_variable, size_variable, rating_variable)
	return gravatar_image_url
}

func (g *Gravatar) Dump() {
	fmt.Println("Gravatar")
	fmt.Println("--------")
	fmt.Println("\tDefault:     ", g.Default)
	fmt.Println("\tSize:        ", g.Size)
	fmt.Println("\tRating:      ", g.Rating)
	fmt.Println("\tSslFlag:     ", g.SslFlag)
}