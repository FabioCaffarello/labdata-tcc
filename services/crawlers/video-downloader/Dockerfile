FROM python:3.10-slim-buster as devimage

COPY dist/*.tar.gz .
RUN pip install *.tar.gz
RUN rm -rf *.tar.gz

COPY downloader/ /app/downloader/
COPY tests/ /app/tests/
COPY hack/ /app/hack/

RUN chmod +x /app/hack/check-integration.sh

WORKDIR /app


ENTRYPOINT ["python", "/app/downloader/main.py"]


FROM devimage as prodimage

COPY --from=devimage /app/downloader/ /app/downloader/

ENTRYPOINT ["python", "/app/downloader/main.py"]