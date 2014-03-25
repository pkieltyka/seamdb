package seamdb

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLevel(t *testing.T) {
	db := &SL{}

	Convey("Open", t, func() {
		err := db.Open()
		So(err, ShouldEqual, nil)

		Convey("Put a bunch of objects", func() {
			e1 := db.Put("hi", []byte{1, 2, 3})
			e2 := db.Put("bye", []byte{4, 5, 6})
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

		Convey("Get those objects", func() {
			v1, _ := db.Get("hi")
			v2, _ := db.Get("bye")
			So(v1, ShouldResemble, []byte{1, 2, 3})
			So(v2, ShouldResemble, []byte{4, 5, 6})
		})

		Convey("Delete those objects", func() {
			e1 := db.Del("hi")
			e2 := db.Del("bye")
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

		Reset(func() {
			db.Close()
		})
	})
}

//--

// func BenchmarkLevelPut(b *testing.B) {
// 	db := &SL{}
// 	db.Open()
// 	defer db.Close()

// 	for n := 0; n < b.N; n++ {
// 		for _, d := range testData {
// 			err := db.Put(string(d[0]), d[1])
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// }

// func BenchmarkLevelGet(b *testing.B) {
// 	db := &SL{}
// 	db.Open()
// 	defer db.Close()

// 	for n := 0; n < b.N; n++ {
// 		for _, d := range testData {
// 			key := string(d[0])
// 			v, err := db.Get(key)
// 			if err != nil {
// 				panic(err)
// 			}
// 			_ = v
// 			fmt.Println("SL", key, len(v))
// 		}
// 	}
// }

// func BenchmarkLevelIterate(b *testing.B) {
// 	db := &SL{}
// 	db.Open()
// 	defer db.Close()

// 	for n := 0; n < b.N; n++ {
// 		i := db.Db.NewIterator(nil, nil)
// 		for i.Next() {
// 			k := i.Key()
// 			v := i.Value()
// 			fmt.Println(string(k), len(v))

// 			err := db.Del(string(k))
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 		i.Release()
// 	}
// }
