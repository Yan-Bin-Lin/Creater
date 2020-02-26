package serve

import (
	"testing"
)

func Testserve_SplitProject(t *testing.T) {
	type except struct {
		project []string
		blog    string
	}

	// build test table
	var splitTest = []struct {
		url     string // input
		except_ except // expected result
	}{
		{"/foo", except{[]string{"foo"}, ""}},
		{"/foo/bar", except{[]string{"foo"}, "bar"}},
		{"/foo/bar/blog", except{[]string{"foo", "bar"}, "blog"}},
	}

	for _, value := range splitTest {
		proj, blog := splitProject(value.url)
		// test project
		for i, v := range proj {
			if v != value.except_.project[i] {
				t.Errorf("splitProject(%s).project[%d] = %s; expected %s\n", value.url, i, v, value.except_.project[i])
			}
		}
		// test blog
		if blog != value.except_.blog {
			t.Errorf("splitProject(%s).blog = %s; expected %s\n", value.url, blog, value.except_.blog)
		}
	}
}
