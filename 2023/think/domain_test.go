package think_test

import (
	"figoxu/think"
	"fmt"
	"testing"
	"time"
)

func TestDomain(t *testing.T) {
	objects := think.GenerateRandomDisplayObjects(1000000)

	start := time.Now()

	l1 := objects.WhereLayerIndexGte(-1).
		WherePositionXGte(-1).
		WherePositionYGte(-1).
		WhereOpacityGte(-1).
		WhereInteractive(true).
		WhereVisible(true)
	middle := time.Now()

	l2 := objects.WhereFn(
		think.PureWhereLayerIndexGte(-1),
		think.PureWherePositionXGte(-1),
		think.PureWherePositionYGte(-1),
		think.PureWhereOpacityGte(-1),
		think.PureWhereInteractive(true),
		think.PureWhereVisible(true),
	)
	end := time.Now()

	fmt.Println(" domain where -> ", middle.Sub(start).String(), " size is :", len(l1))
	fmt.Println(" fn where -> ", end.Sub(middle).String(), " size is :", len(l2))
	//domain where ->  3.116921916s  size is : 250668
	//fn where ->  619.985834ms  size is : 250668
	//	可见,在数据量大的情况下，单次循环更节省时间，好处大于链式调用
}
