#!/usr/bin/env bash

#args:

source setenv.sh

#
echo "#### Deploying ..."

#
function deploy() {
    echo "Pushing service..."

    env_subst manifest_env.yml manifest.yml

    cat manifest.yml

    ##
    cf push -f manifest.yml; if [ $? -ne 0 ]; then
        return 1
    else
        return 0
    fi
}

deploy; if [ $? -ne 0 ]; then
    echo "#### Deploy failed"
    exit 1
fi


echo "#### Deploy successful"

exit 0
##
