# WordUtils Go Library

This library provides two functions for word manipulation:

- **WordJumble**: Randomly shuffles the letters of a word (supports Unicode).
- **WordCypher**: Applies a Caesar cipher to a word, shifting only letters (A-Z, a-z).

## Usage

```go
import "github.com/yourusername/wordutils"

jumbled := wordutils.WordJumble("example")
cyphered := wordutils.WordCypher("Hello, World!", 5)
```

## Functions

### WordJumble

```go
func WordJumble(word string) string
```
Returns a new string with the letters of the input word randomly shuffled.

### WordCypher

```go
func WordCypher(word string, shift int) string
```
Returns a new string with each letter shifted by the given offset using a Caesar cipher. Only letters are shifted; other characters remain unchanged.

## Testing

Run tests with:

```
go test
```

## License
MIT

