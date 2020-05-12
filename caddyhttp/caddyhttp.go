// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package caddyhttp

import (
	// plug in the server
	_ "github.com/billhu1996/caddy/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/billhu1996/caddy/caddyhttp/basicauth"
	_ "github.com/billhu1996/caddy/caddyhttp/bind"
	_ "github.com/billhu1996/caddy/caddyhttp/browse"
	_ "github.com/billhu1996/caddy/caddyhttp/errors"
	_ "github.com/billhu1996/caddy/caddyhttp/expvar"
	_ "github.com/billhu1996/caddy/caddyhttp/extensions"
	_ "github.com/billhu1996/caddy/caddyhttp/fastcgi"
	_ "github.com/billhu1996/caddy/caddyhttp/gzip"
	_ "github.com/billhu1996/caddy/caddyhttp/header"
	_ "github.com/billhu1996/caddy/caddyhttp/index"
	_ "github.com/billhu1996/caddy/caddyhttp/internalsrv"
	_ "github.com/billhu1996/caddy/caddyhttp/limits"
	_ "github.com/billhu1996/caddy/caddyhttp/log"
	_ "github.com/billhu1996/caddy/caddyhttp/markdown"
	_ "github.com/billhu1996/caddy/caddyhttp/mime"
	_ "github.com/billhu1996/caddy/caddyhttp/pprof"
	_ "github.com/billhu1996/caddy/caddyhttp/proxy"
	_ "github.com/billhu1996/caddy/caddyhttp/push"
	_ "github.com/billhu1996/caddy/caddyhttp/redirect"
	_ "github.com/billhu1996/caddy/caddyhttp/requestid"
	_ "github.com/billhu1996/caddy/caddyhttp/rewrite"
	_ "github.com/billhu1996/caddy/caddyhttp/root"
	_ "github.com/billhu1996/caddy/caddyhttp/status"
	_ "github.com/billhu1996/caddy/caddyhttp/templates"
	_ "github.com/billhu1996/caddy/caddyhttp/timeouts"
	_ "github.com/billhu1996/caddy/caddyhttp/websocket"
	_ "github.com/billhu1996/caddy/onevent"
)
