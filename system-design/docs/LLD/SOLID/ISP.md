# Interface Segregation Principle

The Interface Segregation Principle (ISP) says that clients should not be forced to depend on or implement methods they do not use. Prefer small, focused interfaces over one large, general-purpose interface.

## Example

A single interface may force every device to support printing, scanning, and faxing, even when a device does not provide all three capabilities:

```go
type MultifunctionDevice interface {
 Print(document string)
 Scan() string
 Fax(document string)
}
```

Split those capabilities into smaller interfaces instead:

```go
type Printer interface {
 Print(document string)
}

type Scanner interface {
 Scan() string
}

type FaxMachine interface {
 Fax(document string)
}
```

A basic printer can implement only `Printer`, while a multifunction device can implement all three. Each client can depend only on the capability it needs.

## Benefits

- Avoids unnecessary or empty method implementations.
- Reduces coupling between clients and implementations.
- Makes code easier to test, maintain, and extend.
- Lets types expose only the capabilities they actually support.
