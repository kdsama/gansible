# Build Docker containers
build:
	docker build -t ubuntu_password -f testservers/Dockerfile.ubuntu .
	docker build -t ubuntu_sshkey -f testservers/Dockerfile.ubuntu.key .
	docker build -t centos_password -f testservers/Dockerfile.centos .
	docker build -t centos_sshkey -f testservers/Dockerfile.centos.key .
	docker build -t redhat_password -f testservers/Dockerfile.redhat .
	docker build -t redhat_sshkey -f testservers/Dockerfile.redhat.key .

# Run Docker containers
run:
	docker run -d -p 2222:22 --name ubuntu_password_container ubuntu_password
	docker run -d -p 2223:22 --name ubuntu_sshkey_container ubuntu_sshkey
	docker run -d -p 2224:22 --name centos_password_container centos_password
	docker run -d -p 2225:22 --name centos_sshkey_container centos_sshkey
	docker run -d -p 2226:22 --name redhat_password_container redhat_password
	docker run -d -p 2227:22 --name redhat_sshkey_container redhat_sshkey

# Remove Docker containers
clean:
	docker stop ubuntu_password_container ubuntu_sshkey_container centos_password_container centos_sshkey_container redhat_password_container redhat_sshkey_container
	docker rm ubuntu_password_container ubuntu_sshkey_container centos_password_container centos_sshkey_container redhat_password_container redhat_sshkey_container

# Clean up
clean-all: clean
	docker rmi ubuntu_password ubuntu_sshkey centos_password centos_sshkey redhat_password redhat_sshkey
