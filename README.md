# fstabfmt

<img src="img/fstabfmt.svg" alt="logo" width="320px" height="320px">

![Build](https://github.com/xyproto/fstabfmt/workflows/Build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/fstabfmt)](https://goreportcard.com/report/github.com/xyproto/fstabfmt) [![License](https://img.shields.io/badge/license-BSD-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/fstabfmt/main/LICENSE)

Format `/etc/fstab` files.

## Features and limitations

* Can format `/etc/fstab` files.
* Will use 2 spaces between all fields, if they are of equal length.
* The shorter fields are padded with spaces.

## Example use

* Run `fstabfmt /etc/fstab` to see the formatted output. No changes are made to the file.
* Run `fstabfmt -i /etc/fstab` to make changes to `/etc/fstab`. Always keep a backup, just in case.

## Example output

Before:

```sh
# Static information about the filesystems.
# See fstab(5) for details.
#
# <file system> <dir> <type> <options> <dump> <pass>
# /dev/nvme0n1p2 LABEL=root
UUID=2bb3c21b-dc8f-401e-991b-66afd7301cb7	/         	xfs       	rw,relatime,inode64,logbufs=8,logbsize=32k,noquota	0 1

# /dev/nvme0n1p1 LABEL=boot
UUID=1815-DD5D      	/boot     	vfat      	rw,relatime,fmask=0022,dmask=0022,codepage=437,iocharset=iso8859-1,shortname=mixed,utf8,errors=remount-ro	0 2
```

After:

```sh
# Static information about the filesystems.
# See fstab(5) for details.
#
# <file system> <dir> <type> <options> <dump> <pass>
# /dev/nvme0n1p2 LABEL=root
UUID=2bb3c21b-dc8f-401e-991b-66afd7301cb7  /      xfs   rw,relatime,inode64,logbufs=8,logbsize=32k,noquota                                                         0  1
# /dev/nvme0n1p1 LABEL=boot
UUID=1815-DD5D                             /boot  vfat  rw,relatime,fmask=0022,dmask=0022,codepage=437,iocharset=iso8859-1,shortname=mixed,utf8,errors=remount-ro  0  2
```

## General info

* Version: 1.0.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
