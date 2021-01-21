#!/usr/local/bin/python

import configparser
import jenkins
config = configparser.ConfigParser()
config.read(r'/etc/jenkins_jobs/jenkins_jobs.ini')
server = jenkins.Jenkins('http://localhost:8080', username=config['jenkins']['user'], password=config['jenkins']['password'])

print(server.get_whoami())