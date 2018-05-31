package kubeconfig

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

var kubeConfigTemplate = template.Must(template.New("kubeadmConfigTemplate").Parse(`
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{.CA}}
    server: {{.Endpoint}}
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: {{.UserName}}
  name: {{.UserName}}-context
current-context: {{.UserName}}-context
kind: Config
preferences: {}
users:
- name: {{.UserName}}
  user:
    token: {{.Token}}
`))

const homeFilterStr = "~/"

func Generate(config interface{}, path string) error {
	if err := writeConfig(config, path); err != nil {
		return err
	}
	return nil
}

func writeConfig(config interface{}, path string) error {
	p := path
	if index := strings.Index(p, homeFilterStr); index > -1 {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		p = filepath.Join(home, p[index+len(homeFilterStr):])
	}

	dir := filepath.Dir(p)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("Error creating directory: %s", dir)
		}
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := kubeConfigTemplate.Execute(f, config); err != nil {
		return err
	}
	fmt.Printf("Generate the Kubernetes config to `%s`.\n", path)
	return nil
}
