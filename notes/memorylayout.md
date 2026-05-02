# Memory Layout

Structs sit in memory as contiguous blocks. Fields placed sequentially.

## Field Ordering Matters

Good order (4 bytes total):
```go
type stats struct {
    Reach    uint16 // 2 bytes
    NumPosts uint8  // 1 byte
    NumLikes uint8  // 1 byte
}
```
No padding needed — fields naturally aligned.

Bad order (6 bytes total):
```go
type stats struct {
    NumPosts uint8  // 1 byte + 1 byte padding
    Reach    uint16 // 2 bytes
    NumLikes uint8  // 1 byte + 1 byte padding
}
```
Compiler inserts padding between misaligned fields for CPU access speed.

## Rule

Order fields largest to smallest to minimize padding.

## Debugging Layout

```go
typ := reflect.TypeOf(stats{})
fmt.Printf("Struct is %d bytes\n", typ.Size())
```

## When to Care

Rarely. Matters at scale — hundreds of thousands of structs in memory. Reordering fields saved 2+ GB in one production case.
