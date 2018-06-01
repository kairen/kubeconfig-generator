# Kubeconfig Generator Web UI
Provides Kubenconfig Generator with a web UI.

## Quick Start
First, modified the `config/index.js` file to your  Kubernetes Generator backend:
```js
proxyTable: {
  '/login': {
    target: 'http://127.0.0.1:32400',
    changeOrigin: true,
    ws: true
  },
},
```

And then install dependencies and run:
``` bash
$ npm install
$ npm run dev
DONE  Compiled successfully in 475ms

I  Your application is running here: http://0.0.0.0:8080
```
