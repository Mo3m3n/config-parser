/*
Copyright 2019 HAProxy Technologies

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

package parser

import (
	"github.com/haproxytech/config-parser/v2/parsers"
	"github.com/haproxytech/config-parser/v2/parsers/extra"
	"github.com/haproxytech/config-parser/v2/parsers/filters"
	"github.com/haproxytech/config-parser/v2/parsers/http"
	"github.com/haproxytech/config-parser/v2/parsers/simple"
	"github.com/haproxytech/config-parser/v2/parsers/stats"
	"github.com/haproxytech/config-parser/v2/parsers/tcp"
)

func createParsers(parser []ParserInterface) *Parsers {
	p := Parsers{
		Parsers: append(parser, []ParserInterface{
			&extra.Section{Name: "defaults"},
			&extra.Section{Name: "global"},
			&extra.Section{Name: "frontend"},
			&extra.Section{Name: "backend"},
			&extra.Section{Name: "listen"},
			&extra.Section{Name: "resolvers"},
			&extra.Section{Name: "userlist"},
			&extra.Section{Name: "peers"},
			&extra.Section{Name: "mailers"},
			&extra.Section{Name: "cache"},
			&extra.Section{Name: "program"},
			&extra.UnProcessed{},
		}...),
	}
	for _, parser := range p.Parsers {
		parser.Init()
	}
	return &p
}

func getStartParser() *Parsers {
	return createParsers([]ParserInterface{
		&extra.ConfigVersion{},
		&extra.Comments{},
	})
}

func getDefaultParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Mode{},
		&parsers.HashType{},
		&parsers.Balance{},

		&parsers.MaxConn{},
		&parsers.Log{},
		&parsers.OptionHTTPLog{},
		&stats.Stats{Mode: "defaults"},

		&simple.Word{Name: "log-tag"},

		&simple.String{Name: "log-format"},
		&simple.String{Name: "log-format-sd"},
		&parsers.Cookie{},
		&parsers.HTTPCheck{},
		&parsers.BindProcess{},

		&simple.Option{Name: "tcplog"},
		&simple.Option{Name: "httpclose"},
		&simple.Option{Name: "http-use-htx"},
		&parsers.OptionRedispatch{},
		&simple.Option{Name: "dontlognull"},
		&simple.Option{Name: "log-separate-errors"},
		&simple.Option{Name: "http-buffer-request"},
		&simple.Option{Name: "http-server-close"},
		&simple.Option{Name: "http-keep-alive"},
		&simple.Option{Name: "http-pretend-keepalive"},
		&simple.Option{Name: "clitcpka"},
		&simple.Option{Name: "contstats"},
		&simple.Option{Name: "ssl-hello-chk"},
		&parsers.OptionSmtpchk{},
		&simple.Option{Name: "ldap-check"},
		&parsers.OptionMysqlCheck{},
		&simple.Option{Name: "abortonclose"},
		&simple.Option{Name: "pgsql-check"},
		&simple.Option{Name: "tcp-check"},
		&simple.Option{Name: "redis-check"},
		&parsers.OptionHttpchk{},

		&simple.Option{Name: "logasap"},
		&simple.Option{Name: "allbackups"},
		&simple.Option{Name: "external-check"},
		&parsers.OptionForwardFor{},

		&parsers.HTTPReuse{},
		&simple.Timeout{Name: "http-request"},
		&simple.Timeout{Name: "check"},
		&simple.Timeout{Name: "connect"},
		&simple.Timeout{Name: "client"},
		&simple.Timeout{Name: "client-fin"},
		&simple.Timeout{Name: "queue"},
		&simple.Timeout{Name: "server"},
		&simple.Timeout{Name: "server-fin"},
		&simple.Timeout{Name: "tunnel"},
		&simple.Timeout{Name: "http-keep-alive"},

		&simple.Number{Name: "retries"},

		&parsers.ExternalCheckPath{},
		&parsers.ExternalCheckCommand{},
		&parsers.DefaultServer{},
		&parsers.ErrorFile{},
		&parsers.DefaultBackend{},
		&parsers.UniqueIDFormat{},
		&parsers.UniqueIDHeader{},
		&parsers.ConfigSnippet{},
	})
}

func getGlobalParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Daemon{},
		&simple.Word{Name: "chroot"},
		&simple.Word{Name: "user"},
		&simple.Word{Name: "group"},
		//&simple.SimpleFlag{Name: "master-worker"},
		&parsers.MasterWorker{},
		&parsers.ExternalCheck{},
		&parsers.NbProc{},
		&parsers.NbThread{},
		&parsers.CPUMap{},
		&parsers.Mode{},
		&parsers.MaxConn{},
		&simple.String{Name: "pidfile"},
		&parsers.Socket{},
		&parsers.StatsTimeout{},
		&simple.Number{Name: "tune.ssl.default-dh-param"},
		&simple.String{Name: "ssl-default-bind-options"},
		&simple.Word{Name: "ssl-default-bind-ciphers"},
		&simple.String{Name: "ssl-default-server-options"},
		&simple.Word{Name: "ssl-default-server-ciphers"},
		&parsers.Log{},
		&parsers.ConfigSnippet{},
	})
}

func getFrontendParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Mode{},
		&parsers.MaxConn{},
		&parsers.Bind{},
		&parsers.ACL{},
		&parsers.BindProcess{},
		&simple.Word{Name: "log-tag"},
		&simple.String{Name: "log-format"},
		&simple.String{Name: "log-format-sd"},

		&parsers.Log{},

		&simple.Option{Name: "httpclose"},
		&simple.Option{Name: "forceclose"},
		&simple.Option{Name: "http-buffer-request"},
		&simple.Option{Name: "http-server-close"},
		&simple.Option{Name: "http-keep-alive"},
		&simple.Option{Name: "http-use-htx"},
		&parsers.OptionForwardFor{},
		&simple.Option{Name: "tcplog"},
		&simple.Option{Name: "dontlognull"},
		&simple.Option{Name: "contstats"},
		&simple.Option{Name: "log-separate-errors"},
		&simple.Option{Name: "clitcpka"},

		&simple.Option{Name: "logasap"},
		&parsers.OptionHTTPLog{},

		&simple.Timeout{Name: "http-request"},
		&simple.Timeout{Name: "client"},
		&simple.Timeout{Name: "client-fin"},
		&simple.Timeout{Name: "http-keep-alive"},

		&filters.Filters{},
		&tcp.Requests{},
		&stats.Stats{Mode: "frontend"},
		&http.Requests{Mode: "frontend"},
		&http.Redirect{},

		&simple.Word{Name: "monitor-uri"},

		&parsers.ConfigSnippet{},
		&parsers.UseBackend{},
		&parsers.DefaultBackend{},
		&parsers.StickTable{},
		&http.Responses{Mode: "frontend"},
		&parsers.UniqueIDFormat{},
		&parsers.UniqueIDHeader{},
	})
}

func getBackendParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Mode{},
		&parsers.HashType{},
		&parsers.Balance{},
		&parsers.ACL{},
		&parsers.HTTPCheck{},
		&parsers.BindProcess{},

		&simple.Option{Name: "httpclose"},
		&simple.Option{Name: "forceclose"},
		&simple.Option{Name: "http-buffer-request"},
		&simple.Option{Name: "http-server-close"},
		&simple.Option{Name: "http-keep-alive"},
		&simple.Option{Name: "http-pretend-keepalive"},
		&simple.Option{Name: "http-use-htx"},
		&parsers.OptionForwardFor{},
		&simple.Option{Name: "ssl-hello-chk"},
		&parsers.OptionSmtpchk{},
		&simple.Option{Name: "ldap-check"},
		&parsers.OptionMysqlCheck{},
		&simple.Option{Name: "abortonclose"},
		&parsers.OptionPgsqlCheck{},
		&simple.Option{Name: "tcp-check"},
		&simple.Option{Name: "redis-check"},
		&parsers.OptionRedispatch{},
		&simple.Option{Name: "external-check"},

		&simple.String{Name: "log-tag"},
		&simple.Option{Name: "allbackups"},

		&parsers.OptionHttpchk{},
		&parsers.ExternalCheckPath{},
		&parsers.ExternalCheckCommand{},

		&parsers.Log{},

		&simple.Timeout{Name: "http-request"},
		&simple.Timeout{Name: "queue"},
		&simple.Timeout{Name: "http-keep-alive"},
		&simple.Timeout{Name: "check"},
		&simple.Timeout{Name: "tunnel"},
		&simple.Timeout{Name: "server"},
		&simple.Timeout{Name: "server-fin"},
		&simple.Timeout{Name: "connect"},

		&parsers.DefaultServer{},
		&parsers.Stick{},
		&filters.Filters{},
		&tcp.Requests{},
		&stats.Stats{Mode: "backend"},
		&parsers.HTTPReuse{},
		&http.Requests{Mode: "backend"},
		&http.Redirect{},
		&parsers.Cookie{},
		&parsers.UseServer{},
		&parsers.StickTable{},
		&parsers.ConfigSnippet{},
		&parsers.Server{},
		&simple.Number{Name: "retries"},
		&tcp.Responses{},
		&http.Responses{Mode: "backend"},
	})
}

func getListenParser() *Parsers {
	return createParsers([]ParserInterface{})
}

func getResolverParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Nameserver{},

		&simple.TimeTwoWords{Keywords: []string{"hold", "nx"}},
		&simple.TimeTwoWords{Keywords: []string{"hold", "obsolete"}},
		&simple.TimeTwoWords{Keywords: []string{"hold", "other"}},
		&simple.TimeTwoWords{Keywords: []string{"hold", "refused"}},
		&simple.TimeTwoWords{Keywords: []string{"hold", "timeout"}},
		&simple.TimeTwoWords{Keywords: []string{"hold", "valid"}},

		&simple.Timeout{Name: "resolve"},
		&simple.Timeout{Name: "retry"},

		&simple.Word{Name: "accepted_payload_size"},
		&simple.Word{Name: "parse-resolv-conf"},
		&simple.Word{Name: "resolve_retries"},
	})
}

func getUserlistParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Group{},
		&parsers.User{},
	})
}

func getPeersParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Peer{},
	})
}

func getMailersParser() *Parsers {
	return createParsers([]ParserInterface{
		&simple.TimeTwoWords{Keywords: []string{"timeout", "mail"}},
		&parsers.Mailer{},
	})
}

func getCacheParser() *Parsers {
	return createParsers([]ParserInterface{
		&simple.Number{Name: "total-max-size"},
		&simple.Number{Name: "max-object-size"},
		&simple.Number{Name: "max-age"},
	})
}

func getProgramParser() *Parsers {
	return createParsers([]ParserInterface{
		&simple.String{Name: "command"},
		&simple.String{Name: "user"},
		&simple.String{Name: "group"},
		&simple.Option{Name: "start-on-reload"},
	})
}
