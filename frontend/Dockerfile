# Build: $ docker build -t frontend .
# Run: $ docker run --rm -it -p 8080:80 frontend
FROM node:16 AS builder
WORKDIR /app
ENV PATH=${PATH}:./node_modules/.bin
ENV NODE_PATH=/app/node_modules
ADD package.json package-lock.json ./
RUN npm install
ADD . .
RUN node --max_old_space_size=4096 ./node_modules/@angular/cli/bin/ng build --prod

# stage 2
FROM nginx:alpine
WORKDIR /usr/share/nginx/html

# Remove all nginx files
RUN rm -rf ./*

# Use angular files
EXPOSE 80
COPY --from=builder /app/dist/frontend .
ENTRYPOINT ["nginx", "-g", "daemon off;"]

# To build this container, run: $ docker build -t swivel-web .
