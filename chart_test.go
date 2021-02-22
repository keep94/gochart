package gochart_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/keep94/gochart"
	"github.com/keep94/gomath"
)

func TestApply(t *testing.T) {
	xs := gochart.NewInts(5, 2, 3)
	assertValuesEqual(t, xs, int64(5), int64(7), int64(9))
	ys := xs.Apply(func(x int64) int64 { return 3 * x })
	assertValuesEqual(t, ys, int64(15), int64(21), int64(27))
}

func TestApplyBigInt(t *testing.T) {
	xs := gochart.NewInts(10, 1, 4)
	ys := xs.ApplyBigInt(
		func(x int64, result *big.Int) *big.Int {
			bx := big.NewInt(x)
			return result.Mul(bx, bx)
		})
	assertBigValuesEqual(t, ys, 100, 121, 144, 169)
}

func TestApplyBigIntStream(t *testing.T) {
	xs := gochart.NewInts(3, 3, 5)
	ys := xs.ApplyBigIntStream(upBy2())
	assertBigValuesEqual(t, ys, 6, 12, 18, 24, 30)
}

func TestApplyBigIntStreamSame(t *testing.T) {
	xs := gochart.NewInts(3, 0, 5)
	assertPanic(t, func() {
		xs.ApplyBigIntStream(upBy2())
	})
}

func TestApplyBigIntStreamLess1(t *testing.T) {
	xs := gochart.NewInts(0, 3, 5)
	assertPanic(t, func() {
		xs.ApplyBigIntStream(upBy2())
	})
}

func TestApplyBigIntStreamDown(t *testing.T) {
	xs := gochart.NewInts(15, -3, 5)
	assertPanic(t, func() {
		xs.ApplyBigIntStream(upBy2())
	})
}

func TestApplySlice(t *testing.T) {
	xs := gochart.NewInts(2, 2, 5)
	ys := xs.ApplySlice([]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29})
	assertValuesEqual(
		t, ys, int64(3), int64(7), int64(13), int64(19), int64(29))
}

func TestApplySlicePanicNegative(t *testing.T) {
	xs := gochart.NewInts(0, 1, 3)
	assertPanic(t, func() {
		xs.ApplySlice([]int64{1, 2, 3})
	})
}

func TestApplySlicePanicPastLength(t *testing.T) {
	xs := gochart.NewInts(1, 1, 3)
	assertPanic(t, func() {
		xs.ApplySlice([]int64{1, 2})
	})
}

func TestApplyStream(t *testing.T) {
	xs := gochart.NewInts(3, 3, 6)
	ys := xs.ApplyStream(to30By2())
	assertValuesEqual(
		t, ys, int64(6), int64(12), int64(18), int64(24), int64(30), int64(0))
}

func TestApplyStreamSame(t *testing.T) {
	xs := gochart.NewInts(3, 0, 5)
	assertPanic(t, func() {
		xs.ApplyStream(to30By2())
	})
}

func TestApplyStreamLess1(t *testing.T) {
	xs := gochart.NewInts(0, 3, 5)
	assertPanic(t, func() {
		xs.ApplyStream(to30By2())
	})
}

func TestApplyStreamDown(t *testing.T) {
	xs := gochart.NewInts(15, -3, 5)
	assertPanic(t, func() {
		xs.ApplyStream(to30By2())
	})
}

func TestApplyFloat(t *testing.T) {
	xs := gochart.NewFloats(1.0, 2.0, 4)
	ys := xs.Apply(func(x float64) float64 {
		return x * x
	})
	assertValuesEqual(t, xs, 1.0, 3.0, 5.0, 7.0)
	assertValuesEqual(t, ys, 1.0, 9.0, 25.0, 49.0)
}

func TestApplyInvFloat(t *testing.T) {
	xs := gochart.NewFloats(0.0, 1.0, 6)
	ys := xs.ApplyInv(
		func(x float64) float64 {
			return x * x
		},
		1,
		2)
	if ys.Len() != 6 {
		t.Fatal("Expected length of 4")
	}
	if ys.Value(0).(float64) != 1.0 {
		t.Error("Expected 0 to give value of 1")
	}
	if ys.Value(1).(float64) != 1.0 {
		t.Error("Expected 1 to give value of 1")
	}
	assertCloseTo(t, 1.4142, ys.Value(2).(float64))
	assertCloseTo(t, 1.7321, ys.Value(3).(float64))
	if ys.Value(4).(float64) != 2.0 {
		t.Error("Expected 4 to give value of 2")
	}
	if ys.Value(5).(float64) != 2.0 {
		t.Error("Expected 5 to give value of 2")
	}
}

