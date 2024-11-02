package generator

func Sum(a, b int) <-chan int {
  out := make(chan int)
  go func(){
    defer close(out)

    out <- a + b
  }()
  return out
}
