package pricehelper_test

import (
	"testing"

	"github.com/aizeNR/user-balance-service/pkg/pricehelper"
	"github.com/stretchr/testify/assert"
)

func TestPennyToRubles(t *testing.T) {
	t.Parallel()

	type test struct {
		input uint64
		want  float64
	}

	cases := []test{
		{
			input: 1512,
			want:  15.12,
		},
		{
			input: 1599,
			want:  15.99,
		},
		{
			input: 1500,
			want:  15.00,
		},
	}

	for _, tc := range cases {
		actual := pricehelper.PennyToRubles(tc.input)

		assert.Equal(t, tc.want, actual)
	}
}

func TestRublesToPenny(t *testing.T) {
	t.Parallel()

	type test struct {
		input float64
		want  uint64
	}

	cases := []test{
		{
			want:  1514,
			input: 15.14,
		},
		{
			want:  1599,
			input: 15.99,
		},
		{
			want:  1500,
			input: 15.00,
		},
	}

	for _, tc := range cases {
		actual := pricehelper.RublesToPenny(tc.input)

		assert.Equal(t, tc.want, actual)
	}
}
