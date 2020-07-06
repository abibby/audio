#!/usr/bin/env bash

# echo 0.005 \
#     | ./bin/sin | ./bin/add 0.5 | ./bin/mul 2 \
#     | ./bin/sin | ./bin/add 0.5 | ./bin/mul 440 \
#     | ./bin/sin \
#     | ./bin/gate \
#     | ./bin/play
# echo 0.5 \
#     | ./bin/sin | ./bin/add 1 | ./bin/mul 440 \
#     | ./bin/sin \
#     | ./bin/play
echo 440 | ./bin/sin \
    | ./bin/patch \
        <(echo 880 | ./bin/sin | ./bin/mul 0.2) \
        <(echo 1320 | ./bin/sin | ./bin/mul 0.1) \
    | ./bin/play
# echo 440 | ./bin/sin \
#     | ./bin/play