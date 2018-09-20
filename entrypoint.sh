#!/bin/sh

set -e

if [ -z $ARM_CLIENT_ID ]; then
    echo "Azure Client ID is missing"
    exit 1
fi

if [ -z $ARM_CLIENT_SECRET ]; then
    echo "Azure Client Secret is missing"
    exit 2
fi

if [ -z $ARM_TENANT_ID ]; then
    echo "Azure Tenant ID is missing"
    exit 3
fi

if [ -z $PLUGIN_REGISTRY ]; then
    echo "Registry is missing"
    exit 4
fi

if [ -z $PLUGIN_REPOSITORY ]; then
    echo "Repository is missing"
    exit 5
fi

if [ -z $PLUGIN_TAGS ]; then
    echo "Tags to save are missing, it will only safe latest"
fi

az login --service-principal -u ${ARM_CLIENT_ID} -p ${ARM_CLIENT_SECRET} --tenant ${ARM_TENANT_ID}
ll
TAGS=$(echo "$PLUGIN_TAGS" | sed 's/,/ /g')
echo "Registry: $PLUGIN_REGISTRY"
echo "Repository: $PLUGIN_REPOSITORY"
echo "Tags to save: $TAGS"
echo "Command sample: ./regsitry-cleaner -registry=\"$PLUGIN_REGISTRY\" -repository=\"$PLUGIN_REPOSITORY\" $TAGS"
./registry-cleaner -registry="$PLUGIN_REGISTRY" -repository="$PLUGIN_REPOSITORY" $TAGS