FROM node:21-alpine

RUN apk add openjdk17
RUN npm install -g @openapitools/openapi-generator-cli

RUN apk add go
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

RUN npm install -g @apidevtools/swagger-cli