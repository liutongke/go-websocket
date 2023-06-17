package chance

import (
	"math/rand"
)

type ModelChance struct {
	Id     int `json:"id"`
	Chance int `json:"chance"` //概率
}

func Choose(mrs []*ModelChance) int {
	// 定义两个map,一个是数据库里的数据ID==>数据名字,数据ID==>数据的数量
	pack := make(map[int]int)

	for _, v := range mrs {
		pack[v.Id] = v.Chance
	}

	return RedPack(pack)
}

func RedPack(pack map[int]int) (id int) {

	// 定义一个map,key对应数据ID,value是个定长为2的数组,下标0对应数值左范围,下标1对应数值右范围
	randArr := make(map[int][2]int)

	// sum为所有库存的和,并赋值randArr
	var sum int
	for k, v := range pack {
		var randArr1 [2]int
		randArr1[0] = sum
		sum += v
		randArr1[1] = sum
		randArr[k] = randArr1
	}

	// Golang随机得设置seed,否则每次随机的数是一样顺序的,
	//rand.Seed(time.Now().UnixNano())
	// 获取某个随机值
	s := rand.Intn(sum)

	// 查找随机出来的值在哪个范围内
	for m, n := range randArr {
		if s >= n[0] && s < n[1] {
			id = m
			break
		}
	}

	// 返回产品ID,一般后续操作是 返回奖品ID,然后将该产品在数据库里的库存-1
	// 如果你想永远概率都一样,那么每次抽中之后该函数传值 pack后的值不能变,需要另外一个参数比较当前奖品还有库存否,如果没有,返回其他值
	return
}
