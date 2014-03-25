package seamdb

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBolt(t *testing.T) {
	db := &SB{}

	Convey("Open", t, func() {
		err := db.Open()
		So(err, ShouldEqual, nil)

		Convey("Put a bunch of objects", func() {
			e1 := db.Put("main", "hi", []byte{1, 2, 3})
			e2 := db.Put("main", "bye", []byte{4, 5, 6})
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

		Convey("Get those objects", func() {
			v1, _ := db.Get("main", "hi")
			v2, _ := db.Get("main", "bye")
			So(v1, ShouldResemble, []byte{1, 2, 3})
			So(v2, ShouldResemble, []byte{4, 5, 6})
		})

		Convey("Delete those objects", func() {
			e1 := db.Del("main", "hi")
			e2 := db.Del("main", "bye")
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

		Reset(func() {
			db.Close()
		})
	})
}

//--

func TestBoltSetInit(t *testing.T) {
	db := &SB{}
	db.Open()
	defer db.Close()

	for _, d := range testData {
		err := db.Put("main", string(d[0]), d[1])
		if err != nil {
			panic(err)
		}
	}
}

func TestBoltSetGet(t *testing.T) {
	db := &SB{}
	db.Open()
	defer db.Close()

	for _, d := range testData {
		key := string(d[0])
		v, err := db.Get("main", key)
		if err != nil {
			panic(err)
		}
		_ = v
		fmt.Println("SB", key, len(v))
	}
}

func TestBoltSetDel(t *testing.T) {
	db := &SB{}
	db.Open()
	defer db.Close()

	for _, d := range testData {
		key := string(d[0])
		err := db.Del("main", key)
		if err != nil {
			panic(err)
		}
		fmt.Println("SB", key)
	}
}

//--

// func BenchmarkBoltPut(b *testing.B) {
// 	db := &SB{}
// 	db.Open()
// 	defer db.Close()

// 	for n := 0; n < b.N; n++ {
// 		for _, d := range testData {
// 			err := db.Put("main", string(d[0]), d[1])
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// }

// func BenchmarkBoltGet(b *testing.B) {
// 	db := &SB{}
// 	db.Open()
// 	defer db.Close()

// 	for n := 0; n < b.N; n++ {
// 		for _, d := range testData {
// 			key := string(d[0])
// 			v, err := db.Get("main", key)
// 			if err != nil {
// 				panic(err)
// 			}
// 			_ = v
// 			fmt.Println("SB", key, len(v))
// 		}
// 	}
// }
