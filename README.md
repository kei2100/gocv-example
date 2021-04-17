## Getting Started

```bash
$ go get -u -d gocv.io/x/gocv
```

```bash
$ brew install cmake
$ brew install opencv
$ brew install pkgconfig
```

```bash
$ brew info cmake opencv pkgconfig
Warning: Treating cmake as a formula. For the cask, use homebrew/cask/cmake
cmake: stable 3.20.1 (bottled), HEAD
Cross-platform make
https://www.cmake.org/
/usr/local/Cellar/cmake/3.20.1 (6,468 files, 69.6MB) *
  Poured from bottle on 2021-04-17 at 08:35:27
From: https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/cmake.rb
License: BSD-3-Clause
==> Dependencies
Build: sphinx-doc ✘
==> Options
--HEAD
        Install HEAD version
==> Caveats
Emacs Lisp files have been installed to:
  /usr/local/share/emacs/site-lisp/cmake
==> Analytics
install: 163,915 (30 days), 460,965 (90 days), 1,578,893 (365 days)
install-on-request: 114,775 (30 days), 323,053 (90 days), 1,063,109 (365 days)
build-error: 0 (30 days)

opencv: stable 4.5.2 (bottled)
Open source computer vision library
https://opencv.org/
/usr/local/Cellar/opencv/4.5.2 (838 files, 233.6MB) *
  Poured from bottle on 2021-04-17 at 08:27:00
From: https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/opencv.rb
License: Apache-2.0
==> Dependencies
Build: cmake ✔, pkg-config ✔
Required: ceres-solver ✔, eigen ✔, ffmpeg ✔, glog ✔, harfbuzz ✔, jpeg ✔, libpng ✔, libtiff ✔, numpy ✔, openblas ✔, openexr ✔, protobuf ✔, python@3.9 ✔, tbb ✔, vtk ✔, webp ✔
==> Analytics
install: 16,114 (30 days), 45,725 (90 days), 228,293 (365 days)
install-on-request: 15,467 (30 days), 44,053 (90 days), 212,948 (365 days)
build-error: 0 (30 days)

pkg-config: stable 0.29.2 (bottled)
Manage compile and link flags for libraries
https://freedesktop.org/wiki/Software/pkg-config/
/usr/local/Cellar/pkg-config/0.29.2_3 (11 files, 624.0KB) *
  Poured from bottle on 2021-04-17 at 08:22:25
From: https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/pkg-config.rb
License: GPL-2.0-or-later
==> Analytics
install: 205,549 (30 days), 661,128 (90 days), 2,786,916 (365 days)
install-on-request: 28,537 (30 days), 90,602 (90 days), 530,305 (365 days)
build-error: 0 (30 days)
```

```bash
$ go run ~/go/pkg/mod/gocv.io/x/gocv@v0.27.0/cmd/version/main.go 
gocv version: 0.27.0
opencv lib version: 4.5.2
```

## Writing Code

* example/hello-video/
* example/face-detect/
* example/qr-detect/
* more examples... https://gocv.io/writing-code/more-examples/
