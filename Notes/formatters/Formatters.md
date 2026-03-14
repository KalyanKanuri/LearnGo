# General:

- `%v` - default format for any value
- `%+v` - default format with field names for structs
- `%#v` - Go-syntax representation of the value
- `%T` - type of the value
- `%%` - literal percent sign

## Boolean:

- `%t` - true or false

## Integer:

- `%b` - binary
- `%c` - character (Unicode code point)
- `%d` - decimal
- `%o` - octal
- `%O` - octal with 0o prefix
- `%q` - single-quoted character literal
- `%x` - hexadecimal (lowercase a-f)
- `%X` - hexadecimal (uppercase A-F)
- `%U` - Unicode format (U+1234)

## Floating Point:

- `%e` - scientific notation (lowercase e)
- `%E` - scientific notation (uppercase E)
- `%f` - decimal point, no exponent
- `%F` - same as %f
- `%g` - %e or %f, whichever is more compact
- `%G` - %E or %F, whichever is more compact

## String and Byte Slice:

- `%s` - string or byte slice
- `%q` - double-quoted string
- `%x` - hexadecimal (lowercase), byte by byte
- `%X` - hexadecimal (uppercase), byte by byte

## Pointer:

- `%p` - pointer address in hexadecimal with 0x prefix

## Width and Precision:

- `%9d` - width 9, right-justified
- `%-9d` - width 9, left-justified
- `%.2f` - 2 decimal places
- `%9.2f` - width 9, 2 decimal places
- `%09d` - pad with zeros

## Flags:

- `+` - always print sign for numbers
- `-` - left-justify
- `#` - alternate format (0x for hex, 0o for octal)
- `(space)` - leave space for sign
- `0` - pad with zeros
