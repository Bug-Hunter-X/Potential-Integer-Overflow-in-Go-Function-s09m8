func myFunc(a, b int) (int, error) {
  if a == 0 {
    return b, nil
  }
  if (b > 0 && a > math.MaxInt32-b) || (b < 0 && a < math.MinInt32-b) {
    return 0, fmt.Errorf("integer overflow")
  }
  return a + b, nil
}
import "math"
import "fmt"