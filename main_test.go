package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRepeatedNumers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "No numbers repeated")

}

var _ = Describe("No numbers repeated", func() {
	Context("Test of numbers repeated given array of numbers", func() {
		It("check", func() {
			type Equivalences struct {
				InitialArray []int
				ResultArray  []int
			}

			var ArrayEquivalences = []Equivalences{
				Equivalences{
					InitialArray: []int{0, 0, 0, 0},
					ResultArray:  []int{0},
				},
				Equivalences{
					InitialArray: []int{2, 3, 5, 7, 11, 13},
					ResultArray:  []int{2, 3, 5, 7, 11, 13},
				},
				Equivalences{
					InitialArray: []int{2, 8, 9, 1, 1, 11, 3, 5, 7, 11, 13},
					ResultArray:  []int{1, 2, 3, 5, 7, 8, 9, 11, 13},
				},
			}

			for _, ArrayEquivalencesMap := range ArrayEquivalences {
				Expect(NoRepeated(ArrayEquivalencesMap.InitialArray)).Should(BeEquivalentTo(ArrayEquivalencesMap.ResultArray))
			}
		})
	})
})
