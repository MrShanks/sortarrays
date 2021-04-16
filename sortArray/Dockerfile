FROM openjdk:8-jdk-alpine
RUN addgroup -S spring && adduser -S spring -G spring
USER spring:spring
ARG DEPENDENCY=target/dependency
EXPOSE 8080/tcp
COPY ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY ${DEPENDENCY}/META-INF /app/META-INF
COPY ${DEPENDENCY}/BOOT-INF/classes /app
COPY src/main/resources/application.properties /app/application.properties
ENTRYPOINT ["java","-cp","app:app/lib/*","sortArray.SortArrayApplication"]
