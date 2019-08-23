# Go Modules Examples

Experiments with Go modules using Go 1.12. This experiment
uses single module repositories (rather than more complex
[multi-module repositories](https://github.com/golang/go/wiki/Modules#what-are-multi-module-repositories)).

# Notes
The version in the go.mod file must match your semantic version git tag. For example, if go.mod has:

    module github.com/drichardson/hello/v3

Then you also need a git tag created like so:

    git tag v3.0.0
    git push origin v3.0.0

## References

- [Using Go Modules](https://blog.golang.org/using-go-modules)
- [go mod documentation](https://golang.org/cmd/go/#hdr-Module_maintenance)
- [Go 1.11 Modules | golang github Wiki](https://github.com/golang/go/wiki/Modules)
- [Go Modules in 2019](https://blog.golang.org/modules2019)
- [Semantic Import Versioning](https://research.swtch.com/vgo-import) - incompatible versions must have different import paths.
- [vgo-module](https://research.swtch.com/vgo-module)
