#!/usr/local/bin/python

import configparser
import jenkins
config = configparser.ConfigParser()
config.read(r'/etc/jenkins_jobs/jenkins_jobs.ini')
server = jenkins.Jenkins(config['jenkins']['url'], username=config['jenkins']['user'], password=config['jenkins']['password'])

server.build_job('jobs-generator')