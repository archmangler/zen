# zen

* Just a dummy golang module for troubleshooting.

# Testing

* Write the test:

```
package zen_test

import (
        "testing"
        "github.com/archmangler/zen"
)

func TestKoan(t *testing.T) {
        if zen.Koan() != "Seijo's soul separated from her being. Which was the real Seijo?" {
                t.Fatal("Wrong Realization. No Satori for you!")
        }
}
```

E.g:

```
traiano@Traianos-iMac zen % go test -run ^TestKoan                     
--- FAIL: TestKoan (0.00s)
    zen_test.go:11: Wrong Realization. No Satori for you!
FAIL
exit status 1
FAIL    github.com/archmangler/zen      0.327s
traiano@Traianos-iMac zen % 
```

* modify the module output:

```
traiano@Traianos-iMac zen % go test -run ^TestKoan
PASS
ok      github.com/archmangler/zen      0.100s
traiano@Traianos-iMac zen % 
```
