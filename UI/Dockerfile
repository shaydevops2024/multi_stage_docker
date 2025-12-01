FROM python:3.12-alpine

WORKDIR /app

COPY index.html .
COPY static ./static

EXPOSE 9000

CMD ["python", "-m", "http.server", "9000"]

