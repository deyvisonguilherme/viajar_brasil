FROM java:8-alpine
MAINTAINER Your Name <you@example.com>

ADD target/uberjar/senem.jar /senem/app.jar

EXPOSE 3000

CMD ["java", "-jar", "/senem/app.jar"]
