package main

import (
	"fmt"
	"os"
)
func main() {
	fout, err := os.Create("./write4.TXT")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\r\n", "Hello Go", i) //Sprintf格式化
		//写入文件
		fout.WriteString(outstr)            //string信息
		fout.Write([]byte("i love go\r\n")) //byte类型
	}
}
