package polaris

import (
	"fmt"
	"testing"
)

func TestGrade(t *testing.T) {
	result, err := RunGrade(&Config{
		Host:  "172.16.10.81",
		Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6InZ6MldqSnhWZGl0c0QyaG8tSEtSNUhpQW1PbEVIOFltOUxnNEhIeHZFcjAifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrby1hZG1pbi10b2tlbi1menA4eCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJrby1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjMwNGQ5NGI5LWI5MjgtNDAyMi05OTE0LWE1MDFhZGQyOWIwOCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTprby1hZG1pbiJ9.sx1_n2kBz5iNI6vLWgfQMHpFz7eNtAw9zWZoRrDPqu3B29VD7-0h3lBcBzrGYzmdpwdPtLKbH8ZODsSzntz60CeUaWjYub_PsIM2iaVfz-zXt-b8m8sYxv7sdv0owkSEeFQNCTY_B4clkkHl5nmhTZHuuzvkqt3Ue6YqZz-howOwQgLKZ2r-XmNYk3k8NyEC4jfZafy28anMqp3CAYLSNW8D4lIIeDm1ZQNXxGR-RraC4x-DoRbRnvOVtHl612jXjV4NzeRWevbZsmpkaGvZxcOx69xx1iqZOKfeAS2wK_OpvHiAqhFa1mL44wrL8aWE8a-WpnvYb2FsRGKVgBZnpg",
		Port:  8443,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}