#!/bin/bash

#	jenkins

docker image build --progress=plain -f Dockerfile-jenkins -t guther2/jenkins .
docker container run -d -p 8080:8080 -volume /var/run/docker.sock:/var/run/docker.sock --volume  $(pwd):/mnt/host --volume $SSH_AUTH_SOCK:/ssh-agent --env JENKINS_ADMIN_ID=admin --env JENKINS_ADMIN_PASSWORD=password --env SSH_AUTH_SOCK=/ssh-agent -it --rm --name jenkinscontainer guther2/jenkins


#	python
docker pull python
docker run -d -it --network=host --volume $(pwd)/jenkins_jobs:/etc/jenkins_jobs --rm --name python python

docker exec python bash -c "
pip install --user jenkins-job-builder
/root/.local/bin/jenkins-jobs update /etc/jenkins_jobs/seed.yaml
./etc/jenkins_jobs/jobs_generator.py
"


#	golang

docker build --ssh default --progress=plain -f Dockerfile-golang -t guther2/golang .
docker run --volume $SSH_AUTH_SOCK:/ssh-agent --env SSH_AUTH_SOCK=/ssh-agent -it --rm --name guther2/golangcontainer guther2/golang
