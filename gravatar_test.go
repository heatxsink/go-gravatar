package gravatar

import (
	"testing"
	"fmt"
)

var (
	test_email_address = "helloworld000000@gmail.com"
)

func TestNew(t *testing.T) {
	fmt.Println("gravatar.New()")
	gg := New("wavatar", 200, "g", false)
	fmt.Println("gravatar.Dump()")
	gg.Dump()
}

func TestGetImageUrl(t *testing.T) {
	fmt.Println("gravatar.GetImageUrl()")
	gg := New("retro", 100, "g", false)
	image_url := gg.GetImageUrl(test_email_address)
	fmt.Println("Gravatar Image Url: ", image_url)
}