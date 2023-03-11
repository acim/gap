package gap

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	log.SetOutput(&buf)
	logger("foo", "bar", "baz")
	log.SetOutput(os.Stderr)

	if !strings.Contains(buf.String(), "foo bar baz") {
		t.Errorf("want foo bar baz; got %s", buf.String())
	}
}
