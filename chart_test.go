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

func TestApplyBigIntChan(t *testing.T) {
  xs := gochart.NewInts(3, 3, 5)
  ys := xs.ApplyBigIntChan(to30By2())
  assertBigValuesEqual(t, ys, 6, 12, 18, 24, 30)
}

func TestApplyBigIntChanSame(t *testing.T) {
  xs := gochart.NewInts(3, 0, 5)
  assertPanic(t, func() {
    xs.ApplyBigIntChan(to30By2())
  })
}

func TestApplyBigIntChanLess1(t *testing.T) {
  xs := gochart.NewInts(0, 3, 5)
  assertPanic(t, func() {
    xs.ApplyBigIntChan(to30By2())
  })
}

func TestApplyBigIntChanGreater(t *testing.T) {
  xs := gochart.NewInts(3, 3, 6)
  assertPanic(t, func() {
    xs.ApplyBigIntChan(to30By2())
  })
}

func TestApplyBigIntChanDown(t *testing.T) {
  xs := gochart.NewInts(15, -3, 5)
  assertPanic(t, func() {
    xs.ApplyBigIntChan(to30By2())
  })
}

func TestApplyChan(t *testing.T) {
  xs := gochart.NewInts(3, 3, 6)
  ys := xs.ApplyChan(to30By2Int())
  assertValuesEqual(
      t, ys, int64(6), int64(12), int64(18), int64(24), int64(30), int64(0))
}

func TestApplyChanSame(t *testing.T) {
  xs := gochart.NewInts(3, 0, 5)
  assertPanic(t, func() {
    xs.ApplyChan(to30By2Int())
  })
}

func TestApplyChanLess1(t *testing.T) {
  xs := gochart.NewInts(0, 3, 5)
  assertPanic(t, func() {
    xs.ApplyChan(to30By2Int())
  })
}

func TestApplyChanDown(t *testing.T) {
  xs := gochart.NewInts(15, -3, 5)
  assertPanic(t, func() {
    xs.ApplyChan(to30By2Int())
  })
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

func to30By2Int() <-chan int64 {
  result := make(chan int64)
  go func() {
    defer close(result)
    for i := int64(2); i <= 30; i +=2 {
      result <- i
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

func assertPanic(t *testing.T, f func()) {
  t.Helper()
  defer func() {
    recover()
  }()
  f()
  t.Error("Expected panic")
}
