package processor_test

import (
	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/common"
	. "github.com/rawfish-dev/upcode-academy-golang/dataprocessor/processor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Processor", func() {

	It("shoud handle zero cities", func() {
		// Input data
		data := []common.City{}

		// Expected output
		expected := map[string][]common.City{}

		// Call the function
		result := ProcessCities(data)

		// Assert the result is as expected
		Expect(result).To(Equal(expected))
		Expect(result).To(HaveLen(0))
	})

	It("should map a city to it's state id", func() {
		data := []common.City{
			{
				ID:      "322",
				Name:    "Sentosa",
				StateID: "123",
			},
		}

		expected := map[string][]common.City{
			"123": []common.City{
				{
					ID:      "322",
					Name:    "Sentosa",
					StateID: "123",
				},
			},
		}

		result := ProcessCities(data)

		Expect(result).To(Equal(expected))
	})

	It("should map cities to their state ids", func() {
		data := []common.City{
			{
				ID:      "322",
				Name:    "Sentosa",
				StateID: "123",
			},
			{
				ID:      "323",
				Name:    "Singapore",
				StateID: "123",
			},
		}

		expected := map[string][]common.City{
			"123": []common.City{
				{
					ID:      "322",
					Name:    "Sentosa",
					StateID: "123",
				},
				{
					ID:      "323",
					Name:    "Singapore",
					StateID: "123",
				},
			},
		}

		result := ProcessCities(data)

		Expect(result).To(Equal(expected))
	})
})
