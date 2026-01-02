# Go 1.26 Playground

## Test Artifacts

Go 1.26 introduces the `-artifacts` flag.

`t.ArtifactDir()` directories can be saved under a specified directory.

```bash
go test -artifacts . -v .
```

After execution, temporary files for each test remain under the `_artifacts/` directory.