package main

import "log"

type kuma struct {
	name          string
	WakeUpHandler AngryFn
	EatHandler    HungryFn
}

type AngryFn func(int)
type HungryFn func(int, string)

func (k *kuma) HandleWake(fn func(int)) {
	k.WakeUpHandler = fn
}

func (k *kuma) HandleEat(fn func(int, string)) {
	k.EatHandler = fn
}

func main() {

	garyKuma := &kuma{
		name:       "AloHa",
		EatHandler: func(int, string) {}, // 基本上要init 比較保險
		// WakeUpHander: func(int) {},
	}

	garyKuma.HandleEat(func(action int, food string) {
		log.Println("action:", action, "thank you for your ", food)
	})

	if garyKuma.WakeUpHandler == nil {
		log.Println("wake up handler is nil !")

		// 沒有assign 東西，garyKuma.WakeUpHandler(5)會造成nil pointer dereference
		garyKuma.WakeUpHandler = func(int) {}
		garyKuma.WakeUpHandler(5)
	}

	// 丟肉
	garyKuma.EatHandler(0, "meat")

	// 變換掛載的handler
	changeHandler(garyKuma)

	// 丟肉
	garyKuma.EatHandler(0, "meat")
}

func changeHandler(k *kuma) {
	k.HandleEat(func(action int, food string) {
		log.Println(k.name+" don't like this :", food)
	})
}

// result:

// wake up handler is nil !
// action: 0 thank you for your  meat
// AloHa don't like this : meat
