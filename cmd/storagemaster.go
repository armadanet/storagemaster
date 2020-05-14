package main

import(
  "github.com/armadanet/storagemaster"
)

func main() {
  s := storagemaster.New()
  s.Run(8086)
}
