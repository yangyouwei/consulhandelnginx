upstream {{.Name}} {
	zone upstream-{{.Name}} 64k;{{range .Servers}}
	server {{.Addr}}:{{.Port}} max_fails=3 fail_timeout=60 weight=1;{{else}}server 127.0.0.1:65535; # force a 502{{end}}
	ip_hash;
}