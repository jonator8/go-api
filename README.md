# go-api

template for tests

```
func Test$NAME$(t *testing.T) {
    testCases := []struct {
        desc string
        
    }{
        {
            desc: "",
            
        },
    }
    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            $END$
        })
    }
}
```