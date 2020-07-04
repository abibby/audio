echo 0.005 \
    | ./bin/sin | ./bin/add 0.5 | ./bin/mul 2 \
    | ./bin/sin | ./bin/add 0.5 | ./bin/mul 440 \
    | ./bin/sin \
    | ./bin/gate \
    | ./bin/play