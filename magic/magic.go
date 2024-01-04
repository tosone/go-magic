package magic

/*
 #cgo darwin CFLAGS: -DHAVE_STRLCPY -DHAVE_STRLCAT -DHAVE_MKSTEMP -DVERSION="1.3"
 #cgo linux CFLAGS: -DHAVE_MKSTEMP -DVERSION="1.3"
 #include <magic.h>
 #include <stdlib.h>
 #include <unistd.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	MAGIC_NONE              = C.MAGIC_NONE
	MAGIC_DEBUG             = C.MAGIC_DEBUG
	MAGIC_SYMLINK           = C.MAGIC_SYMLINK
	MAGIC_COMPRESS          = C.MAGIC_COMPRESS
	MAGIC_DEVICES           = C.MAGIC_DEVICES
	MAGIC_MIME_TYPE         = C.MAGIC_MIME_TYPE
	MAGIC_CONTINUE          = C.MAGIC_CONTINUE
	MAGIC_CHECK             = C.MAGIC_CHECK
	MAGIC_PRESERVE_ATIME    = C.MAGIC_PRESERVE_ATIME
	MAGIC_RAW               = C.MAGIC_RAW
	MAGIC_ERROR             = C.MAGIC_ERROR
	MAGIC_MIME_ENCODING     = C.MAGIC_MIME_ENCODING
	MAGIC_MIME              = C.MAGIC_MIME
	MAGIC_APPLE             = C.MAGIC_APPLE
	MAGIC_NO_CHECK_COMPRESS = C.MAGIC_NO_CHECK_COMPRESS
	MAGIC_NO_CHECK_TAR      = C.MAGIC_NO_CHECK_TAR
	MAGIC_NO_CHECK_SOFT     = C.MAGIC_NO_CHECK_SOFT
	MAGIC_NO_CHECK_APPTYPE  = C.MAGIC_NO_CHECK_APPTYPE
	MAGIC_NO_CHECK_ELF      = C.MAGIC_NO_CHECK_ELF
	MAGIC_NO_CHECK_TEXT     = C.MAGIC_NO_CHECK_TEXT
	MAGIC_NO_CHECK_CDF      = C.MAGIC_NO_CHECK_CDF
	MAGIC_NO_CHECK_TOKENS   = C.MAGIC_NO_CHECK_TOKENS
	MAGIC_NO_CHECK_ENCODING = C.MAGIC_NO_CHECK_ENCODING
	MAGIC_NO_CHECK_ASCII    = C.MAGIC_NO_CHECK_ASCII
	MAGIC_NO_CHECK_FORTRAN  = C.MAGIC_NO_CHECK_FORTRAN
	MAGIC_NO_CHECK_TROFF    = C.MAGIC_NO_CHECK_TROFF
)
const (
	MAGIC_NO_CHECK_BUILTIN = MAGIC_NO_CHECK_COMPRESS |
		MAGIC_NO_CHECK_TAR |
		MAGIC_NO_CHECK_APPTYPE |
		MAGIC_NO_CHECK_ELF |
		MAGIC_NO_CHECK_TEXT |
		MAGIC_NO_CHECK_CDF |
		MAGIC_NO_CHECK_TOKENS |
		MAGIC_NO_CHECK_ENCODING
)

// Magic ...
type Magic C.magic_t

// Instance ...
type Instance struct {
	Magic C.magic_t
}

// New ...
func New(flags int, magicFile string) (*Instance, error) {
	inst := (Magic)(C.magic_open(C.int(flags)))
	if magicFile == "" {
		ret := (int)(C.magic_load((C.magic_t)(inst), nil))
		if ret != 0 {
			return &Instance{Magic: (C.magic_t)(inst)}, fmt.Errorf("load magic file failed: %s", C.GoString((C.magic_error((C.magic_t)(inst)))))
		}
	}
	filename := C.CString(magicFile)
	defer C.free(unsafe.Pointer(filename))
	ret := (int)(C.magic_load((C.magic_t)(inst), filename))
	if ret != 0 {
		return &Instance{Magic: (C.magic_t)(inst)}, fmt.Errorf("load magic file failed: %s", C.GoString((C.magic_error((C.magic_t)(inst)))))
	}
	return &Instance{Magic: (C.magic_t)(inst)}, nil
}

// Close ...
func (i *Instance) Close() {
	C.magic_close((C.magic_t)(i.Magic))
}

// File ...
func (i *Instance) File(file string) (string, error) {
	filename := C.CString(file)
	defer C.free(unsafe.Pointer(filename))
	ret := C.magic_file((C.magic_t)(i.Magic), filename)
	if ret == nil {
		return "", fmt.Errorf("detect file type failed: %v", C.GoString((C.magic_error((C.magic_t)(i.Magic)))))
	}
	return C.GoString(ret), nil
}

// Buffer ...
func (i *Instance) Buffer(data []byte) (string, error) {
	ret := C.magic_buffer((C.magic_t)(i.Magic), unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if ret == nil {
		return "", fmt.Errorf("detect file type failed: %v", C.GoString((C.magic_error((C.magic_t)(i.Magic)))))
	}
	return C.GoString(ret), nil
}
