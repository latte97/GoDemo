package ChanDemo

import "fmt"

func ChanOpertion()  {
	//定义一个chan，这里单独做解释，就算chan里面的类型是指针类型的，通过chan一定是值传递
	UserChan := make(chan *map[string]string)
	//执行从chan写数据的方法
	go ChanInput(UserChan)


	//执行从chan出数据的方法
	ChanOutput(UserChan)

}
//NKbeQ#5HnWZPAx

//在chan中写入数据
func ChanInput(ch chan *map[string]string){
	fmt.Println("执行入chan方法")
	var a = map[string]string{"A":"a"}
	fmt.Printf("input----%p\n",&a)//打印传入chan前 ----内存地址
	ch <- &a
}

func ChanOutput(ch chan *map[string]string){
	fmt.Println("执行出chan方法")
	out_data := <- ch

	fmt.Printf("output---%p\n",&out_data) //打印从chan读取出来后---- 内存地址
}

func main()  {


}
