package util

import (
	"fmt"
	"os"
)

/*
 @Author: zhijian
 @Date: 2021/5/31 10:07
 @Description:
*/

//判断文件或文件夹是否存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
