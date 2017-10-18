/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-10-16 22:20:36
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-16 22:44:25
***********************************************/
package libs

import (
	"crypto/md5"
	"fmt"
)

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
