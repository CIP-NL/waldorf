package waldorf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getWaldorf() *Observation {
	w := Observe()
	w.RidiculedTwiceFunc = func(observation *Observation, ridiculedAlready bool) {

	}
	return w
}

func TestObservation_IsFalse(t *testing.T) {
	tcs := []struct {
		statement bool
		wantErr   bool
		desc      string
	}{
		{
			statement: 1 == 2,
			wantErr:   false,
			desc:      "ShouldBeFalse(1 == 2) should return True",
		},
		{
			statement: 1 == 1,
			wantErr:   true,
			desc:      "ShouldBeFalse(1 == 1) should return False",
		},
	}

	waldorf := getWaldorf()
	lenComps := 0

	for _, tc := range tcs {
		if tc.wantErr {
			lenComps += 1
			assert.False(t, waldorf.ShouldBeFalse(tc.statement, ""), tc.desc)
		} else {
			assert.True(t, waldorf.ShouldBeFalse(tc.statement, ""), tc.desc)
		}
	}
	assertLenIs(t, waldorf, lenComps)
}

func TestObservation_IsTrue(t *testing.T) {
	tcs := []struct {
		statement bool
		wantErr   bool
		desc      string
	}{
		{
			statement: 1 == 1,
			wantErr:   false,
			desc:      "ShouldBeTrue(1 == 1) should return %s",
		},
		{
			statement: 1 == 2,
			wantErr:   true,
			desc:      "ShouldBeTrue(1 == 2) should return %s",
		},
	}

	waldorf := getWaldorf()
	lenComps := 0

	for _, tc := range tcs {
		if tc.wantErr {
			lenComps += 1
			assert.False(t, waldorf.ShouldBeTrue(tc.statement, ""), tc.desc, !tc.wantErr)
		} else {
			assert.True(t, waldorf.ShouldBeTrue(tc.statement, ""), tc.desc, !tc.wantErr)
		}
	}
	assertLenIs(t, waldorf, lenComps)
}

func TestObservation_EitherShouldBeTrue(t *testing.T) {
	tcs := []struct {
		statementOne bool
		statementTwo bool
		wantErr      bool
		desc         string
	}{
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 1,
			wantErr:      false,
			desc:         "EitherShouldBeTrue(1 == 1, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 1,
			wantErr:      false,
			desc:         "EitherShouldBeTrue(1 == 2, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 2,
			wantErr:      false,
			desc:         "EitherShouldBeTrue(1 == 1, 1 == 2) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 2,
			wantErr:      true,
			desc:         "EitherShouldBeTrue(1 == 2, 1 == 2) should be %s",
		},
	}

	waldorf := getWaldorf()
	lenComps := 0

	for _, tc := range tcs {
		if tc.wantErr {
			lenComps += 1
			assert.False(t, waldorf.EitherShouldBeTrue(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		} else {
			assert.True(t, waldorf.EitherShouldBeTrue(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		}
	}
	assertLenIs(t, waldorf, lenComps)
}

func TestObservation_BothShouldBeTrue(t *testing.T) {
	tcs := []struct {
		statementOne bool
		statementTwo bool
		wantErr      bool
		desc         string
	}{
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 1,
			wantErr:      false,
			desc:         "BothShouldBeTrue(1 == 1, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 1,
			wantErr:      true,
			desc:         "BothShouldBeTrue(1 == 2, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 2,
			wantErr:      true,
			desc:         "BothShouldBeTrue(1 == 1, 1 == 2) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 2,
			wantErr:      true,
			desc:         "BothShouldBeTrue(1 == 2, 1 == 2) should be %s",
		},
	}

	waldorf := getWaldorf()
	lenComps := 0

	for _, tc := range tcs {
		if tc.wantErr {
			lenComps += 1
			assert.False(t, waldorf.BothShouldBeTrue(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		} else {
			assert.True(t, waldorf.BothShouldBeTrue(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		}
	}
	assertLenIs(t, waldorf, lenComps)
}

func TestObservation_EitherNeither(t *testing.T) {
	tcs := []struct {
		statementOne bool
		statementTwo bool
		wantErr      bool
		desc         string
	}{
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 1,
			wantErr:      true,
			desc:         "EitherNeither(1 == 1, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 1,
			wantErr:      false,
			desc:         "EitherNeither(1 == 2, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 2,
			wantErr:      false,
			desc:         "EitherNeither(1 == 1, 1 == 2) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 2,
			wantErr:      false,
			desc:         "EitherNeither(1 == 2, 1 == 2) should be %s",
		},
	}

	waldorf := getWaldorf()
	lenComps := 0

	for _, tc := range tcs {
		if tc.wantErr {
			lenComps += 1
			assert.False(t, waldorf.EitherNeither(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		} else {
			assert.True(t, waldorf.EitherNeither(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		}
	}
	assertLenIs(t, waldorf, lenComps)
}

func TestObservation_IfThisThenThat(t *testing.T) {
	tcs := []struct {
		statementOne bool
		statementTwo bool
		wantErr      bool
		desc         string
	}{
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 1,
			wantErr:      false,
			desc:         "IfThisThenThat(1 == 1, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 1,
			wantErr:      false,
			desc:         "IfThisThenThat(1 == 2, 1 == 1) should be %s",
		},
		{
			statementOne: 1 == 1,
			statementTwo: 1 == 2,
			wantErr:      true,
			desc:         "IfThisThenThat(1 == 1, 1 == 2) should be %s",
		},
		{
			statementOne: 1 == 2,
			statementTwo: 1 == 2,
			wantErr:      false,
			desc:         "IfThisThenThat(1 == 2, 1 == 2) should be %s",
		},
	}

	waldorf := getWaldorf()
	lenComps := 0

	for _, tc := range tcs {
		if tc.wantErr {
			lenComps += 1
			assert.False(t, waldorf.IfThisThenThat(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		} else {
			assert.True(t, waldorf.IfThisThenThat(tc.statementOne, tc.statementTwo, ""), tc.desc, !tc.wantErr)
		}
	}
	assertLenIs(t, waldorf, lenComps)
}

func assertLenIs(t *testing.T, observation *Observation, i int) {
	comp := observation.Ridicule()
	if i == 0 {
		assert.Nil(t, comp, "waldorf should not have any complaints")
		return
	}
	if assert.NotNil(t, comp) {
		assert.True(t, len(comp.comps) == i, "Waldor has an incorrect amount of complaints")
	}
}

func TestObservation_Ridicule(t *testing.T) {
	waldorf := Observe()

	const (
		notRidiculed = "Are numbers equal?"
		incorrectOne = "are Jesus slippers Cool?"
		incorrectTwo = "Switserland is next to Sweden right?"
	)

	// correct observation, should not be in the list
	waldorf.ShouldBeTrue(1 == 1, notRidiculed)

	// incorrect observations, should be returned
	waldorf.ShouldBeTrue(1 == 2, incorrectOne)
	waldorf.ShouldBeFalse("A" == "A", incorrectTwo)

	everythingWrong := waldorf.Ridicule()
	assert.NotContains(t, everythingWrong.comps, notRidiculed)
	assert.Contains(t, everythingWrong.comps, incorrectOne)
	assert.Contains(t, everythingWrong.comps, incorrectTwo)
}
