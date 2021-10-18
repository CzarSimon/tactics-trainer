FROM node:14.18.1-alpine3.14 AS build

# Copy package specifications
WORKDIR /app/web-app
COPY package.json .
COPY package-lock.json .
COPY tsconfig.json .

# Download dependencies application
RUN npm ci

# Copy source
COPY public public
COPY src src

# Build application
RUN npm run build

FROM nginxinc/nginx-unprivileged AS runtime
COPY --from=build /app/web-app/build /usr/share/nginx/html
COPY nginx /etc/nginx/conf.d/
CMD ["nginx", "-g", "daemon off;"]