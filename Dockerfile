FROM devops-hub.tutorabc.com.cn/library/cy:0.3
ADD ./main /opt/
WORKDIR /opt/
ENTRYPOINT ["/opt/main"]