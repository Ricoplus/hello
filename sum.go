package main

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 //发信号告诉任务管理者我已经计算完成了
}

const NCPU = 16

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) //勇于接受每个cpu完成信号

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	for i := 0; i < NCPU; i++ {
		<-c //获取到一个数据,表示一个cpu计算完成了
	}

	//到这里表示所有计算已经结束
}
