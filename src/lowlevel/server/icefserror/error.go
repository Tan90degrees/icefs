/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 08:07:47
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-07 16:01:20
 * @FilePath: /icefs/src/lowlevel/server/icefserror/errorcode.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefserror

import (
	"log"
	"syscall"
)

const (
	ICEFS_EOK     = 0
	ICEFS_BUG_ERR = 9999
)

func IcefsStdErrno(err error) int32 {
	// if err != nil {
	// 	pc, _, _, _ := runtime.Caller(1)
	// 	log.Println(runtime.FuncForPC(pc).Name(), err)
	// }
	switch err {
	case nil:
		return ICEFS_EOK
	case syscall.E2BIG:
		return 0x7
	case syscall.EACCES:
		return 0xd
	case syscall.EADDRINUSE:
		return 0x62
	case syscall.EADDRNOTAVAIL:
		return 0x63
	case syscall.EADV:
		return 0x44
	case syscall.EAFNOSUPPORT:
		return 0x61
	case syscall.EAGAIN:
		return 0xb
	case syscall.EALREADY:
		return 0x72
	case syscall.EBADE:
		return 0x34
	case syscall.EBADF:
		return 0x9
	case syscall.EBADFD:
		return 0x4d
	case syscall.EBADMSG:
		return 0x4a
	case syscall.EBADR:
		return 0x35
	case syscall.EBADRQC:
		return 0x38
	case syscall.EBADSLT:
		return 0x39
	case syscall.EBFONT:
		return 0x3b
	case syscall.EBUSY:
		return 0x10
	case syscall.ECANCELED:
		return 0x7d
	case syscall.ECHILD:
		return 0xa
	case syscall.ECHRNG:
		return 0x2c
	case syscall.ECOMM:
		return 0x46
	case syscall.ECONNABORTED:
		return 0x67
	case syscall.ECONNREFUSED:
		return 0x6f
	case syscall.ECONNRESET:
		return 0x68
	case syscall.EDEADLOCK:
		return 0x23
	case syscall.EDESTADDRREQ:
		return 0x59
	case syscall.EDOM:
		return 0x21
	case syscall.EDOTDOT:
		return 0x49
	case syscall.EDQUOT:
		return 0x7a
	case syscall.EEXIST:
		return 0x11
	case syscall.EFAULT:
		return 0xe
	case syscall.EFBIG:
		return 0x1b
	case syscall.EHOSTDOWN:
		return 0x70
	case syscall.EHOSTUNREACH:
		return 0x71
	case syscall.EIDRM:
		return 0x2b
	case syscall.EILSEQ:
		return 0x54
	case syscall.EINPROGRESS:
		return 0x73
	case syscall.EINTR:
		return 0x4
	case syscall.EINVAL:
		return 0x16
	case syscall.EIO:
		return 0x5
	case syscall.EISCONN:
		return 0x6a
	case syscall.EISDIR:
		return 0x15
	case syscall.EISNAM:
		return 0x78
	case syscall.EKEYEXPIRED:
		return 0x7f
	case syscall.EKEYREJECTED:
		return 0x81
	case syscall.EKEYREVOKED:
		return 0x80
	case syscall.EL2HLT:
		return 0x33
	case syscall.EL2NSYNC:
		return 0x2d
	case syscall.EL3HLT:
		return 0x2e
	case syscall.EL3RST:
		return 0x2f
	case syscall.ELIBACC:
		return 0x4f
	case syscall.ELIBBAD:
		return 0x50
	case syscall.ELIBEXEC:
		return 0x53
	case syscall.ELIBMAX:
		return 0x52
	case syscall.ELIBSCN:
		return 0x51
	case syscall.ELNRNG:
		return 0x30
	case syscall.ELOOP:
		return 0x28
	case syscall.EMEDIUMTYPE:
		return 0x7c
	case syscall.EMFILE:
		return 0x18
	case syscall.EMLINK:
		return 0x1f
	case syscall.EMSGSIZE:
		return 0x5a
	case syscall.EMULTIHOP:
		return 0x48
	case syscall.ENAMETOOLONG:
		return 0x24
	case syscall.ENAVAIL:
		return 0x77
	case syscall.ENETDOWN:
		return 0x64
	case syscall.ENETRESET:
		return 0x66
	case syscall.ENETUNREACH:
		return 0x65
	case syscall.ENFILE:
		return 0x17
	case syscall.ENOANO:
		return 0x37
	case syscall.ENOBUFS:
		return 0x69
	case syscall.ENOCSI:
		return 0x32
	case syscall.ENODATA:
		return 0x3d
	case syscall.ENODEV:
		return 0x13
	case syscall.ENOENT:
		return 0x2
	case syscall.ENOEXEC:
		return 0x8
	case syscall.ENOKEY:
		return 0x7e
	case syscall.ENOLCK:
		return 0x25
	case syscall.ENOLINK:
		return 0x43
	case syscall.ENOMEDIUM:
		return 0x7b
	case syscall.ENOMEM:
		return 0xc
	case syscall.ENOMSG:
		return 0x2a
	case syscall.ENONET:
		return 0x40
	case syscall.ENOPKG:
		return 0x41
	case syscall.ENOPROTOOPT:
		return 0x5c
	case syscall.ENOSPC:
		return 0x1c
	case syscall.ENOSR:
		return 0x3f
	case syscall.ENOSTR:
		return 0x3c
	case syscall.ENOSYS:
		return 0x26
	case syscall.ENOTBLK:
		return 0xf
	case syscall.ENOTCONN:
		return 0x6b
	case syscall.ENOTDIR:
		return 0x14
	case syscall.ENOTEMPTY:
		return 0x27
	case syscall.ENOTNAM:
		return 0x76
	case syscall.ENOTRECOVERABLE:
		return 0x83
	case syscall.ENOTSOCK:
		return 0x58
	case syscall.ENOTTY:
		return 0x19
	case syscall.ENOTUNIQ:
		return 0x4c
	case syscall.ENXIO:
		return 0x6
	case syscall.EOPNOTSUPP:
		return 0x5f
	case syscall.EOVERFLOW:
		return 0x4b
	case syscall.EOWNERDEAD:
		return 0x82
	case syscall.EPERM:
		return 0x1
	case syscall.EPFNOSUPPORT:
		return 0x60
	case syscall.EPIPE:
		return 0x20
	case syscall.EPROTO:
		return 0x47
	case syscall.EPROTONOSUPPORT:
		return 0x5d
	case syscall.EPROTOTYPE:
		return 0x5b
	case syscall.ERANGE:
		return 0x22
	case syscall.EREMCHG:
		return 0x4e
	case syscall.EREMOTE:
		return 0x42
	case syscall.EREMOTEIO:
		return 0x79
	case syscall.ERESTART:
		return 0x55
	case syscall.ERFKILL:
		return 0x84
	case syscall.EROFS:
		return 0x1e
	case syscall.ESHUTDOWN:
		return 0x6c
	case syscall.ESOCKTNOSUPPORT:
		return 0x5e
	case syscall.ESPIPE:
		return 0x1d
	case syscall.ESRCH:
		return 0x3
	case syscall.ESRMNT:
		return 0x45
	case syscall.ESTALE:
		return 0x74
	case syscall.ESTRPIPE:
		return 0x56
	case syscall.ETIME:
		return 0x3e
	case syscall.ETIMEDOUT:
		return 0x6e
	case syscall.ETOOMANYREFS:
		return 0x6d
	case syscall.ETXTBSY:
		return 0x1a
	case syscall.EUCLEAN:
		return 0x75
	case syscall.EUNATCH:
		return 0x31
	case syscall.EUSERS:
		return 0x57
	case syscall.EXDEV:
		return 0x12
	case syscall.EXFULL:
		return 0x36
	default:
		return ICEFS_BUG_ERR
	}
}

func Must(v any, err error) any {
	if err != nil {
		log.Fatal(err)
	}
	return v
}
