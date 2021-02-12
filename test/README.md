# Local Integration Tests

To run the tests.

```bash
sudo -E go test . -v
```

### Adding a test

To add a test please try to have both `Happy` and `Sad` tests defined for all new SDK methods.

Example test:

##### mymethod_test.go

```go
 
// TestHappyMethod will test my new method
func TestHappyMethod(t *testing.T) {
    params := "my good input"
    _, err := Client.V1().Method(params)
	if err != nil {
		t.Errorf("expected success running method: %v", err)
		t.FailNow()
	}
}

// TestSadMethod will false positive test my new method
func TestSadMethod(t *testing.T) {
    params := "my bad input"
    _, err := Client.V1().Method(params)
    if err == nil {
        t.Errorf("expected failure running method: %v", err)
        t.FailNow()
    }
}

```
