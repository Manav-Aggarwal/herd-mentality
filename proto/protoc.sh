#!/usr/bin/env bash

set -eo pipefail

buf generate --template="buf.gen.yaml" --config="buf.yaml"
