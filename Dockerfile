FROM microsoft/azure-cli:2.0.45

COPY registry-cleaner /
ADD /entrypoint.sh /
RUN ls -la

ENTRYPOINT [ "/entrypoint.sh" ]