func TestApplyInvFloatReverse(t *testing.T) {
	xs := gochart.NewFloats(0.0, 1.0, 4)
	ys := xs.ApplyInv(
		func(x float64) float64 {
			return 4.0 - x*x
		},
		0,
		2)
	if ys.Len() != 4 {
		t.Fatal("Expected length of 4")
	}
	if ys.Value(0).(float64) != 2.0 {
		t.Error("Expected value of 2")
	}
	assertCloseTo(t, 1.7321, ys.Value(1).(float64))
	assertCloseTo(t, 1.4142, ys.Value(2).(float64))
	if ys.Value(3).(float64) != 1.0 {
		t.Error("Expected value of 1")
	}
}

func TestChartDimensions(t *testing.T) {
	xs := gochart.NewInts(1, 1, 100)
	chart := gochart.NewChart(xs, xs)
	assertEqual(t, 100, chart.NumRows())
	assertEqual(t, 1, chart.NumCols())
	chart = gochart.NewChart(xs, xs, gochart.NumRows(33))
	assertEqual(t, 33, chart.NumRows())
	assertEqual(t, 4, chart.NumCols())
	chart = gochart.NewChart(xs, xs, gochart.NumCols(3))
	assertEqual(t, 34, chart.NumRows())
	assertEqual(t, 3, chart.NumCols())
	chart = gochart.NewChart(xs, xs, gochart.NumCols(5), gochart.NumRows(25))
	assertEqual(t, 25, chart.NumRows())
	assertEqual(t, 5, chart.NumCols())
}

func TestOptions(t *testing.T) {
	xs := gochart.NewInts(1, 1, 100)
	options := gochart.Options{gochart.NumCols(6), gochart.NumRows(26)}
	chart := gochart.NewChart(xs, xs, options)
	assertEqual(t, 26, chart.NumRows())
	assertEqual(t, 6, chart.NumCols())
}

func TestNewChartPanic(t *testing.T) {
	xs := gochart.NewInts(1, 1, 10)
	ys := gochart.NewInts(1, 1, 9)
	assertPanic(
		t,
		func() {
			gochart.NewChart(xs, ys)
		},
	)
}

func TestFloatsPanic(t *testing.T) {
	xs := gochart.NewFloats(1.0, 1.0, 10)
	assertPanic(t, func() { xs.Value(10) })
	assertPanic(t, func() { xs.Value(-1) })
}

func TestIntsPanic(t *testing.T) {
	xs := gochart.NewInts(1, 1, 10)
	assertPanic(t, func() { xs.Value(10) })
	assertPanic(t, func() { xs.Value(-1) })
}

func TestValuesPanic(t *testing.T) {
	xs := gochart.NewFloats(1.0, 1.0, 10)
	ys := xs.Apply(math.Sqrt)
	assertPanic(t, func() { ys.Value(10) })
	assertPanic(t, func() { ys.Value(-1) })
}

func assertValuesEqual(
	t *testing.T, ys gochart.Values, expectedValues ...interface{}) {
	t.Helper()
	if ys.Len() != len(expectedValues) {
		t.Fatalf("Expected %v values, but got %v", len(expectedValues), ys.Len())
	}
	for i := 0; i < ys.Len(); i++ {
		if ys.Value(i) != expectedValues[i] {
			t.Errorf("Expected %v, got %v", expectedValues[i], ys.Value(i))
		}
	}
}

func assertEqual(
	t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertBigValuesEqual(
	t *testing.T, ys gochart.Values, expectedValues ...int64) {
	t.Helper()
	if ys.Len() != len(expectedValues) {
		t.Fatalf("Expected %v values, but got %v", len(expectedValues), ys.Len())
	}
	for i := 0; i < ys.Len(); i++ {
		if ys.Value(i).(*big.Int).Cmp(big.NewInt(expectedValues[i])) != 0 {
			t.Errorf("Expected %v, got %v", expectedValues[i], ys.Value(i))
		}
	}
}

type linearBigIntStream struct {
	start *big.Int
	incr  *big.Int
}

func (s *linearBigIntStream) Next(value *big.Int) *big.Int {
	value.Set(s.start)
	s.start.Add(s.start, s.incr)
	return value
}

func upBy2() gomath.BigIntStream {
	return &linearBigIntStream{start: big.NewInt(2), incr: big.NewInt(2)}
}

type linearIntStream struct {
	start int64
	incr  int64
	max   int64
}

func (s *linearIntStream) Next() (result int64, ok bool) {
	if s.start > s.max {
		return
	}
	result = s.start
	ok = true
	s.start += s.incr
	return
}

func to30By2() gomath.IntStream {
	return &linearIntStream{start: 2, incr: 2, max: 30}
}

func assertCloseTo(t *testing.T, expected float64, actual float64) {
	t.Helper()
	if math.Abs((expected-actual)/expected) > 0.0001 {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		recover()
	}()
	f()
	t.Error("Expected panic")
}
