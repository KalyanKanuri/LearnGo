# Liskov Substitution Principle

The Liskov Substitution Principle (LSP) states that an implementation of a type should be usable wherever its parent type or interface is expected, without changing the correctness of the program. Implementations must honor the behavior promised by that contract.

## Example

In Go, interfaces make it easy to define interchangeable behavior. Both email and SMS notifiers can be passed to code that expects a `Notifier`:

```go
type Notifier interface {
 Send(message string) error
}

func NotifyUser(n Notifier, message string) error {
 return n.Send(message)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) error {
 // Send the message by email.
 return nil
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) error {
 // Send the message by SMS.
 return nil
}
```

`NotifyUser` works with either implementation without needing to know its concrete type. An implementation that silently discards valid messages or unexpectedly panics would break the expected contract and violate LSP.

## Benefits

- Lets callers use different implementations interchangeably.
- Reduces coupling to concrete types and makes systems easier to extend.
- Makes testing easier by allowing compatible fakes or mocks.
- Prevents unexpected behavior when introducing a new implementation.
