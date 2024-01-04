# go-magic

[![codecov](https://codecov.io/gh/tosone/go-magic/graph/badge.svg?token=UZ4O8G0TTG)](https://codecov.io/gh/tosone/go-magic)

binding libmagic which is used to detect file type.

You should download the [magic.mgc](https://raw.githubusercontent.com/tosone/go-magic/main/magic/magic.mgc) file, because new magic instance required.

``` go
inst, err := New(MAGIC_MIME_TYPE, "./magic.mgc")
assert.NoError(t, err)
ret, err := inst.File("./2401.01663.pdf")
assert.NoError(t, err)
assert.Equal(t, "application/pdf", ret)
```
