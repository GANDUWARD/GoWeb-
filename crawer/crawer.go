package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Get(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页的body内容
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				result += string(buf[:n])
				break
			} else {
				fmt.Println("resp.Body.err =", err)
				break
			}
		}
		result += string(buf[:n])
	}
	fmt.Print(result)
	return
}

func SpiderPage(URL string, i int, page chan<- int) {
	fmt.Printf("正在爬取第%d个网页\n", i)
	//将所有网页内容爬取下来
	url := URL + strconv.Itoa((i-1)*50)
	result, err := Get(url)
	if err != nil {
		fmt.Println("http.Get err =", err)
		return
	}
	//把内容写入文件
	filename := "page" + strconv.Itoa(i) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err =", err)
		return
	}
	//写内容
	f.WriteString(result)
	//关闭文件
	f.Close()
	page <- i
}
func SpiderPage_forone(URL string) {
	//将所有网页内容爬取下来
	result, err := Get(URL)
	if err != nil {
		fmt.Println("http.Get err =", err)
		return
	}
	//把内容写入文件
	filename := "page" + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err =", err)
		return
	}
	//写内容
	f.WriteString(result)
	//关闭文件
	f.Close()
}

func Run(URL string, start, end int) {
	fmt.Printf("正在爬取%d页到%d页\n", start, end)
	//因为很有可能爬虫还没有结束下面的循环就已经结束了，所以这里就需要将数据传入通道
	page := make(chan int)
	for i := start; i <= end; i++ {
		//将page阻塞
		go SpiderPage(URL, i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page) //直接将面码传给点位符,值直接从管道里取出
	}
}
func main() {
	url := "http://www.xx671.com"
	/*var start, end int
	fmt.Printf("请输入起始页的数字>=1:>")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页的数字>=1:>")
	fmt.Scan(&end)
	Run(url, start, end)*/
	SpiderPage_forone(url)
}
