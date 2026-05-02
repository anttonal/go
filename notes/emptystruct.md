# Empty Struct

`struct{}` is a unary value — only one possible instance exists. Zero bytes in memory.

## Usage

```go
// anonymous
empty := struct{}{}

// named
type emptyStruct struct{}
empty := emptyStruct{}
```

## Why Use It

- Signals presence without carrying data
- Common as map value for sets: `map[string]struct{}`
- 0 bytes vs `bool`'s 1 byte per entry

## Memory Comparison

| Type | Size |
|------|------|
| struct{}{} | 0 bytes |
| bool | 1 byte |
| uint16 | 2 bytes |
| int64 | 8 bytes |

Mostly used with maps and channels.
