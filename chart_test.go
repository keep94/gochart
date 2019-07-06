package gochart_test

import (
  "math"
  "math/big"
  "testing"

  "github.com/keep94/gochart"
)

func TestApply(t *testing.T) {
  xs := gochart.NewInts(5, 2, 3)
  assertValuesEqual(t, xs, int64(5), int64(7), int64(9))
  ys := xs.Apply(func(x int64) int64 { return 3*x })
  assertValuesEqual(t, ys, int64(15), int64(21), int64(27))
}

func TestApplyBigInt(t *testing.T) {
  xs := gochart.NewInts(10, 1, 4)
  ys := xs.ApplyBigInt(
      func(x int64) *big.Int {
        bx := big.NewInt(x)
        return new(big.Int).Mul(bx, bx)
      })
  assertBigValuesEqual(t, ys, 100, 121, 144, 169)
}

func TestApplyBigIntCh(t *testing.T) {
  xs := gochart.NewInts(-1, 3, 10)
  ys := xs.ApplyBigIntCh(to30By2())
  assertBigValuesEqual(t, ys, 0, 4, 10, 16, 22, 28, 0, 0, 0, 0)
}

func TestApplyFloat(t *testing.T) {
  xs := gochart.NewFloats(1.0, 2.0, 4)
  ys := xs.Apply(func(x float64) float64 {
          return x*x
        })
  assertValuesEqual(t, xs, 1.0, 3.0, 5.0, 7.0)
  assertValuesEqual(t, ys, 1.0, 9.0, 25.0, 49.0)
}

func TestApplyInvFloat(t *testing.T) {
  xs := gochart.NewFloats(0.0, 1.0, 6)
  ys := xs.ApplyInv(
      func(x float64) float64 {
        return x*x
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
  assertEqual(t, 100, chart.RowCount())
  assertEqual(t, 1, chart.ColCount())
  chart = gochart.NewChart(xs, xs, gochart.RowCount(33))
  assertEqual(t, 33, chart.RowCount())
  assertEqual(t, 4, chart.ColCount())
  chart = gochart.NewChart(xs, xs, gochart.ColCount(3))
  assertEqual(t, 34, chart.RowCount())
  assertEqual(t, 3, chart.ColCount())
  chart = gochart.NewChart(xs, xs, gochart.ColCount(5), gochart.RowCount(25))
  assertEqual(t, 25, chart.RowCount())
  assertEqual(t, 5, chart.ColCount())
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

func to30By2() <-chan *big.Int {
  result := make(chan *big.Int)
  go func() {
    defer close(result)
    for i := 2; i <= 30; i += 2 {
      result <- big.NewInt(int64(i))
    }
  }()
  return result
}

func assertCloseTo(t *testing.T, expected float64, actual float64) {
  t.Helper()
  if math.Abs((expected - actual) / expected) > 0.0001 {
    t.Errorf("Expected %v, got %v", expected, actual)
  }
}
