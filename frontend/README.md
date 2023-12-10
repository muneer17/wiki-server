# wiki-server

This is a brief description of your project.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Docker installed on your machine. You can download Docker [here](https://www.docker.com/products/docker-desktop).

### Building the Docker Image

To build the Docker image, navigate to the directory containing the Dockerfile and run the following command:

docker build -t your-image-name .

This will build a Docker image with the tag your-image-name. You can choose any name you like.

Running the Docker Container
After the image has been built, you can run the server with the following command:

This will start a Docker container from the image you built, and map port 8080 in the container to port 8080 on your machine.

Now, you should be able to access the server at http://localhost:8080.
