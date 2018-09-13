FROM microsoft/azure-cli:2.0.45

COPY registry-cleaner ./
ADD /entrypoint.sh /

ENTRYPOINT [ "/entrypoint.sh" ]