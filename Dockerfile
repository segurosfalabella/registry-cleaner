FROM microsoft/azure-cli:2.0.45

ADD /registry-cleaner /usr/bin
ADD /entrypoint.sh /
RUN ls -la

ENTRYPOINT [ "/entrypoint.sh" ]