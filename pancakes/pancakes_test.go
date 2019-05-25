package pancakes_test

import (
	"github.com/cozzbp/pancakes/pancakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pancakes", func() {
	Context("FlipAtIndex", func() {
		It("Should flip single rune at index 0", func() {
			flipped := pancakes.FlipAtIndex([]rune("-"), 0)

			Expect(flipped).Should(Equal([]rune("+")))
		})
		It("Should flip at index", func() {
			flipped := pancakes.FlipAtIndex([]rune("-----+"), 4)

			Expect(flipped).Should(Equal([]rune("++++++")))
		})
	})

	Context("Flip", func() {
		It("Should not flip single happy-side up pancake", func() {
			p := pancakes.New("+")
			num := p.Flip()

			Expect(num).Should(Equal(0))
		})
		It("Should not flip multiple happy-side up pancake", func() {
			p := pancakes.New("+++++++++")
			num := p.Flip()

			Expect(num).Should(Equal(0))
		})
		It("Should flip single blank-side up pancake", func() {
			p := pancakes.New("-")
			num := p.Flip()
			Expect(num).Should(Equal(1))
		})
		It("Should flip pancakes 1 time", func() {
			p := pancakes.New("-+")
			num := p.Flip()
			Expect(num).Should(Equal(1))
		})
		It("Should flip pancakes 2 times", func() {
			p := pancakes.New("+-")
			num := p.Flip()
			Expect(num).Should(Equal(2))
		})
		It("Should flip pancakes 3 times", func() {
			p := pancakes.New("--+-")
			num := p.Flip()
			Expect(num).Should(Equal(3))
		})
		It("Should flip pancakes 4 times", func() {
			p := pancakes.New("+-+-")
			num := p.Flip()
			Expect(num).Should(Equal(4))
		})
		It("Should flip pancakes 1 time", func() {
			p := pancakes.New("------------------------")
			num := p.Flip()
			Expect(num).Should(Equal(1))
		})
		It("Should flip pancakes 0 times", func() {
			p := pancakes.New("+++++++++++++++++++++")
			num := p.Flip()
			Expect(num).Should(Equal(0))
		})
		It("Should flip pancakes 2 times", func() {
			p := pancakes.New("++++++++++++++++++++-")
			num := p.Flip()
			Expect(num).Should(Equal(2))
		})
	})
})
