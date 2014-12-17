go-gravatar
===========

A simple gravatar api wrapper.

Example
-------
```go
import(
	"fmt"
	"github.com/heatxsink/go-gravatar"
)

func main() {
	g := gravatar.New("retro", 100, "g", false)
	url := g.GetImageUrl("someones@email.com")
	fmt.Println("Gravatar Image Url: ", url)
}
```