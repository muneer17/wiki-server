# Build frontend
FROM node:14 as frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# Build backend
FROM golang:1.16 as backend
WORKDIR /app
COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM nginx:alpine
RUN mkdir /app
COPY --from=frontend /app/build /usr/share/nginx/html
COPY --from=backend /app/main /app
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
RUN echo "#!/bin/sh" > /app/start.sh && \
    echo "/app/main & nginx -g 'daemon off;'" >> /app/start.sh && \
    chmod +x /app/start.sh
CMD ["/app/start.sh"]