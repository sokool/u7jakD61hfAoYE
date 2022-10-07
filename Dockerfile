############################
# STEP 1 build executable binary
############################
FROM golang:1.19-alpine AS builder

# Set env values
ENV HTTP_PORT=8080

WORKDIR /dictionary

COPY . .

# build the application binary
RUN go build -o dict

# Run the api binary.
CMD ["./dict"]
EXPOSE 8080