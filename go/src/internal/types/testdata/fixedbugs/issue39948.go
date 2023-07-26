// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

type T[P any] interface{
	P // ERROR term cannot be a type parameter
}
