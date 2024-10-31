package think

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	faker "github.com/bxcodec/faker/v4"
	"math/rand"
	"time"
)

// DisplayObject 表示一个显示场景中的对象
type DisplayObject struct {
	ID          string  // 对象的唯一标识符
	PositionX   float64 // X轴位置
	PositionY   float64 // Y轴位置
	Width       float64 // 对象的宽度
	Height      float64 // 对象的高度
	Rotation    float64 // 旋转角度
	Color       string  // 对象的颜色
	Opacity     float64 // 对象的透明度 (0.0 - 1.0)
	Visible     bool    // 是否可见
	LayerIndex  int     // 对象在场景中的层次位置
	Interactive bool    // 是否可以交互
}

// DisplayObjectList 表示 DisplayObject 的列表
type DisplayObjectList []*DisplayObject

// 使用 faker 库生成随机的 DisplayObject 实例
func NewRandomDisplayObject() *DisplayObject {

	return &DisplayObject{
		ID:          faker.UUIDDigit(),
		PositionX:   rand.Float64() * 1920,    // 随机X位置在屏幕宽度内
		PositionY:   rand.Float64() * 1080,    // 随机Y位置在屏幕高度内
		Width:       rand.Float64()*100 + 10,  // 随机宽度，10-110之间
		Height:      rand.Float64()*100 + 10,  // 随机高度，10-110之间
		Rotation:    rand.Float64() * 360,     // 随机旋转角度，0-360度
		Color:       generateRandomHexColor(), // 随机十六进制颜色
		Opacity:     rand.Float64(),           // 随机透明度，0.0-1.0
		Visible:     rand.Intn(2) == 1,        // 随机可见性
		LayerIndex:  rand.Intn(10),            // 随机层次位置
		Interactive: rand.Intn(2) == 1,        // 随机交互性
	}
}

func generateRandomHexColor() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// GenerateRandomDisplayObjects 生成指定数量的随机 DisplayObject
func GenerateRandomDisplayObjects(count int) DisplayObjectList {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	var objects DisplayObjectList
	for i := 0; i < count; i++ {
		objects = append(objects, NewRandomDisplayObject())
	}
	return objects
}

type WhereDisplayObject func(*DisplayObject) bool

func (p DisplayObjectList) WhereFn(fns ...WhereDisplayObject) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(func(in *DisplayObject) bool {
		for _, fn := range fns {
			if !fn(in) {
				return false
			}
		}
		return true
	}).ToSlice(&l)
	return l
}

func PureWherePositionYGte(positionY float64) WhereDisplayObject {
	return func(x *DisplayObject) bool {
		return x.PositionY >= positionY
	}
}

func PureWherePositionXGte(positionX float64) WhereDisplayObject {
	return func(x *DisplayObject) bool {
		return x.PositionX >= positionX
	}
}

func PureWhereLayerIndexGte(layerIndex int) WhereDisplayObject {
	return func(x *DisplayObject) bool {
		return x.LayerIndex >= layerIndex
	}
}

func PureWhereVisible(visible bool) WhereDisplayObject {
	return func(x *DisplayObject) bool {
		return x.Visible == visible
	}
}

func PureWhereOpacityGte(opacity float64) WhereDisplayObject {
	return func(x *DisplayObject) bool {
		return x.Opacity >= opacity
	}
}

func PureWhereInteractive(interactive bool) WhereDisplayObject {
	return func(x *DisplayObject) bool {
		return x.Interactive == interactive
	}
}

// WhereOpacity 根据透明度筛选对象
func (p DisplayObjectList) WhereOpacityGte(opacity float64) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(PureWhereOpacityGte(opacity)).ToSlice(&l)
	return l
}

// WhereVisible 根据是否可见筛选对象
func (p DisplayObjectList) WhereVisible(visible bool) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(PureWhereVisible(visible)).ToSlice(&l)
	return l
}

// WhereLayerIndex 根据层次位置筛选对象
func (p DisplayObjectList) WhereLayerIndexGte(layerIndex int) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(PureWhereLayerIndexGte(layerIndex)).ToSlice(&l)
	return l
}

// WherePosition 根据 X 和 Y 位置筛选对象
func (p DisplayObjectList) WherePositionXGte(positionX float64) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(PureWherePositionXGte(positionX)).ToSlice(&l)
	return l
}

// WherePosition 根据 X 和 Y 位置筛选对象
func (p DisplayObjectList) WherePositionYGte(positionY float64) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(PureWherePositionYGte(positionY)).ToSlice(&l)
	return l
}

// WhereInteractive 根据是否可以交互筛选对象
func (p DisplayObjectList) WhereInteractive(interactive bool) DisplayObjectList {
	var l []*DisplayObject
	linq.From(p).WhereT(PureWhereInteractive(interactive)).ToSlice(&l)
	return l
}

func main() {
	// 生成 10000 个随机的 DisplayObject
	objects := GenerateRandomDisplayObjects(10000)

	start := time.Now()

	l1 := objects.WhereLayerIndexGte(-1).
		WherePositionXGte(-1).
		WherePositionYGte(-1).
		WhereOpacityGte(-1).
		WhereInteractive(true).
		WhereVisible(true)
	middle := time.Now()

	l2 := objects.WhereFn(
		PureWhereLayerIndexGte(-1),
		PureWherePositionXGte(-1),
		PureWherePositionYGte(-1),
		PureWhereOpacityGte(-1),
		PureWhereInteractive(true),
		PureWhereVisible(true),
	)
	end := time.Now()

	fmt.Println(" domain where -> ", middle.Sub(start).String(), " size is :", len(l1))
	fmt.Println(" fn where -> ", end.Sub(middle).String(), " size is :", len(l2))

}
