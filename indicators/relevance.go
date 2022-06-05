package indicators

import (
	"math"

	"github.com/peter9207/stocks/readers"
)

// var SignificanceWindow = float64(3)

func Relevance(data []readers.Data, sigma float64) (result []readers.Data) {

	result = []readers.Data{}

	sum := float64(0)
	sumSqErr := float64(0)

	for i, v := range data {
		if i != 0 {
			prevSD := math.Sqrt(sumSqErr / float64(i))
			diff := math.Abs(float64(v.Value-(sum/float64(i))) / prevSD)

			if diff > sigma {
				result = append(result, v)
			}
		}

		sum += v.Value
		mean := sum / float64(i+1)
		sumSqErr += math.Pow(v.Value-mean, 2)
	}
	return result

}
