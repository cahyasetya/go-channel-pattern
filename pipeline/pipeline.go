package pipeline

func Gen(nums ...int) <-chan int {
  out := make(chan int)
  go func() {
    for _, num := range nums {
      out <- num
    }
    close(out)
  }()
  return out
}

func Sq(in <-chan int) <-chan int {
  out := make(chan int)
  go func ()  {
    for num := range in {
      out <- num*num
    }
  }()
  return out
}
