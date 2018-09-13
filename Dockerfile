FROM microsoft/azure-cli:2.0.45

ADD /registry-cleaner /
ADD /entrypoint.sh /

ENTRYPOINT [ "/entrypoint.sh" ]