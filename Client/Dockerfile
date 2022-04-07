FROM node:12-alpine3.14 as build

MAINTAINER yeol

ADD . /app
WORKDIR /app

RUN yarn config set registry https://registry.npm.taobao.org/ && \
    yarn install && \
    yarn build


FROM nginx:latest

COPY --from=build /app/dist/ /usr/share/nginx/html/
COPY ./config/nginx/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080
RUN service nginx start