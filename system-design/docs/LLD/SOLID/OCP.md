# Open/Closed Principle

The **Open/Closed Principle (OCP)** states that software entities—such as classes, modules, and functions—should be **open for extension but closed for modification**.

This means you should be able to add or change behavior without rewriting stable, existing code. A common way to achieve this is to define an abstraction and provide new implementations of it.

## Example

Suppose a report can be exported in multiple formats. Instead of adding a new conditional branch to a single exporter every time a format is introduced, define an export interface and implement it for each format:

```go
type Exporter interface {
 Export(data []byte) error
}

type PDFExporter struct{}

func (PDFExporter) Export(data []byte) error {
 // Write data as a PDF.
 return nil
}

type CSVExporter struct{}

func (CSVExporter) Export(data []byte) error {
 // Write data as a CSV.
 return nil
}
```

The code that uses an `Exporter` can accept new formats through additional implementations, without needing to change its existing logic.

## Benefits

- Reduces the risk of introducing regressions in stable code.
- Makes new features easier to add and test independently.
- Encourages designs that depend on abstractions rather than concrete implementations.

OCP does not mean that existing code must never change. It is a guideline: identify the behavior likely to vary and design extension points where they provide a practical benefit.
