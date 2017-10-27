/* グループ化していない宣言方法
package main

import "fmt"
import "os"


const i := 100
const pi = 3.1415
const prefix = "Go_"


var i int
var pi float32
var prefix string
*/

package main

import (
	"fmt"
)

const (
	i      = 100
	pi     = 3.1415
	prefix = "Go_"
	x      = iota // x == 0
	y      = iota // y == 1
	z      = iota // z == 2
	w             // w == 3 。iotaを宣言しなくても連続している場合デフォでiotaになる。
)

const v = iota // グルーピングの意味はここにある。constが宣言されるたびにiotaは0に初期化される。

const (
	e, f, g = iota, iota, iota //iotaは同一の行では同じ数字。0,0,0
)

var (
	j         int
	ki        float32
	preprefix string
)

func main() {
	fmt.Print(i)

	//iota列挙型

}
