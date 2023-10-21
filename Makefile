# Build Docker containers
build:
	docker build -t ubuntu_password -f testservers/Dockerfile.ubuntu .


# Run Docker containers
run:
	docker run -d -p 2222:22 --name ubuntu_password_container ubuntu_password


# Remove Docker containers
clean:
	docker stop ubuntu_password_container 
	docker rm ubuntu_password_container 

# Clean up
clean-all: clean
	docker rmi ubuntu_password 
