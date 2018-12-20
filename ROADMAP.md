# Roadmap

long-term
1. Waldorf is currently used for API validation. Once switched to using autogenerated validation, add a waldorf.Validator interface 
and a method to use the generated validation logic; i.e.

```go
package waldorf 

type Validator interface {
    Validate() error
}

func (w *Waldorf) Validate(msg Validator) {
...
}
```