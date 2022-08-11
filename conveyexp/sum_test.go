package conveyexp

import (
	"testing"

	"github.com/ag9920/go-test-demo/biz"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSumScore(t *testing.T) {

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
		initialSum := biz.GetSumScore(students)

		Convey("The result of GetSumScore should be equal to the sum of scores", func() {
			So(30, ShouldEqual, initialSum)
		})

		Convey("When any of the score is incremented", func() {
			students[0].Score++

			Convey("The sum should be greater by one", func() {
				newSum := biz.GetSumScore(students)
				So(newSum, ShouldEqual, initialSum+1)
			})
		})

		Convey("When any of the score is decremented", func() {
			students[1].Score--

			Convey("The sum should be less by one", func() {
				newSum := biz.GetSumScore(students)
				So(newSum, ShouldEqual, initialSum-1)
			})
		})
	})
}
