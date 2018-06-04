FROM node:8.9-alpine
LABEL maintainer="Ellis Wu<ellis.w@inwinstack.com>"

ENV KG_APISERVER_URL http://localhost:8080
ENV KG_UI_PORT 9090

COPY . /kg-ui
WORKDIR /kg-ui
RUN npm install

EXPOSE 9090

ENTRYPOINT ["npm", "run", "dev"]
