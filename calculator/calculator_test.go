package calculator_test

import (
	"calc/calculator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	var (
		x int
		y int
	)

	BeforeEach(func() {
		x = 10
		y = 20
	})

	Describe("Test Plus", func() {
		Context("Plus test result", func() {
			It("should be 30", func() {
				Expect(calculator.Plus(x, y)).To(Equal(30))
			})
		})
	})

	Describe("Test Minus", func() {
		Context("Minus test result1", func() {
			It("should be -10", func() {
				Expect(calculator.Minus(x, y)).To(Equal(-10))
			})
		})

		Context("Minus test result2", func() {
			It("should be 10", func() {
				Expect(calculator.Minus(y, x)).To(Equal(10))
			})
		})
	})

	Describe("Test Multiply", func() {
		Context("Multiply test result", func() {
			It("should be 200", func() {
				Expect(calculator.Multiply(x, y)).To(Equal(200))
			})
		})
	})

	Describe("Test Pow", func() {
		Context("Pow test result", func() {
			It("should be over 7766279631452241920", func() {
				Expect(calculator.Pow(x, y)).To(Equal(7766279631452241920))
			})
		})
	})

	Describe("Test Divide", func() {
		Context("Divide test result1", func() {
			It("should be 0", func() {
				Expect(calculator.Devide(x, y)).To(Equal(0))
			})
		})

		Context("Divide test result2", func() {
			It("should be 2", func() {
				Expect(calculator.Devide(y, x)).To(Equal(2))
			})
		})
	})
})
