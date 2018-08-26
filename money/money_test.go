package money_test

import (
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitNewMoney(t *testing.T) {
	t.Run("Does not return error", func(t *testing.T) {
		_, err := money.NewMoney(1)
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitMoneyAdd(t *testing.T) {

	testCases := []struct {
		name         string
		initialValue int64
		addValue     int64
		wantOutput   string
		wantError    error
	}{
		{
			name:         "Adds money - negative input value",
			initialValue: 10,
			addValue:     -10,
			wantOutput:   "0",
		},
		{
			name:         "Adds money - both positive values",
			initialValue: 10,
			addValue:     30,
			wantOutput:   "40",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m1, err := money.NewMoney(tc.initialValue)
			testkit.AssertIsNil(t, err)
			m2, err := money.NewMoney(tc.addValue)
			testkit.AssertIsNil(t, err)

			m3, err := m1.Add(m2)

			testkit.AssertEqual(t, tc.wantError, err)
			testkit.AssertEqual(t, tc.wantOutput, m3.String())
		})
	}
}
