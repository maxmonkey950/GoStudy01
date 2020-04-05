FROM tomcat:jdk11-openjdk
ENV MAVEN_HOME "/usr/local/maven"
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN wget https://downloads.apache.org/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.tar.gz \
&& tar zxvf apache-maven-3.6.3-bin.tar.gz \
&& mv apache-maven-3.6.3 /usr/local/maven \
&& rm -f apache-maven-3.6.3-bin.tar.gz
RUN export PATH=apache-maven-3.6.3/bin:$PATH
RUN export PATH=/usr/local/bin/apache-maven-3.3.9/bin:$PATH
RUN ln -s /usr/local/maven/bin/mvn /usr/local/bin/mvn
CMD ["catalina.sh", "run"]
