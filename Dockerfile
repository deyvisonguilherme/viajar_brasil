FROM java:8-alpine
MAINTAINER Deyvison Guilherme <deyvison.guilherme@live.com>

ADD target/uberjar/senem.jar /senem/app.jar

EXPOSE 3000

CMD ["java", "-jar", "/senem/app.jar"]
