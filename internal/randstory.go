package internal

import (
	"log"
	"math"
	"math/rand"
	"time"
)

// RandStory は建物階数をランダムに設定する
func RandStory() (story int) {
	seed := time.Now().UnixNano()
	log.Println("seed=", seed)
	r := rand.New(rand.NewSource(seed))
	log.Println("r=", r)
	val := r.Float64()
	log.Println("val=", val)
	// s = int(val)
	// log.Println("s=", s)

	if val < 0.9 {
		spr := 0.700458600675584 * math.Exp(2.23524816406936*val)
		log.Println("spr=", spr)
		story = int(spr + 0.5)
		log.Println("story1=", story)
	} else {
		spr1 := 0.700458600675584 * math.Exp(2.23524816406936*0.9)
		spr2 := 3.35035038917538 * math.Exp(1.36764285404573*(val-0.9))
		spr3 := 3.35035038917538 * math.Exp(1.36764285404573*0.0)
		sprt := spr1 + spr2 - spr3
		log.Println("sprt=", sprt)
		story = int(sprt + 0.5)
		log.Println("story2=", story)
	}

	return story
}
