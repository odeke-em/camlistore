/*
Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"net/http"
	"syscall"

	"camlistore.org/pkg/httputil"
	"camlistore.org/pkg/netutil"
)

func setupHome(rw http.ResponseWriter, req *http.Request) {
	port := httputil.RequestTargetPort(req)
	ourAddr := fmt.Sprintf("127.0.0.1:%d", port)
	// TODO(mpl): fix IPv4 assumption
	uid, err := netutil.AddrPairUserid(req.RemoteAddr, ourAddr)

	fmt.Fprintf(rw, "Hello %q\n", req.RemoteAddr)
	fmt.Fprintf(rw, "<p>uid = %d\n", syscall.Getuid())
	fmt.Fprintf(rw, "<p>euid = %d\n", syscall.Geteuid())

	fmt.Fprintf(rw, "<p>http_local_uid(%q => %q) = %d (%v)\n", req.RemoteAddr, ourAddr, uid, err)
}
