package caddyfile

import "testing"

var tests = []struct {
	caddyfile, json string
}{
	{ // 0
		caddyfile: `foo {
	root /bar
}`,
		json: `[{"hosts":["foo"],"body":[["root","/bar"]]}]`,
	},
	{ // 1
		caddyfile: `host1, host2 {
	dir {
		def
	}
}`,
		json: `[{"hosts":["host1","host2"],"body":[["dir",[["def"]]]]}]`,
	},
	{ // 2
		caddyfile: `host1, host2 {
	dir abc {
		def ghi
		jkl
	}
}`,
		json: `[{"hosts":["host1","host2"],"body":[["dir","abc",[["def","ghi"],["jkl"]]]]}]`,
	},
	{ // 3
		caddyfile: `host1:1234, host2:5678 {
	dir abc {
	}
}`,
		json: `[{"hosts":["host1:1234","host2:5678"],"body":[["dir","abc",[]]]}]`,
	},
	{ // 4
		caddyfile: `host {
	foo "bar baz"
}`,
		json: `[{"hosts":["host"],"body":[["foo","bar baz"]]}]`,
	},
	{ // 5
		caddyfile: `host, host:80 {
	foo "bar \"baz\""
}`,
		json: `[{"hosts":["host","host:80"],"body":[["foo","bar \"baz\""]]}]`,
	},
	{ // 6
		caddyfile: `host {
	foo "bar
baz"
}`,
		json: `[{"hosts":["host"],"body":[["foo","bar\nbaz"]]}]`,
	},
	{ // 7
		caddyfile: `host {
	dir 123 4.56 true
}`,
		json: `[{"hosts":["host"],"body":[["dir","123","4.56","true"]]}]`, // NOTE: I guess we assume numbers and booleans should be encoded as strings...?
	},
	{ // 8
		caddyfile: `http://host, https://host {
}`,
		json: `[{"hosts":["host:http","host:https"],"body":[]}]`, // hosts in JSON are always host:port format (if port is specified), for consistency
	},
	{ // 9
		caddyfile: `host {
	dir1 a b
	dir2 c d
}`,
		json: `[{"hosts":["host"],"body":[["dir1","a","b"],["dir2","c","d"]]}]`,
	},
	{ // 10
		caddyfile: `host {
	dir a b
	dir c d
}`,
		json: `[{"hosts":["host"],"body":[["dir","a","b"],["dir","c","d"]]}]`,
	},
	{ // 11
		caddyfile: `host {
	dir1 a b
	dir2 {
		c
		d
	}
}`,
		json: `[{"hosts":["host"],"body":[["dir1","a","b"],["dir2",[["c"],["d"]]]]}]`,
	},
	{ // 12
		caddyfile: `host1 {
	dir1
}

host2 {
	dir2
}`,
		json: `[{"hosts":["host1"],"body":[["dir1"]]},{"hosts":["host2"],"body":[["dir2"]]}]`,
	},
}

func TestToJSON(t *testing.T) {
	for i, test := range tests {
		output, err := ToJSON([]byte(test.caddyfile))
		if err != nil {
			t.Errorf("Test %d: %v", i, err)
		}
		if string(output) != test.json {
			t.Errorf("Test %d\nExpected:\n'%s'\nActual:\n'%s'", i, test.json, string(output))
		}
	}
}

func TestFromJSON(t *testing.T) {
	for i, test := range tests {
		output, err := FromJSON([]byte(test.json))
		if err != nil {
			t.Errorf("Test %d: %v", i, err)
		}
		if string(output) != test.caddyfile {
			t.Errorf("Test %d\nExpected:\n'%s'\nActual:\n'%s'", i, test.caddyfile, string(output))
		}
	}
}
