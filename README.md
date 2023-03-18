Go Memory Cache
================================

Go Memory Cache helps you store and manage data in memory.

See it in action:

## Example #1

```go
package main

import "github.com/Anatolii1108/golang-ninja-cache"


func main() {
	cache := cache.NewMemoryCache()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
```
