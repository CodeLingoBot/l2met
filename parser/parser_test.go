package parser

import (
	"bufio"
	"bytes"
	"github.com/ryandotsmith/l2met/bucket"
	"testing"
)

var parseTest = []struct {
	tname string
	in    string
	opts  options
	names []string
}{
	{
		"simple",
		`88 <174>1 2013-07-22T00:06:26-00:00 somehost name test - measure#hello=1 measure#world=1ms\n`,
		options{"user": []string{"u"}, "password": []string{"p"}},
		[]string{"hello", "world"},
	},
	{
		"legacy",
		`70 <174>1 2013-07-22T00:06:26-00:00 somehost name test - measure.hello=1\n`,
		options{"user": []string{"u"}, "password": []string{"p"}},
		[]string{"hello"},
	},
}

func TestBuildBuckets(t *testing.T) {
	for _, test := range parseTest {
		body := bufio.NewReader(bytes.NewBufferString(test.in))
		buckets := make([]*bucket.Bucket, 0)
		for b := range BuildBuckets(body, test.opts) {
			buckets = append(buckets, b)
		}
		if len(buckets) != len(test.names) {
			t.Errorf("test=%s actual-len=%d expected-len=%d\n",
				test.tname, len(buckets), len(test.names))
			t.FailNow()
		}
		for i := range test.names {
			if buckets[i].Id.Name != test.names[i] {
				t.Errorf("test=%s actual-name=%s expected-name=%s\n",
					test.tname, test.names[i], buckets[i].Id.Name)
			}
		}
	}
}
