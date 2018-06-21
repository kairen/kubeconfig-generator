<template>
  <v-app id="inspire">
    <v-container fill-height>
      <v-layout>
        <v-flex>
          <span class="headline">Kubernetes Configuration</span>
          <v-divider class="my-3"></v-divider>
          <v-alert :value="true" type="info">
            A file that is used to configure access to a cluster is sometimes called a kubeconfig file. This is a generic way of referring to configuration files. It does not mean that there is a file named kubeconfig.
          </v-alert>
          <div class="subheading my-3">
            <span>Copy your kubeconfig or </span>
            <a v-on:click="download">download</a>
            <span> it. And move the kubeconfig to the </span>
            <code>$HOME/.kube</code>
            <span>directory:</span>
          </div>
          <markup v-if="kubeconfig" :lang="'yaml'">{{ kubeconfig }}</markup>

          <span class="headline">Install kubectl binary via curl</span>
          <v-divider class="my-3"></v-divider>
          <v-tabs color="cyan" dark slider-color="yellow" class="mb-3">
            <template v-for="item in installKCTabs">
              <v-tab :key="item.title" ripple>{{ item.title }}</v-tab>
              <v-tab-item :key="item.title">
                <v-card>
                  <v-card-text>
                    <template v-for="step in item.step">
                      <span v-if="step.title" class="subheading" :key="step.title">{{ step.title }}</span>
                      <markup v-if="kubeconfig && step.cmd" :lang="item.lang" :key="step.cmd">{{ step.cmd }}</markup>
                    </template>
                  </v-card-text>
                </v-card>
              </v-tab-item>
            </template>
          </v-tabs>
          <div class="body-1 mb-3">
            <span>For more details of install kubectl, please see </span>
            <a href="https://kubernetes.io/docs/tasks/tools/install-kubectl/" target="_blank">
              Install and Set Up kubectl
            </a>
            <span>.</span>
          </div>

          <span class="headline">Using</span>
          <v-divider class="my-3"></v-divider>
          <v-tabs color="cyan" dark slider-color="yellow">
            <template v-for="item in usingKCTabs">
              <v-tab :key="item.title" ripple>{{ item.title }}</v-tab>
              <v-tab-item :key="item.title">
                <v-card>
                  <v-card-text>
                    <markup v-if="kubeconfig && item.cmd" :lang="item.lang" :key="item.title">{{ item.cmd }}</markup>
                  </v-card-text>
                </v-card>
              </v-tab-item>
            </template>
          </v-tabs>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script>
import Markup from '@/components/markup'

const cmd = {
  linux: {
    curl: [
      'STABLE_VERSION=$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)',
      'curl -LO https://storage.googleapis.com/kubernetes-release/release/$STABLE_VERSION/bin/linux/amd64/kubectl'
    ],
    make: [
      'chmod +x ./kubectl',
      'sudo mv ./kubectl /usr/local/bin/kubectl'
    ]
  },
  windows: {
    curl: [
      'curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.10.3/bin/windows/amd64/kubectl.exe'
    ]
  },
  macOS: {
    curl: [
      'STABLE_VERSION=$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)',
      'curl -LO https://storage.googleapis.com/kubernetes-release/release/$STABLE_VERSION/bin/darwin/amd64/kubectl'
    ],
    make: [
      'chmod +x ./kubectl',
      'sudo mv ./kubectl /usr/local/bin/kubectl'
    ]
  }
}

export default {
  name: 'Home',
  components: {
    Markup
  },
  data () {
    return {
      kubeconfig: '',
      installKCTabs: [
        {
          title: 'Linux',
          lang: 'cli',
          step: [
            {
              title: '1. Download the latest release with the command:',
              cmd: cmd.linux.curl.join('\n')
            },
            {
              title: '2. Make the kubectl binary executable and move the binary in to you Path:',
              cmd: cmd.linux.make.join('\n')
            }
          ]
        },
        {
          title: 'Windows',
          lang: 'cli',
          step: [
            {
              title: '1. If you have curl installed, use this command:',
              cmd: cmd.windows.curl.join('\n')
            },
            {
              title: '2. Add the binary in to your PATH.'
            }
          ]
        },
        {
          title: 'MacOS',
          lang: 'cli',
          step: [
            {
              title: '1. Download the latest release with the command:',
              cmd: cmd.macOS.curl.join('\n')
            },
            {
              title: '2. Make the kubectl binary executable and move the binary in to you Path:',
              cmd: cmd.macOS.make.join('\n')
            }
          ]
        }
      ],
      usingKCTabs: [
        {
          title: 'Linux',
          lang: 'cli',
          cmd: 'kubectl get po -n <namespace>'
        },
        {
          title: 'Windows',
          lang: 'cli',
          cmd: 'kubectl get po -n <namespace>'
        },
        {
          title: 'MacOS',
          lang: 'cli',
          cmd: 'kubectl get po -n <namespace>'
        }
      ]
    }
  },
  created () {
    if (this.$store.getters.username) {
      var yaml = require('js-yaml')
      var kubeconfigJson = {
        apiVersion: 'v1',
        clusters: [
          {
            cluster: {
              'certificate-authority-data': this.$store.getters.ca,
              server: this.$store.getters.endpoint
            },
            name: 'kubernetes'
          }
        ],
        contexts: [
          {
            context: {
              cluster: 'kubernetes',
              user: this.$store.getters.username
            },
            name: this.$store.getters.username + '-context'
          }
        ],
        'current-context': this.$store.getters.username + '-context',
        kind: 'Config',
        preferences: {},
        users: [
          {
            name: this.$store.getters.username,
            user: {
              token: this.$store.getters.token
            }
          }
        ]
      }
      var kubeconfigYaml = yaml.safeDump(kubeconfigJson)
      this.kubeconfig = kubeconfigYaml
    }
  },
  methods: {
    download () {
      var FileSaver = require('file-saver')
      var file = new File([this.kubeconfig], {type: ''})
      FileSaver.saveAs(file, 'config')
    }
  }
}
</script>
