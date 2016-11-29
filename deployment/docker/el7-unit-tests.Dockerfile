FROM centos:7
WORKDIR /var/tmp
ENTRYPOINT ["/sbin/init"]
RUN yum install -y epel-release
ADD deployment/docker/yum.repo/dev.repo /etc/yum.repos.d/
RUN yum install -y http://repos.mesosphere.io/el/7/noarch/RPMS/mesosphere-el-repo-7-1.noarch.rpm
RUN yum install -y http://openvdc.org/openvdc-release.rpm
RUN yum install -y openvdc
