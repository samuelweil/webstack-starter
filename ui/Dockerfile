FROM node:16-alpine as builder

WORKDIR /build

COPY package*.json ./
RUN npm ci

COPY public public/
COPY src src
COPY tsconfig.json .
RUN npm run build

FROM nginx

COPY --from=builder /build/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/nginx.conf