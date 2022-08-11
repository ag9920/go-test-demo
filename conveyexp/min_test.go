package conveyexp

import (
	"testing"

	"github.com/ag9920/go-test-demo/biz"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMin(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some students with scores", t, func() {
		students := []biz.Student{
			{
				ID:    1,
				Name:  "join",
				Score: 12,
			},
			{
				ID:    2,
				Name:  "michelle",
				Score: 13,
			},
			{
				ID:    3,
				Name:  "kelly",
				Score: 5,
			},
		}
		initialMin := biz.GetMinimumScore(students)

		Convey("The result of GetMinimumScore should be the minimum score", func() {
			So(5, ShouldEqual, initialMin)
		})

		Convey("When a new score becomes the minimum", func() {
			students[0].Score = 3

			Convey("The minimum score should be updated", func() {
				newMin := biz.GetMinimumScore(students)
				So(newMin, ShouldEqual, students[0].Score)
			})
		})
	})
}
