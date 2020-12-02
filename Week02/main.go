package Week02

import (
	"fmt"
	"main/dao"
)

func main() {

	userInfo, error := dao.GetUserName(1)
	if error != nil {
		//返回notFound状态码
	}

	fmt.Println(userInfo)
}
