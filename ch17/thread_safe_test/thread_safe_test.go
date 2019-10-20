package thread_safe_test

import (
	"fmt"
	"sync"
	"testing"
)

/*
一个线程安全的操作
定义了一个Mutex类型的mut
其中defer相当于finally，不管在什么情况下都会在defer里释放锁
*/
/*
对于RW操作的时候尽量使用RWlock，go里面的rwlock里的读锁是不互斥的，
当锁定一个读的时候，另一个读锁也可以进入，而写锁是互斥的，
这样就比Java里的读写都互斥提高了效率
*/
func TestThreadSafe(t *testing.T) {
	count := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
