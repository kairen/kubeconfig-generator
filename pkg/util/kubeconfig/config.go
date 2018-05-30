package kubeconfig

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
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

func Generate(config interface{}, path string) error {
	if err := writeConfig(config, path); err != nil {
		return err
	}
	return nil
}

func writeConfig(config interface{}, path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("Error creating directory: %s", dir)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := kubeConfigTemplate.Execute(f, config); err != nil {
		return err
	}
	return nil
}
