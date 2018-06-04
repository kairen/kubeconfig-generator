# Kubeconfig Generator Web UI
Provides Kubenconfig Generator with a web UI.

## Quick Start
Run with Docker:
```shell
$ docker run -dti -p 9090:9090 \
  -e KG_APISERVER_URL=<yout backend server url> \
  --name kg-ui \
  inwinstack/kg-ui:v0.1.0
```

The environment you can set:

| ENV              | Description                                                                  |
| :--------------- | :--------------------------------------------------------------------------- |
| KG_APISERVER_URL | The Kubeconfig Generator backend server URL.(Default: http://localhost:8080) |
| KG_UI_PORT       | The Kubeconfig Generator web UI expose port.(Default: 9090)                  |

## Building from Source
Fisrst, set the environment with Kubeconfig Generator web UI:
```shell
# default: http://localhost:8080
export KG_APISERVER_URL=http://localhost:8080
# default: 9090
export KG_UI_PORT=9090
```

Or you can modified the `config/index.js` file:
```js
// Modified the value with Kubeconfig Generator backend server URL
proxyTable: {
  '/login': {
    target: process.env.KG_APISERVER_URL || 'http://localhost:8080',
    changeOrigin: true,
    ws: true
  },
}

// Modified the value with Kubeconfig Generator web UI expose port
port: process.env.KG_UI_PORT || 9090
```

And then install dependencies and run:
``` bash
$ npm install
$ npm run dev
DONE  Compiled successfully in 475ms

I  Your application is running here: http://0.0.0.0:9090
```
