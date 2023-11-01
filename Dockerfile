# Lightweight with Alpine
FROM nginx:1.23.0-alpine

# Install dependancies
RUN apk add --update docker openrc go
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.29.2
RUN trivy i ubuntu

# Build frontend outside & copy 
COPY /front/dist /usr/share/nginx/html

# Build backend inside for alpine architecture
COPY /back /docktor
WORKDIR /docktor
RUN go build main.go

# Activate host docker
RUN rc-update add docker boot

# Prepare start frontend & backend
RUN echo "nginx -g 'daemon off;' &" >> custom-start.sh
RUN echo "./main" >> custom-start.sh
RUN chmod u+x custom-start.sh

# Expose frontend & backend
EXPOSE 3030 4040

CMD ["./custom-start.sh"]