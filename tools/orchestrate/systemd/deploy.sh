#!/bin/bash

# Abort if anything fails
set -e

DASHES=$(echo $SERVICE | tr "." "-")

ssh -t -oStrictHostKeyChecking=no -oUserKnownHostsFile=/dev/null "$TARGET_USERNAME"@"$DEPLOYMENT_TARGET" << EOF
    # Abort if anything fails
    set -e

    cd $TARGET_DIRECTORY

    if [ ! -d "src" ]; then
        git clone https://github.com/shadracnicholas/home-automation.git src
    fi

    cd src

    git checkout -- .
    git checkout master
    git pull

    # The or true will suppress the non-zero exit code if the service does not exist
    # Note that this will still print an error to the console though
    sudo systemctl stop "$DASHES".service || true

    cd $SERVICE
    bash ../tools/orchestrate/systemd/$LANG.sh

    echo "Creating systemd service"
    # The quotes are needed around the variable to preserve the new lines
    echo "$SYSTEMD_SERVICE" | sudo tee /lib/systemd/system/$DASHES.service
    sudo chmod 644 /lib/systemd/system/$DASHES.service

    sudo systemctl daemon-reload
    sudo systemctl enable "$DASHES".service
    sudo systemctl start "$DASHES".service

EOF
