/*
 * File: \src\ch9\memo1\main.go                                                *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/20 , 12:52:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 13:39:27                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package memo

import "sync"

//the result of calling Func was stored in Memo
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (any, error)

type result struct {
	value any
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: map[string]result{}}
}

//非并发安全
func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]

	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()

	return res.value, res.err
}
