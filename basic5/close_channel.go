package main

import "fmt"

func main() {
ch := make(chan int)


go func() {
	for i:=0;i<6;i++{
		ch <- i
	}

	close(ch)
}()

	//for   {
	//	if data, has := <- ch; has == true{
	//		fmt.Println("zhu go dudao ",data)
	//	}else {
	//		break
	//	}
	//}

	for data := range ch{  //ch有数据可读 保存在data ，1 无数据 对端没关闭，阻塞等待， 2 对端关闭close， 结束range
		fmt.Println(data)
	}

}
