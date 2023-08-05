#!/bin/bash

echo "
=================================
||                             ||
||                             ||
||       FEATURE MATCHER       ||
||                             ||
||                       v0.1  ||
=================================
"

if [[ $1 -eq "prod" ]]; then
    LOGURU_LEVEL=INFO uvicorn \
        src.main:app \
        --host 0.0.0.0 \
        --port 30557 \
        --workers 1
else
    LOGURU_LEVEL=DEBUG uvicorn src.main:app --reload
fi
