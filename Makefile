# Build Docker containers
build:
	docker build -t ubuntu_password -f testservers/Dockerfile.ubuntu .





# Remove Docker containers
clean:
	docker rm --force server1 
	docker rm --force server2
	docker rm --force server3 
	docker rm --force server4 



# Run Docker containers
env-up: 
	docker run -d -p 2225:22 --name server4 mmumshad/ubuntu-ssh-enabled
	docker run -d -p 2224:22 --name server3 mmumshad/ubuntu-ssh-enabled
	docker run -d -p 2223:22 --name server2 mmumshad/ubuntu-ssh-enabled
	docker run -d -p 2222:22 --name server1 mmumshad/ubuntu-ssh-enabled 

