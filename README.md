# regexp

A feature-rich RegExp engine for Go+, compatible with Perl5 and .NET

## example

```go
import "github.com/goplus/regexp"

re := regexp`^[a-z]+\[[0-9]+\]$`.compile!
echo re.matchString("adam[23]")
```
