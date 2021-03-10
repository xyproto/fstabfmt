# fstabfmt

Format `/etc/fstab` files.

## Features and limitations

* Can format `/etc/fstab` files.
* Will use 2 spaces between all fields, if they are of equal length.
* The shorter fields are padded with spaces.

## Example use

* Run `fstabfmt /etc/fstab` to see the formatted output. No changes are made to the file.
* Run `fstabfmt -i /etc/fstab` to make changes to `/etc/fstab`. Make sure to have a backup.

## General info

* Version: 1.0.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
