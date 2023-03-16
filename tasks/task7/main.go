package main

import (
	"fmt"
	"reflect"
	"time"
)

// Реализовать функцию, которая будет объединять один или более done-каналов в
// single-канал, если один из его составляющих каналов закроется.

// Очевидным вариантом решения могло бы стать выражение при использованием select,
// которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов,
// с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной
// функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.

func reflect_or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	var cases = make([]reflect.SelectCase, 0, len(channels))
	for _, c := range channels {
		// Собираем слайс кейсов
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv, // case <-Chan:
			Chan: reflect.ValueOf(c), // наш канал
		})
	}
	go func() {
		reflect.Select(cases) //ожидает пока не выполнится один из кейсов
		result <- struct{}{}
	}()
	return result
}

func select_or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		// Если каналов нет то сразу Done
		if len(channels) == 0 {
			res <- struct{}{}
			return
		}
		// Проходимся по списку каналов пока какойнибудь не выполнится
		for {
			for i := 0; i < len(channels); i++ {
				select {
				case <-channels[i]:
					res <- struct{}{}
					return
				default:
				}
			}
		}
	}()
	return res
}

func gorutines_or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		// Если каналов нет то сразу Done
		if len(channels) == 0 {
			res <- struct{}{}
			return
		}

		// Запускаем горутины ожидающие сигнала,
		// при желании можно релизовать отключение горутин, при выполнении хотя бы одной, через контекст
		for i := 0; i < len(channels); i++ {
			go func(out *chan interface{}, in *<-chan interface{}) {
				<-*in
				*out <- struct{}{}
			}(&res, &channels[i])
		}
	}()
	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-gorutines_or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))

	start = time.Now()
	<-reflect_or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))

	start = time.Now()
	<-select_or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
