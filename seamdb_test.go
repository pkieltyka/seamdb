package seamdb

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	// . "github.com/smartystreets/goconvey/convey"
)

var (
	testData = buildTestData(100)
)

func buildTestData(numRepeatSet int) [][2][]byte {
	files, _ := filepath.Glob("./testdata/images/*")
	setNum := len(files)
	if numRepeatSet == 0 {
		numRepeatSet = 1
	}
	data := make([][2][]byte, setNum*numRepeatSet)

	for i, filename := range files {
		d, _ := ioutil.ReadFile(filename)
		k := []byte(filepath.Base(filename))

		data[i] = [2][]byte{k, d}

		if numRepeatSet > 1 {
			for n := 0; n < numRepeatSet; n++ {
				nk := []byte(fmt.Sprintf("%s-%d", k, n))
				data[(setNum*n)+i] = [2][]byte{nk, d}
			}
		}
	}

	return data
}
